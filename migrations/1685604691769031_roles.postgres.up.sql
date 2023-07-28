-- Drop table if exists
drop table if exists roles cascade;

-- Create roles table
CREATE TABLE roles (
  id serial PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  abbreviation VARCHAR(8) NOT NULL,
  color VARCHAR(32),
  icon text,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now()
);

ALTER TABLE users
  ADD COLUMN role_id integer,
  ADD CONSTRAINT users_role_id_fk FOREIGN KEY (role_id) REFERENCES roles(id);
