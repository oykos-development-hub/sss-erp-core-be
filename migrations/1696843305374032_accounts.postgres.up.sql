CREATE TABLE IF NOT EXISTS accounts (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    parent_id INTEGER,
    serial_number TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
);
