package data

import (
	"errors"
	"regexp"
	"time"

	up "github.com/upper/db/v4"
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
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	RoleId         int       `db:"role_id"`
}

// Table returns the table name associated with this model in the database
func (u *User) Table() string {
	return "users"
}

// GetAll returns a slice of all users
func (u *User) GetAll(page *int, size *int, conditions *up.AndExpr) ([]*User, *uint64, error) {
	collection := upper.Collection(u.Table())

	var all []*User
	var res up.Result

	if conditions != nil {
		res = collection.Find(conditions)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, nil
}

// GetByEmail gets one user, by email
func (u *User) GetByEmail(email string) (*User, error) {
	var theUser User
	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{"email =": email})
	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	return &theUser, nil
}

// Get gets one user by id
func (u *User) Get(id int) (*User, error) {
	var theUser User
	collection := upper.Collection(u.Table())
	res := collection.Find(up.Cond{"id =": id})

	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	return &theUser, nil
}

// Update updates a user record in the database
func (u *User) Update(theUser User) error {
	theUser.UpdatedAt = time.Now()
	collection := upper.Collection(u.Table())
	res := collection.Find(theUser.ID)
	err := res.Update(&theUser)
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
		return errors.New("password must be at least 8 characters long")
	}
	if !upperCase.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !lowerCase.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !number.MatchString(password) {
		return errors.New("password must contain at least one number")
	}
	if !specialChars.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// Insert inserts a new user, and returns the newly inserted id
func (u *User) Insert(theUser User) (int, error) {
	if err := ValidatePassword(theUser.Password); err != nil {
		return 0, err
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(theUser.Password), 12)
	if err != nil {
		return 0, err
	}

	theUser.CreatedAt = time.Now()
	theUser.UpdatedAt = time.Now()
	theUser.Password = string(newHash)

	collection := upper.Collection(u.Table())
	res, err := collection.Insert(theUser)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

// ResetPassword resets a users's password, by id, using supplied password
func (u *User) ResetPassword(id int, password string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	theUser, err := u.Get(id)
	if err != nil {
		return err
	}

	theUser.Password = string(newHash)

	err = theUser.Update(*theUser)
	if err != nil {
		return err
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
			return false, err
		}
	}

	return true, nil
}

// Delete deletes a record from the database by id, using upper
func (t *User) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}
