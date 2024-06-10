package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/contextutil"

	up "github.com/upper/db/v4"
)

// Account struct
type Account struct {
	ID           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	ParentID     *int      `db:"parent_id"`
	SerialNumber string    `db:"serial_number"`
	Version      int       `db:"version"`
	CreatedAt    time.Time `db:"created_at,omitempty"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Account) Table() string {
	return "accounts"
}

// GetAll gets all records from the database, using upper
func (t *Account) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Account, int, error) {
	collection := Upper.Collection(t.Table())
	var all []*Account
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, -1, err
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.OrderBy("serial_number").All(&all)
	if err != nil {
		return nil, -2, err
	}

	return all, int(total), err
}

// Get gets one record from the database, by id, using upper
func (t *Account) Get(id int) (*Account, error) {
	var one Account
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Account) Update(ctx context.Context, m Account) error {

	m.UpdatedAt = time.Now()

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		// Set the user ID in the session
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}
		// Perform the update within the transaction
		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Account) Delete(ctx context.Context, id int) error {

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Account) Insert(ctx context.Context, m Account) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {
		// Set the user ID in the session
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}
		// Perform the update within the transaction
		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
