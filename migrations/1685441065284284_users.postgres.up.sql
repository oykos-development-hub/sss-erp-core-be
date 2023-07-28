ALTER TABLE users
    ADD COLUMN secondary_email VARCHAR(255),
    ADD COLUMN phone VARCHAR(255) NOT NULL,
    ADD COLUMN pin VARCHAR(16) NOT NULL,
    ADD COLUMN active BOOLEAN DEFAULT true,
    ADD COLUMN verified_email BOOLEAN DEFAULT false,
    ADD COLUMN verified_phone BOOLEAN DEFAULT false,
    ADD COLUMN folder_id INTEGER;
