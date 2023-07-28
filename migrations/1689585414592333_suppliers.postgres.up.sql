CREATE TABLE IF NOT EXISTS suppliers (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    abbreviation TEXT,
    official_id TEXT,
    address TEXT,
    description TEXT,
    folder_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
