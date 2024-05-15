CREATE TABLE IF NOT EXISTS accounts (
    id serial PRIMARY KEY,
    title TEXT NOT NULL,
    parent_id INTEGER,
    serial_number TEXT,
    version INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
);


CREATE OR REPLACE FUNCTION assign_parent_id()
RETURNS TRIGGER AS $$
DECLARE
    parent_serial varchar;
    parent_id int;
BEGIN
    -- Izračunavanje osnovnog serial_number-a sa jednom cifrom manje
    parent_serial := SUBSTRING(NEW.serial_number FROM 1 FOR LENGTH(NEW.serial_number) - 1);

    -- Pronalaženje parent_id na osnovu izračunatog parent_serial
    SELECT id INTO parent_id FROM accounts WHERE serial_number = parent_serial ORDER BY id DESC LIMIT 1;

    -- Postavljanje parent_id za novi unos ako je pronađen odgovarajući parent
    IF parent_id IS NOT NULL THEN
        NEW.parent_id := parent_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Kreiranje trigger-a koji se aktivira pri svakom insertu
CREATE TRIGGER assign_parent_id_trigger
BEFORE INSERT ON accounts
FOR EACH ROW
EXECUTE FUNCTION assign_parent_id();


INSERT INTO accounts (title, serial_number, created_at, updated_at, version) VALUES
('Izdaci', 4, now(), now(), 1),
('Tekući izdaci', 41, now(), now(), 1),
('Transferi institucijama, pojedincima, nevladinom i javnom sektoru', 43, now(), now(), 1),
('Kapitalni izdaci', 44, now(), now(), 1),
('Otplata dugova', 46, now(), now(), 1),
('Bruto zarade: doprinosi na teret poslodavca', 411, now(), now(), 1),
('Ostala lična primanja', 412, now(), now(), 1),
('Rashodi za materijal', 413, now(), now(), 1),
('Rashodi za usluge', 414, now(), now(), 1),
('Rashodi za tekuće održavanje', 415, now(), now(), 1),
('Renta', 417, now(), now(), 1),
('Ostali izdaci', 419, now(), now(), 1),
('Transferi institucijama pojedincima, nevladinom i javnom sektoru', 431, now(), now(), 1),
('Otplata obaveza iz prethodnog perioda', 463, now(), now(), 1),
('Neto zarade', 4111, now(), now(), 1),
('Porez na zarade', 4112, now(), now(), 1),
('Doprinosi na teret zaposlenog', 4113, now(), now(), 1),
('Doprinosi na teret poslodavca', 4114, now(), now(), 1),
('Opštinski prirez', 4115, now(), now(), 1),
('Jubilarne nagrade', 4124, now(), now(), 1),
('Otpremnine', 4125, now(), now(), 1),
('Ostale naknade', 4127, now(), now(), 1),
('Administrativni materijal', 4131, now(), now(), 1),
('Materijal za posebne namjene', 4133, now(), now(), 1),
('Rashodi za energiju', 4134, now(), now(), 1),
('Rashodi za gorivo', 4135, now(), now(), 1),
('Ostali rashodi za materijal', 4139, now(), now(), 1),
('Službena putovanja', 4141, now(), now(), 1),
('Reprezentacija', 4142, now(), now(), 1),
('Komunikacione usluge', 4143, now(), now(), 1),
('Advokatske, notarske i pravne usluge', 4146, now(), now(), 1),
('Usluge stručnog usavršavanja', 4148, now(), now(), 1),
('Ostale usluge', 4149, now(), now(), 1),
('Tekuće održavanje građevinskih objekata', 4152, now(), now(), 1),
('Tekuće održavanje opreme', 4153, now(), now(), 1),
('Zakup objekata', 4171, now(), now(), 1),
('Izdaci po osnovu isplate ugovora o djelu', 4191, now(), now(), 1),
('Izdaci po osnovu troškova sudskih postupaka', 4192, now(), now(), 1),
('Osiguranje', 4194, now(), now(), 1),
('Ostalo', 4199, now(), now(), 1),
('Ostali transferi pojedincima', 4318, now(), now(), 1),
('Izdaci za opremu', 4415, now(), now(), 1),
('Otplata obaveza iz prethodnog perioda', 4630, now(), now(), 1),
('Kancelarijski materijal', 41311, now(), now(), 1),
('Sitan inventar', 41312, now(), now(), 1),
('Sredstva higijene', 41313, now(), now(), 1),
('Rezervni djelovi', 41314, now(), now(), 1),
('Radna odjeća', 41315, now(), now(), 1),
('Publikacije, časopisi i glasila', 41335, now(), now(), 1),
('Ostalo', 41337, now(), now(), 1),
('Rashodi za električnu energiju', 41341, now(), now(), 1),
('Ostali rashodi za energiju', 41342, now(), now(), 1),
('Rashodi za tečna goriva (dizel, benzin, mazut)', 41351, now(), now(), 1),
('Rashodi za gas', 41352, now(), now(), 1),
('Rashodi za čvrsto gorivo (drvo, ugalj)', 41353, now(), now(), 1),
('Sredstva transporta', 44151, now(), now(), 1),
('Kompjuterska oprema', 44152, now(), now(), 1),
('Kancelarijska oprema', 44153, now(), now(), 1),
('Telekomunikaciona oprema', 44154, now(), now(), 1);

