package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"
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
	collection := Upper.Collection(t.Table())
	var all []*Permission
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("id").All(&all)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper order")
	}

	return all, err
}

// GetAll gets all records from the database, using upper
func (p *Permission) GetAllPermissionOfRole(roleID int) ([]*PermissionWithRoles, error) {
	collection := Upper.Collection(p.Table())
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
		return nil, newErrors.Wrap(err, "upper exec query")
	}

	return all, err
}

func (p *Permission) GetUsersByPermission(operation string, title string) ([]User, error) {
	query1 := `select u.id from users u 
	           left join roles r on r.id = u.role_id 
			   left join roles_permissions rp on rp.role_id = r.id
			   left join permissions p on p.id = rp.permission_id
			   where p.path = $1 `

	var whereCond string

	if operation == "full_access" {
		whereCond = " and rp.can_create = true and rp.can_delete = true and rp.can_update = true and rp.can_read = true"
	} else if operation == "can_create" {
		whereCond = " and rp.can_create = true"
	} else if operation == "can_update" {
		whereCond = " and rp.can_update = true"
	} else if operation == "can_delete" {
		whereCond = " and rp.can_delete = true"
	} else if operation == "can_read" {
		whereCond = " and rp.can_read = true"
	}

	query := query1 + whereCond + ";"

	rows1, err := Upper.SQL().Query(query, title)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper exec")
	}
	defer rows1.Close()

	var users []User

	for rows1.Next() {
		var user User
		err = rows1.Scan(&user.ID)

		if err != nil {
			return nil, newErrors.Wrap(err, "upper scan")
		}

		users = append(users, user)

	}
	return users, nil
}

// Get gets one record from the database, by id, using upper
func (t *Permission) Get(id int) (*Permission, error) {
	var one Permission
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Permission) Update(m Permission) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Permission) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Permission) Insert(m Permission) (int, error) {
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
