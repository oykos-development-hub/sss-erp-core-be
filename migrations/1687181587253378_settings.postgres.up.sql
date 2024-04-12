CREATE TABLE settings (
    id serial PRIMARY KEY,
    title text NOT NULL,
    entity text NOT NULL,
    abbreviation text NOT NULL,
    description text,
    value text,
    color text,
    icon text,
    parent_id integer REFERENCES settings(id) ON DELETE CASCADE,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);
