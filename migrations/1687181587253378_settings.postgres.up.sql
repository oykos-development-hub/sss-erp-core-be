CREATE TABLE settings (
    id serial PRIMARY KEY,
    title text NOT NULL,
    entity text NOT NULL,
    abbreviation text NOT NULL,
    description text,
    value text,
    color text,
    icon text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

-- add auto update of updated_at. If you already have this trigger
-- you can delete the next 7 lines
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON settings
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();