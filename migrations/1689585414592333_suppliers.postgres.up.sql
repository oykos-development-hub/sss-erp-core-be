CREATE TABLE IF NOT EXISTS suppliers (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    abbreviation TEXT,
    official_id TEXT,
    tax_percentage FLOAT,
    address TEXT,
    description TEXT,
    entity TEXT DEFAULT 'supplier',
    folder_id INTEGER,
    parent_id INTEGER REFERENCES suppliers(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
