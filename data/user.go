package data

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	up "github.com/upper/db/v4"
	"gitlab.sudovi.me/erp/core-ms-api/contextutil"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User is the type for a user
type User struct {
	ID             int       `db:"id,omitempty"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	SecondaryEmail *string   `db:"secondary_email,omitempty"`
	Active         bool      `db:"active"`
	Password       string    `db:"password"`
	Pin            string    `db:"pin"`
	Phone          string    `db:"phone"`
	VerifiedEmail  bool      `db:"verified_email"`
	VerifiedPhone  bool      `db:"verified_phone"`
	FolderId       *int      `db:"folder_id"`
	CreatedAt      time.Time `db:"created_at,omitempty"`
	UpdatedAt      time.Time `db:"updated_at"`
	RoleId         *int      `db:"role_id"`
}

// Table returns the table name associated with this model in the database
func (u *User) Table() string {
	return "users"
}

// GetAll returns a slice of all users
func (u *User) GetAll(page *int, size *int, conditions *up.AndExpr) ([]*User, *uint64, error) {
	collection := Upper.Collection(u.Table())

	var all []*User
	var res up.Result

	if conditions != nil {
		res = collection.Find(conditions)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper get all")
	}

	return all, &total, nil
}

// GetByEmail gets one user, by email
func (u *User) GetByEmail(email string) (*User, error) {
	var theUser User
	collection := Upper.Collection(u.Table())
	res := collection.Find(up.Cond{"email =": email})
	err := res.One(&theUser)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}

	return &theUser, nil
}

// Get gets one user by id
func (u *User) Get(id int) (*User, error) {
	var theUser User
	collection := Upper.Collection(u.Table())
	res := collection.Find(up.Cond{"id =": id})

	err := res.One(&theUser)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}

	return &theUser, nil
}

// Update updates a user record in the database
func (u *User) Update(ctx context.Context, theUser User) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	if theUser.RoleId != nil && *theUser.RoleId == 0 {
		theUser.RoleId = nil
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(u.Table())
		res := collection.Find(theUser.ID)
		if err := res.Update(&theUser); err != nil {
			return newErrors.Wrap(err, "upper update")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// ValidatePassword checks if the given password meets the required criteria
func ValidatePassword(password string) error {
	var (
		upperCase    = regexp.MustCompile(`[A-Z]`)
		lowerCase    = regexp.MustCompile(`[a-z]`)
		number       = regexp.MustCompile(`[0-9]`)
		specialChars = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	)

	if len(password) < 8 {
		err := errors.New("password must be at least 8 characters long")
		return newErrors.Wrap(err, "validation")
	}
	if !upperCase.MatchString(password) {
		err := errors.New("password must contain at least one uppercase letter")
		return newErrors.Wrap(err, "validation")
	}
	if !lowerCase.MatchString(password) {
		err := errors.New("password must contain at least one lowercase letter")
		return newErrors.Wrap(err, "validation")
	}
	if !number.MatchString(password) {
		err := errors.New("password must contain at least one number")
		return newErrors.Wrap(err, "validation")
	}
	if !specialChars.MatchString(password) {
		err := errors.New("password must contain at least one special character")
		return newErrors.Wrap(err, "validation")
	}

	return nil
}

// Insert inserts a new user, and returns the newly inserted id
func (u *User) Insert(ctx context.Context, theUser User) (int, error) {
	if err := ValidatePassword(theUser.Password); err != nil {
		return 0, newErrors.Wrap(err, "validate password")
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(theUser.Password), 12)
	if err != nil {
		return 0, newErrors.Wrap(err, "bcrycpt generate password")
	}

	theUser.CreatedAt = time.Now()
	theUser.UpdatedAt = time.Now()
	theUser.Password = string(newHash)
	if theUser.RoleId != nil && *theUser.RoleId == 0 {
		theUser.RoleId = nil
	}

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return 0, newErrors.Wrap(err, "context get user id")
	}

	var id int

	err = Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec query")
		}

		collection := sess.Collection(u.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(theUser); err != nil {
			return newErrors.Wrap(err, "upper insert")
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// ResetPassword resets a users's password, by id, using supplied password
func (u *User) ResetPassword(id int, password string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return newErrors.Wrap(err, "bcrypt generate password")
	}

	theUser, err := u.Get(id)
	if err != nil {
		return newErrors.Wrap(err, "upper get")
	}

	theUser.Password = string(newHash)

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, id)

	err = theUser.Update(ctx, *theUser)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}

	return nil
}

// PasswordMatches verifies a supplied password against the hash stored in the database.
// It returns true if valid, and false if the password does not match, or if there is an
// error. Note that an error is only returned if something goes wrong (since an invalid password
// is not an error -- it's just the wrong password))
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			// some kind of error occurred
			return false, newErrors.Wrap(err, "bcrypt match password")
		}
	}

	return true, nil
}

// Delete deletes a record from the database by id, using upper
func (t *User) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID not found in context")
		return newErrors.Wrap(err, "context get user id")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return newErrors.Wrap(err, "upper exec")
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return newErrors.Wrap(err, "upper delete")
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
