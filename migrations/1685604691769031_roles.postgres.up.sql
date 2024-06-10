-- Drop table if exists
drop table if exists roles cascade;

-- Create roles table
CREATE TABLE roles (
  id serial PRIMARY KEY,
  title TEXT NOT NULL,
  abbreviation TEXT NOT NULL,
  active BOOLEAN NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE users
  ADD COLUMN role_id integer,
  ADD CONSTRAINT users_role_id_fk FOREIGN KEY (role_id) REFERENCES roles(id);

  INSERT INTO roles (title, abbreviation, created_at, updated_at, active) VALUES (
    'Admin', 'ADM', '2024-02-05 13:03:53.244118', '2024-02-05 13:03:53.24412', true);

INSERT INTO users (first_name, last_name, user_active, email, password, created_at, 
    updated_at, secondary_email, phone, pin, active, verified_email, verified_phone, folder_id, role_id
) VALUES (
    'Admin', 'Admin', 0, 'admin@example.com', 
    '$2a$12$yqHOCXa8DtSO6NA6mksuIeQ8LVnnW6bAKaJnbaHwwZAeODjivbsce', 
    '2023-08-08 12:00:00', '2023-11-29 13:14:12.77536', 
    '', '382-67-1234567', 1234, true, true, false, 1, 1
);
