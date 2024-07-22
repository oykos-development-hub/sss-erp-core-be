CREATE TABLE IF NOT EXISTS customer_supports (
    id serial PRIMARY KEY,
    user_documentation_file_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO customer_supports(id, user_documentation_file_id) 
VALUES(1, 0);
