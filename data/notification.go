package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Notification struct
type Notification struct {
	ID         int       `db:"id,omitempty"`
	From       string    `db:"from"`
	FromUserID int       `db:"from_user_id"`
	ToUserID   int       `db:"to_user_id"`
	Module     string    `db:"module"`
	Content    string    `db:"content"`
	IsRead     bool      `db:"is_read"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Notification) Table() string {
	return "notifications"
}

// GetAll gets all records from the database, using upper
func (t *Notification) GetAll(condition *up.Cond) ([]*Notification, error) {
	collection := upper.Collection(t.Table())
	var all []*Notification
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
func (t *Notification) Get(id int) (*Notification, error) {
	var one Notification
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Notification) Update(m Notification) error {
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
func (t *Notification) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Notification) Insert(m Notification) (int, error) {
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
