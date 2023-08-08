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
    (5, 'Ana', 'Djurović', 'zaposleni4@example.com', NULL, TRUE, 'zaposleni123', '1234', '382-67-1415161', TRUE, FALSE, NULL, '2023-08-08 16:00:00', '2023-08-08 16:00:00', 2);
    (6, 'Igor', 'Milošević', 'igor.milosevic@example.com', NULL, TRUE, 'zaposleni123', '1234', '382-67-1617181', TRUE, FALSE, NULL, '2023-08-08 17:00:00', '2023-08-08 17:00:00', 2); -- Adjust the role_id as needed

-- engagements for 5 users
INSERT INTO settings (
    id, title, abbreviation, entity, description, value, color, icon, created_at, updated_at
) VALUES
    (1, 'Sudija', 'SUD', 'engagement_types', 'Angažman kao sudija u pravosudnom sistemu.', 'sudija', '#D32F2F', 'judge_icon.png', '2023-08-08 17:00:00', '2023-08-08 17:00:00'),
    (2, 'Tužilac', 'TUZ', 'engagement_types', 'Angažman kao tužilac u pravosudnom sistemu.', 'tuzilac', '#1976D2', 'prosecutor_icon.png', '2023-08-08 18:00:00', '2023-08-08 18:00:00'),
    (3, 'Advokat', 'ADV', 'engagement_types', 'Angažman kao advokat u pravosudnom sistemu.', 'advokat', '#388E3C', 'lawyer_icon.png', '2023-08-08 19:00:00', '2023-08-08 19:00:00'),
    (4, 'Pravni Sekretar', 'PSK', 'engagement_types', 'Angažman kao pravni sekretar u pravosudnom sistemu.', 'pravni_sekretar', '#8E24AA', 'legal_secretary_icon.png', '2023-08-08 21:00:00', '2023-08-08 21:00:00');
    (5, 'IT Stručnjak', 'ITS', 'engagement_types', 'Angažman kao IT stručnjak u pravosudnom sistemu.', 'it_strucnjak', '#7B1FA2', 'it_expert_icon.png', '2023-08-08 22:00:00', '2023-08-08 22:00:00');

INSERT INTO settings (
    id, title, abbreviation, entity, description, created_at, updated_at)
VALUES
    ('Bachelor of Law', 'LLB', 'education_types', 'A basic degree in law and legal studies', NOW(), NOW()),
    ('Master of Law', 'LLM', 'education_types', 'An advanced, postgraduate academic degree in law', NOW(), NOW()),
    ('Doctor of Law', 'JD', 'education_types', 'A professional doctorate and first professional graduate degree in law', NOW(), NOW()),
    ('Master of Criminal Justice', 'MCJ', 'education_types', 'A postgraduate degree focusing on the study of criminal justice and criminology', NOW(), NOW()),
    ('Paralegal Certificate', 'PC', 'education_types', 'A certificate program in legal studies for paralegals', NOW(), NOW()),
    ('Ph.D. in Legal Studies', 'PhD', 'education_types', 'A doctoral degree focusing on research and academic approach to law and legal studies', NOW(), NOW());
