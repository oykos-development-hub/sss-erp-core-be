CREATE TABLE IF NOT EXISTS permissions (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    path TEXT NOT NULL,
    parent_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO permissions (title, path, parent_id, created_at, updated_at) 
VALUES 
    ('Moduli', '/', null, now(), now()),
    ('Finansije', '/finance', 1, now(), now()),
    ('Osnovna sredstva', '/inventory', 1, now(), now()),
    ('Javne nabavke', '/procurements', 1, now(), now()),
    ('Materijalno knjigovođstvo', '/accounting', 1, now(), now()),
    ('Ljudski resursi', '/hr', 1, now(), now()),
    ('Budžet', '/finance/budget', 2, now(), now()),
    ('Trenutni budžet', '/finance/current-budget', 2, now(), now()),
    ('Obaveze i potraživanja', '/finance/liabilities-receivables', 2, now(), now()),
    ('Izvještaji', '/finance/reports', 2, now(), now()),
    ('Pokretna sredstva', '/inventory/movable-inventory', 3, now(), now()),
    ('Nepokretna sredstva', '/inventory/immovable-inventory', 3, now(), now()),
    ('Sitan inventar', '/inventory/small-inventory', 3, now(), now()),
    ('Izvještaji', '/inventory/reports', 3, now(), now()),
    ('Planovi', '/procurements/plans', 4, now(), now()),
    ('Ugovori', '/procurements/contracts', 4, now(), now()),
    ('Izvještaji', '/procurements/reports', 4, now(), now()),
    ('Narudžbenica', '/accounting/order-form', 5, now(), now()),
    ('Ugovori', '/accounting/contracts', 5, now(), now()),
    ('Zalihe robe', '/accounting/stock', 5, now(), now()),
    ('Izvještaji', '/accounting/reports', 5, now(), now()),
    ('Kadrovi', '/hr/employees', 6, now(), now()),
    ('Sistematizacija', '/hr/systematization', 6, now(), now()),
    ('Sudije', '/hr/judges', 6, now(), now()),
    ('Oglasi', '/hr/job-tenders', 6, now(), now()),
    ('Preporuke interne revizije', '/hr/revision-recommendations', 6, now(), now()),
    ('Izvještaji', '/hr/reports', 6, now(), now());