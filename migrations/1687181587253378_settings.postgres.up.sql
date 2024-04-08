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

INSERT INTO settings(title, entity, abbreviation, value) VALUES (
  'Rješenje o prekidu radnog odnosa',
  'resolution_types',
  'PRO',
  'employment_termination'
);

INSERT INTO settings(title, entity, abbreviation, value) VALUES (
  'Rješenje o korišćenju I dijela godišnjeg odmora',
  'resolution_types',
  '1GO',
  'vacation_details'
);

INSERT INTO settings(title, entity, abbreviation, value) VALUES (
  'Rješenje o prekidu radnog odnosa',
  'resolution_types',
  'GO',
  'vacation'
);