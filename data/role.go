package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Role struct
type Role struct {
	ID           int       `db:"id,omitempty"`
	Title        string    `db:"title"`
	Abbreviation string    `db:"abbreviation"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Role) Table() string {
	return "roles"
}

// GetAll gets all records from the database, using upper
func (t *Role) GetAll(condition *up.Cond) ([]*Role, error) {
	collection := Upper.Collection(t.Table())
	var all []*Role
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Role) Get(id int) (*Role, error) {
	var one Role
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Role) Update(m Role) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Role) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Role) Insert(m Role) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
