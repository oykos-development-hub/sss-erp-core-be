package data

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	up "github.com/upper/db/v4"
)

// BankAccount struct
type BankAccount struct {
    ID        int       `db:"id,omitempty"`
    Title     string    `db:"title"`
		SupplierID int `db:"supplier_id"`
    CreatedAt time.Time `db:"created_at,omitempty"`
    UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *BankAccount) Table() string {
    return "bank_accounts"
}

// GetAll gets all records from the database, using upper
func (t *BankAccount) GetAll(page *int, size *int, condition *up.AndExpr, orders []interface{}) ([]*BankAccount, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*BankAccount
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
		return nil,nil, err
	}

	return all, &total, err
}

func (t *BankAccount) GetSupplierBankAccounts(id int) ([]string, error) {
	collection := Upper.Collection("bank_accounts")
	var bankAccounts []BankAccount
	var bankAccountValues []string
	
	res := collection.Find(up.Cond{"supplier_id": id})   

	err := res.All(&bankAccounts)
	if err != nil {
		return nil, err
	}

	for _, bankAccount := range bankAccounts {
		bankAccountValues = append(bankAccountValues, bankAccount.Title)
	}

	return bankAccountValues, nil
}

// Get gets one record from the database, by id, using upper
func (t *BankAccount) Get(id int) (*BankAccount, error) {
    var one BankAccount
    collection := Upper.Collection(t.Table())

    res := collection.Find(up.Cond{"id": id})
    err := res.One(&one)
    if err != nil {
        return nil, err
    }
    return &one, nil
}

// Update updates a record in the database, using upper
func (t *BankAccount) Update(m BankAccount) error {
    m.UpdatedAt = time.Now()
    collection := Upper.Collection(t.Table())
    res := collection.Find(m.ID)
    err := res.Update(&m)
    
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
					return fmt.Errorf("bankovni račun %s već postoji", m.Title)
			}
		}

    return nil
}

// Delete deletes a record from the database by id, using upper
func (t *BankAccount) Delete(tx up.Session, title string) error {
    collection := tx.Collection(t.Table())
    res := collection.Find("title", title)
    err := res.Delete()
    if err != nil {
        return err
    }
    return nil
}

// Insert inserts a model into the database, using upper
func (t *BankAccount) Insert(tx up.Session, m BankAccount) (int, error) {
    m.CreatedAt = time.Now()
    m.UpdatedAt = time.Now()
    collection := tx.Collection(t.Table())
    res, err := collection.Insert(m)

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
						return 0, fmt.Errorf("bankovni račun %s već postoji", m.Title)
				}
		}
		
    id := getInsertId(res.ID())

    return id, nil
}
