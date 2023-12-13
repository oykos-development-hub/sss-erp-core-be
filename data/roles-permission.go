package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// RolesPermission struct
type RolesPermission struct {
	ID           int       `db:"id,omitempty"`
	PermissionID int       `db:"permission_id"`
	RoleID       int       `db:"role_id"`
	CanCreate    bool      `db:"can_create"`
	CanRead      bool      `db:"can_read"`
	CanUpdate    bool      `db:"can_update"`
	CanDelete    bool      `db:"can_delete"`
	CreatedAt    time.Time `db:"created_at,omitempty"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *RolesPermission) Table() string {
	return "roles_permissions"
}

// GetAll gets all records from the database, using upper
func (t *RolesPermission) GetAll(condition *up.Cond) ([]*RolesPermission, error) {
	collection := upper.Collection(t.Table())
	var all []*RolesPermission
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
func (t *RolesPermission) Get(id int) (*RolesPermission, error) {
	var one RolesPermission
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *RolesPermission) Update(m RolesPermission) error {
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
func (t *RolesPermission) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *RolesPermission) DeleteAllPermissionsByRole(roleID int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"role_id": roleID})
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *RolesPermission) Insert(m RolesPermission) (int, error) {
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
