-- Admin and user roles
INSERT INTO roles (
    id, title, abbreviation, color, icon, created_at, updated_at
) VALUES
    (1, 'Admin', 'ADM', '#FF0000', 'admin_icon.png', '2023-08-08 10:00:00', '2023-08-08 10:00:00'),
    (2, 'User', 'USR', '#0000FF', 'user_icon.png', '2023-08-08 11:00:00', '2023-08-08 11:00:00');

-- 5 users (1 admin)
INSERT INTO users (
    id, first_name, last_name, email, secondary_email, active, password, pin, phone, verified_email, verified_phone, folder_id, created_at, updated_at, role_id
) VALUES
    (1, 'Admin', 'Admin', 'admin@example.com', NULL, TRUE, 'admin123', '1234', '382-67-1234567', TRUE, FALSE, 1, '2023-08-08 12:00:00', '2023-08-08 12:00:00', 1),
    (2, 'Marko', 'Radulović', 'zaposleni1@example.com', 'marko.radulovic2@example.com', TRUE, 'zaposleni123', '1234', '382-67-7654321', TRUE, TRUE, 2, '2023-08-08 13:00:00', '2023-08-08 13:00:00', 2),
    (3, 'Milica', 'Petrović', 'zaposleni2@example.com', NULL, TRUE, 'zaposleni123', '1234', '382-67-1011121', TRUE, FALSE, 3, '2023-08-08 14:00:00', '2023-08-08 14:00:00', 2),
    (4, 'Nikola', 'Ivanović', 'zaposleni3@example.com', 'nikola.ivanovic2@example.com', TRUE, 'zaposleni123', '1234', '382-67-1213141', TRUE, TRUE, 4, '2023-08-08 15:00:00', '2023-08-08 15:00:00', 2),
    (5, 'Ana', 'Djurović', 'zaposleni4@example.com', NULL, TRUE, 'zaposleni123', '1234', '382-67-1415161', TRUE, FALSE, NULL, '2023-08-08 16:00:00', '2023-08-08 16:00:00', 2),
    (6, 'Igor', 'Milošević', 'igor.milosevic@example.com', NULL, TRUE, 'zaposleni123', '1234', '382-67-1617181', TRUE, FALSE, NULL, '2023-08-08 17:00:00', '2023-08-08 17:00:00', 2);

-- engagements for 5 users
INSERT INTO settings 
    (id, title, abbreviation, entity, description, value, color, icon, created_at, updated_at) 
VALUES
    (1, 'Sudija', 'SUD', 'engagement_types', 'Angažman kao sudija u pravosudnom sistemu.', 'sudija', '#D32F2F', 'judge_icon.png', '2023-08-08 17:00:00', '2023-08-08 17:00:00'),
    (2, 'Tužilac', 'TUZ', 'engagement_types', 'Angažman kao tužilac u pravosudnom sistemu.', 'tuzilac', '#1976D2', 'prosecutor_icon.png', '2023-08-08 18:00:00', '2023-08-08 18:00:00'),
    (3, 'Advokat', 'ADV', 'engagement_types', 'Angažman kao advokat u pravosudnom sistemu.', 'advokat', '#388E3C', 'lawyer_icon.png', '2023-08-08 19:00:00', '2023-08-08 19:00:00'),
    (4, 'Pravni Sekretar', 'PSK', 'engagement_types', 'Angažman kao pravni sekretar u pravosudnom sistemu.', 'pravni_sekretar', '#8E24AA', 'legal_secretary_icon.png', '2023-08-08 21:00:00', '2023-08-08 21:00:00'),
    (5, 'IT Stručnjak', 'ITS', 'engagement_types', 'Angažman kao IT stručnjak u pravosudnom sistemu.', 'it_strucnjak', '#7B1FA2', 'it_expert_icon.png', '2023-08-08 22:00:00', '2023-08-08 22:00:00');

-- contract types
INSERT INTO settings 
    (id, title, abbreviation, entity, description, value, color, icon, created_at, updated_at) 
VALUES
    (6,'Ugovor na odredjeno vrijeme', 'UOV', 'contract_types', 'Ugovor na odredjeno vrijeme', NULL, NULL, NULL, NOW(), NOW()),
    (7,'Ugovor na neodredjeno vrijeme', 'UNV', 'contract_types', 'Ugovor o radu', NULL, NULL, NULL, NOW(), NOW());


-- resolution types
INSERT INTO settings (
    id, title, abbreviation, entity, description, value, color, icon, created_at, updated_at
) VALUES
    (8, 'Rešenje', 'RŠ', 'resolution_types', 'Rešavanje problema na konkretnom radnom mestu.', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW()),
	(9, 'Odluka', 'OD', 'resolution_types', 'Donošenje odluke o promeni radnog mesta, povećanju plata.', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW()),
	(10, 'Resavanje problema', 'RP', 'resolution_types', 'Rešavanje problema na radnom mestu.', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW());

-- education types
INSERT INTO settings (
    id, title, abbreviation, entity, description, created_at, updated_at)
VALUES
    (11, 'Nivo 1 - Osnovna stručna osposobljenost', '1', 'education_types', 'opis', NOW(), NOW()),
    (12, 'Nivo 2 - Srednja stručna osposobljenost', '2', 'education_types', 'opis', NOW(), NOW()),
    (13, 'Nivo 3 - Srednja stručna osposobljenost s dodatnim obrazovanjem', '3', 'education_types', 'opis', NOW(), NOW()),
    (14, 'Nivo 4 - Visoka stručna sprema', '4', 'education_types', 'opis', NOW(), NOW()),
    (15, 'Nivo 5 - Diplomske studije (prva i druga godina)', '5', 'education_types', 'opis', NOW(), NOW()),
    (16, 'Nivo 6 - Bachelor diploma', '6', 'education_types', 'opis', NOW(), NOW()),
    (17, 'Nivo 7 - Postdiplomski studiji', '7', 'education_types', 'opis', NOW(), NOW());

INSERT INTO settings (
    id, title, abbreviation, entity, description, value, color, icon, created_at, updated_at
) VALUES
    (18, 'Dobar', 'A', 'evaluation_types', 'Izvrsni rezultati', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW()),
	(19, 'Zadovoljio', 'B', 'evaluation_types', 'Dobri rezultati', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW()),
	(20, 'Nije zadovoljio', 'C', 'evaluation_types', 'Losi rezultati', NULL, '#006699', 'fa fa-file-text-o', NOW(), NOW());

--suppliers
INSERT INTO suppliers (
    id,	title, abbreviation, official_id, address, description, folder_id, created_at,updated_at)
VALUES
    (1, 'Namos', 'S1', '123456789', 'Address 1', 'Namos Dajkovic Company', 101, NOW(), NOW()),
    (2,'EPCG', 'S2', '987654321', 'Address 2', 'Elektroprivreda', 102, NOW(), NOW()),
    (3,'Telekom', 'S3', '456789123', 'Address 3', 'Telekom Crna Gora', 103, NOW(), NOW());