package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Permission struct
type Permission struct {
	ID        int       `db:"id,omitempty"`
	Title     string    `db:"title"`
	Path      string    `db:"path"`
	ParentID  *int      `db:"parent_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PermissionWithRoles struct {
	Permission
	CanCreate bool `db:"can_create"`
	CanRead   bool `db:"can_read"`
	CanUpdate bool `db:"can_update"`
	CanDelete bool `db:"can_delete"`
}

// Table returns the table name
func (t *Permission) Table() string {
	return "permissions"
}

// GetAll gets all records from the database, using upper
func (t *Permission) GetAll(condition *up.Cond) ([]*Permission, error) {
	collection := upper.Collection(t.Table())
	var all []*Permission
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("id").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// GetAll gets all records from the database, using upper
func (p *Permission) GetAllPermissionOfRole(roleID int) ([]*PermissionWithRoles, error) {
	collection := upper.Collection(p.Table())
	var all []*PermissionWithRoles

	err := collection.Session().
		SQL().
		Select(
			"p.*",
			up.Raw("COALESCE(rp.can_read, FALSE) AS can_read"),
			up.Raw("COALESCE(rp.can_delete, FALSE) AS can_delete"),
			up.Raw("COALESCE(rp.can_update, FALSE) AS can_update"),
			up.Raw("COALESCE(rp.can_create, FALSE) AS can_create"),
		).
		From("permissions p").
		LeftJoin("roles_permissions rp").On("p.id = rp.permission_id AND rp.role_id = ?", roleID).
		OrderBy("id").
		All(&all)

	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *Permission) Get(id int) (*Permission, error) {
	var one Permission
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Permission) Update(m Permission) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Permission) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Permission) Insert(m Permission) (int, error) {
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
