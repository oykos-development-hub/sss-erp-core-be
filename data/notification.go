package data

import (
	"encoding/json"
	"time"

	up "github.com/upper/db/v4"
)

// Notification struct
type Notification struct {
	ID          int             `db:"id,omitempty"`
	FromContent string          `db:"from_content"`
	FromUserID  int             `db:"from_user_id"`
	ToUserID    int             `db:"to_user_id"`
	Path        string          `db:"path"`
	Module      string          `db:"module"`
	Content     string          `db:"content"`
	IsRead      bool            `db:"is_read"`
	Data        json.RawMessage `db:"data"`
	CreatedAt   time.Time       `db:"created_at,omitempty"`
	UpdatedAt   time.Time       `db:"updated_at"`
}

// Table returns the table name
func (t *Notification) Table() string {
	return "notifications"
}

// GetAll gets all records from the database, using upper
func (t *Notification) GetAll(page *int, size *int, condition *up.Cond) ([]*Notification, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Notification
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
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

	res.OrderBy("is_read", "-created_at")

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
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
