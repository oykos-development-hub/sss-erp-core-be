package data

import (
	"time"

	db2 "github.com/upper/db/v4"
	up "github.com/upper/db/v4"
)

// Supplier struct
type Supplier struct {
	ID            int    `db:"id,omitempty"`
	Title         string `db:"title" validate:"required"`
	Abbreviation  string `db:"abbreviation"`
	OfficialID    string `db:"official_id"`
	Address       string `db:"address"`
	Description   string `db:"description"`
	FolderID      int    `db:"folder_id"`
	Entity        string `db:"entity"`
	ParentID      *int   `db:"supplier_id"`
	BankAccounts  []string
	TaxPercentage float32   `db:"tax_percentage"`
	CreatedAt     time.Time `db:"created_at,omitempty"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Supplier) Table() string {
	return "suppliers"
}

// GetAll gets all records from the database, using upper
func (t *Supplier) GetAll(page *int, size *int, condition *up.AndExpr) ([]*Supplier, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*Supplier
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

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Supplier) Get(id int) (*Supplier, error) {
	var one Supplier
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}

	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Supplier) Update(m Supplier) error {
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
func (t *Supplier) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Supplier) Insert(tx db2.Session, m Supplier) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := tx.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
