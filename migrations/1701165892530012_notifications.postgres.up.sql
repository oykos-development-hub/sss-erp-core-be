CREATE TABLE IF NOT EXISTS notifications (
    id serial PRIMARY KEY,
    content TEXT NOT NULL,
    module TEXT NOT NULL,
    from_user_id INTEGER,
    to_user_id INTEGER NOT NULL,
    from_content TEXT NOT NULL,
    is_read bool NOT NULL,
    data JSON,
    path TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
