CREATE TABLE IF NOT EXISTS permissions (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    path TEXT NOT NULL,
    parent_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
