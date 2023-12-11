CREATE TABLE IF NOT EXISTS roles_permissions (
    id serial PRIMARY KEY,
    permission_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    can_create BOOLEAN NOT NULL,
    can_read BOOLEAN NOT NULL,
    can_update BOOLEAN NOT NULL,
    can_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
