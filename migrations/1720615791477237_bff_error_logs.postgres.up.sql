CREATE TABLE IF NOT EXISTS bff_error_logs (
    id serial PRIMARY KEY,
    error VARCHAR ( 255 ) NOT NULL,
    code INTEGER,
    entity TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
