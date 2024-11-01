package data

import (
	"database/sql"
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"
)

// Setting struct
type Setting struct {
	ID           int          `db:"id,omitempty"`
	Title        string       `db:"title"`
	Abbreviation string       `db:"abbreviation"`
	Entity       string       `db:"entity"`
	Description  *string      `db:"description,omitempty"`
	ParentID     *int         `db:"parent_id"`
	Value        *string      `db:"value,omitempty"`
	Color        *string      `db:"color,omitempty"`
	Icon         *string      `db:"icon,omitempty"`
	CreatedAt    time.Time    `db:"created_at,omitempty"`
	UpdatedAt    time.Time    `db:"updated_at,omitempty"`
	DeletedAt    sql.NullTime `db:"deleted_at,omitempty"`
}

// Table returns the table name
func (t *Setting) Table() string {
	return "settings"
}

// GetAll gets all records from the database, using upper
func (t *Setting) GetAll(page *int, size *int, conditions *up.AndExpr) ([]*Setting, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Setting
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

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Setting) Get(id int) (*Setting, error) {
	var one Setting
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Setting) Update(m Setting) error {
	m.UpdatedAt = time.Now()

	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	if err := res.Update(&m); err != nil {
		return newErrors.Wrap(err, "upper update")
	}

	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Setting) Delete(id int) error {
	var setting Setting

	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})

	err := res.One(&setting)
	if err != nil {
		return newErrors.WrapNotFoundError(err, "get setting")
	}

	setting.DeletedAt = sql.NullTime{Time: time.Now()}

	if err := res.Update(&setting); err != nil {
		return newErrors.Wrap(err, "upper update")
	}

	return nil
}

// Insert inserts a model into the database, using upper
func (t *Setting) Insert(m Setting) (int, error) {
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
