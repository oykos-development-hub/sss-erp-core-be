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

INSERT INTO roles_permissions (role_id, permission_id, can_create, can_update, can_read, can_delete)
SELECT roles.id, permissions.id, true, true, true, true
FROM roles
CROSS JOIN permissions;