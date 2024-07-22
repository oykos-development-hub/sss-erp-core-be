package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// CustomerSupport struct
type CustomerSupport struct {
	ID                      int       `db:"id,omitempty"`
	UserDocumentationFileID int       `db:"user_documentation_file_id"`
	CreatedAt               time.Time `db:"created_at,omitempty"`
	UpdatedAt               time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *CustomerSupport) Table() string {
	return "customer_supports"
}

// GetAll gets all records from the database, using upper
func (t *CustomerSupport) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*CustomerSupport, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*CustomerSupport
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

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *CustomerSupport) Get(id int) (*CustomerSupport, error) {
	var one CustomerSupport
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *CustomerSupport) Update(m CustomerSupport) error {
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
func (t *CustomerSupport) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *CustomerSupport) Insert(m CustomerSupport) (int, error) {
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
