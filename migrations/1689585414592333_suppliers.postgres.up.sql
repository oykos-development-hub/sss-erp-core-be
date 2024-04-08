CREATE TABLE IF NOT EXISTS suppliers (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    abbreviation TEXT,
    official_id TEXT,
    tax_percentage INTEGER,
    address TEXT,
    description TEXT,
    entity TEXT DEFAULT 'supplier',
    folder_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
