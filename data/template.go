package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"
)

// Template struct
type Template struct {
	ID        int       `db:"id,omitempty"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Template) Table() string {
	return "templates"
}

// GetAll gets all records from the database, using Upper
func (t *Template) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*Template, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Template
	var res up.Result

	if condition != nil {
		res = collection.Find(condition)
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

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order")
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *Template) Get(id int) (*Template, error) {
	var one Template
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Template) Update(m Template) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Template) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Template) Insert(m Template) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
