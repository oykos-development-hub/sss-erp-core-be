CREATE TABLE IF NOT EXISTS roles_permissions (
    id serial PRIMARY KEY,
    permission_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    can_create BOOLEAN NOT NULL,
    can_read BOOLEAN NOT NULL,
    can_update BOOLEAN NOT NULL,
    can_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO roles_permissions (
    permission_id, role_id, can_create, can_read, can_update, can_delete, created_at, updated_at
)
SELECT 
    p.id AS permission_id, 
    r.id AS role_id, 
    TRUE AS can_create, 
    TRUE AS can_read, 
    TRUE AS can_update, 
    TRUE AS can_delete,
    NOW() AS created_at,
    NOW() AS updated_at
FROM roles r, permissions p;
