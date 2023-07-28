DROP TABLE if exists roles cascade;

ALTER TABLE users
  DROP CONSTRAINT users_role_id_fk,
  DROP COLUMN role_id;
