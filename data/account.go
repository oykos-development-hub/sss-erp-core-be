package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Account struct
type Account struct {
	ID           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	ParentID     *int      `db:"parent_id"`
	SerialNumber string    `db:"serial_number"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Account) Table() string {
	return "accounts"
}

// GetAll gets all records from the database, using upper
func (t *Account) GetAll(condition *up.Cond) ([]*Account, error) {
	collection := upper.Collection(t.Table())
	var all []*Account
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("serial_number").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Account) Get(id int) (*Account, error) {
	var one Account
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Account) Update(m Account) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Account) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Account) Insert(m Account) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
