package data

import (
	"encoding/json"
	"time"

	up "github.com/upper/db/v4"
)

// UserAccountLog struct
type UserAccountLog struct {
	ID                  int             `db:"id,omitempty"`
	CreatedAt           time.Time       `db:"created_at,omitempty"`
	TargetUserAccountID int             `db:"target_user_account_id"`
	SourceUserAccountID int             `db:"source_user_account_id"`
	ChangeType          int             `db:"change_type"`
	PreviousValue       json.RawMessage `db:"previous_value"`
	NewValue            json.RawMessage `db:"new_value"`
}

// Table returns the table name
func (t *UserAccountLog) Table() string {
	return "user_account_logs"
}

// GetAll gets all records from the database, using upper
func (t *UserAccountLog) GetAll(condition *up.Cond) ([]*UserAccountLog, error) {
	collection := Upper.Collection(t.Table())
	var all []*UserAccountLog
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
func (t *UserAccountLog) Get(id int) (*UserAccountLog, error) {
	var one UserAccountLog
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *UserAccountLog) Update(m UserAccountLog) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *UserAccountLog) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *UserAccountLog) Insert(m UserAccountLog) (int, error) {
	m.CreatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}

// Builder is an example of using upper's sql builder
func (t *UserAccountLog) Builder(id int) ([]*UserAccountLog, error) {
	collection := Upper.Collection(t.Table())

	var result []*UserAccountLog

	err := collection.Session().
		SQL().
		SelectFrom(t.Table()).
		Where("id > ?", id).
		OrderBy("id").
		All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
