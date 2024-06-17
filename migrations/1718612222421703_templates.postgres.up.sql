CREATE TABLE IF NOT EXISTS templates (
    id serial PRIMARY KEY,
    title VARCHAR ( 255 ) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
