CREATE TABLE IF NOT EXISTS list_of_parameters (
    id serial PRIMARY KEY,
    title VARCHAR ( 255 ) NOT NULL,
    description text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO list_of_parameters (title, description, created_at, updated_at) 
VALUES 
('tekuci_datum', 'Trenutni datum.', NOW(), NOW()),
('tekuca_godina', 'Trenutna godina.', NOW(), NOW()),
('tekuci_mjesec', 'Trenutni mjesec.', NOW(), NOW()),
('ime_prezime', 'Ime i prezime zaposlenog.', NOW(), NOW()),
('jmbg', 'Jedinstveni matični broj zaposlenog.', NOW(), NOW()),
('ulica', 'Adresa ulice zaposlenog.', NOW(), NOW()),
('organizaciona_jedinica', 'Organizaciona jedinica.', NOW(), NOW()),
('odjeljenje', 'Odjeljenje unutar organizacije.', NOW(), NOW()),
('radno_mjesto', 'Radno mjesto.', NOW(), NOW()),
('radno_mjesto_uslovi', 'Uslovi za radno mjesto zaposlenog.', NOW(), NOW()),
('broj_sistematizacije', 'Broj trenutno aktivne sistematizacije.', NOW(), NOW()),
('datum_sistematizacije', 'Datum donošenja trenutno aktivne sistematizacije.', NOW(), NOW()),
('datum_pocetka_ugovora', 'Datum početka posljednjeg aktivnog ugovora ili aneksa.', NOW(), NOW()),
('datum_pocetka_rada', 'Datum početka rada.', NOW(), NOW()),
('datum_isteka_ugovora', 'Datum isteka ugovora.', NOW(), NOW()),
('trajanje_ugovora_u_danima', 'Trajanje ugovora u danima.', NOW(), NOW()),
('steceni_broj_dana_odmora', 'Stečeni broj dana godišnjeg odmora.', NOW(), NOW()),
('preostali_broj_dana_odmora', 'Preostali broj dana godišnjeg odmora.', NOW(), NOW()),
('potroseni_broj_dana_odmora', 'Potrošeni broj dana godišnjeg odmora.', NOW(), NOW()),
('radni_sati_sedmicno', 'Broj radnih sati sedmično.', NOW(), NOW()),
('datum_pocetka_godisnjeg_odmora', 'Datum početka godišnjeg odmora.', NOW(), NOW()),
('datum_kraja_godisnjeg_odmora', 'Datum kraja godišnjeg odmora.', NOW(), NOW()),
('ocjena',  'Posljednja aktuelna ocjena zaposlenog.', NOW(), NOW()),
('obrazovanje', 'Nivo obrazovanja zaposlenog.', NOW(), NOW());
