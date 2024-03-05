CREATE TABLE IF NOT EXISTS bank_accounts (
    id serial PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
