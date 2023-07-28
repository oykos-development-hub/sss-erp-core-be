CREATE TABLE user_account_logs (
    id serial PRIMARY KEY,
    created_at TIMESTAMP,
    target_user_account_id INTEGER NOT NULL,
    source_user_account_id INTEGER NOT NULL,
    change_type INTEGER NOT NULL,
    previous_value JSONB NOT NULL,
    new_value JSONB NOT NULL,
    FOREIGN KEY (target_user_account_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (source_user_account_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (change_type) REFERENCES settings(id) ON UPDATE CASCADE ON DELETE CASCADE
);
