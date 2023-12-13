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
