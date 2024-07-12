INSERT INTO roles (title, abbreviation, created_at, updated_at, active) VALUES (
    'Menadžer OJ', 'MOJ', NOW(), NOW(), true);

    INSERT INTO roles_permissions (
    permission_id, role_id, can_create, can_read, can_update, can_delete, created_at, updated_at
)
SELECT 
    p.id AS permission_id, 
    2 AS role_id,  --dodajem sve permisije za usera id 2
    TRUE AS can_create, 
    TRUE AS can_read, 
    TRUE AS can_update, 
    TRUE AS can_delete,
    NOW() AS created_at,
    NOW() AS updated_at
FROM permissions p
WHERE NOT EXISTS (
    SELECT 1 
    FROM roles_permissions rp 
    WHERE rp.permission_id = p.id AND rp.role_id = 2
);

INSERT INTO users (
    first_name, last_name, user_active, email, password, created_at, updated_at, 
    secondary_email, phone, pin, active, verified_email, verified_phone, folder_id, role_id
) VALUES
('Menadžer', 'Apelacioni Sud Crne Gore', 0, 'menadzerascg@example.com', '$2a$12$yqHOCXa8DtSO6NA6mksuIeQ8LVnnW6bAKaJnbaHwwZAeODjivbsce', NOW(), NOW(), 'sadio.mane@gmail.com', '+382 68 878 434', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Sekretarijat Sudskog savjeta', 0, 'menadzersss@example.com', '$2a$12$yqHOCXa8DtSO6NA6mksuIeQ8LVnnW6bAKaJnbaHwwZAeODjivbsce',  NOW(), NOW(), 'vulja22@example.com', '555-555-5555', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Viši sud Podgorica', 0, 'menadzervspg@example.com', '$2a$12$yqHOCXa8DtSO6NA6mksuIeQ8LVnnW6bAKaJnbaHwwZAeODjivbsce',  NOW(), NOW(), 'rafa@gmail.com', '+382 68 878 434', '3424', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Žabljak', 0, 'menadzeroszabljak@example.com', '$2a$12$yqHOCXa8DtSO6NA6mksuIeQ8LVnnW6bAKaJnbaHwwZAeODjivbsce',  NOW(), NOW(), 'lebron.james@gmail.me', '+382 68 878 434', '3242', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Upravni sud Crne Gore', 0, 'menadzeruscg@example.com', '$2a$12$vJA1hCeFvy/4JXXMrEu4BeTpAf.rkmsilh.4lt99tTOFUpq5RAEz.',  NOW(), NOW(), 'menadzeruscg1@example.com', '068954789', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Plav', 0, 'menadzerosplav@example.com', '$2a$12$slPNPmU56yh./ndiss76H.xKmvxFlItVy0yvSn5IDgz95LWhJPxYy',  NOW(), NOW(), 'menadzerosplav1@example.com', '068954821', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Sud za prekršaje Bijelo Polje', 0, 'menadzerszpbp@example.com', '$2a$12$STdPzi7OWaMedvM40EVfZOB7fqRnXupBlEMrDeWRwTCvXiCxIO2Hy',  NOW(), NOW(), 'menadzerszpbp1@example.com', '069874512', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Berane', 0, 'menadzerosberane@example.com', '$2a$12$xxC8vxjorgaFoepx01rJ..tX6m7J0F3/Op957GOmnvHTRW3qeQNXC',  NOW(), NOW(), 'menadzerosberane1@example.com', '067842659', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Cetinje', 0, 'menadzeroscetinje@example.com', '$2a$12$UhxarTfdixdidPFdqyj1MODA9e4B8TdwGGA0ugjSF7JPuTSj15Loy',  NOW(), NOW(), 'menadzeroscetinje1@example.com', '069265842', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Danilovgrad', 0, 'menadzerosdanilovgrad@example.com', '$2a$12$1A4K.QPgy2M1Q/OBUfJ7betdJoHg0I7gb0rB754m7uGx/BM7RZwvG',  NOW(), NOW(), 'menadzerosdanilovgrad1@example.com', '069547856', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Herceg Novi', 0, 'menadzeroshercegnovi@example.com', '$2a$12$s1KQq4foB9r/COmztiY9VecbNefEtLXXv7/bv0oqRiqaq8WZqRlVG',  NOW(), NOW(), 'menadzeroshercegnovi@oykos.com', '32424424', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Kolašin', 0, 'menadzeroskolasin@example.com', '$2a$12$5apmABVYTuR9c9jrI.8XAesrnN9ELuVzgNuynQBsQzksj..veMZyy',  NOW(), NOW(), 'menadzeroskolasin1@example.com', '068745962', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Nikšić', 0, 'menadzerosniksic@example.com', '$2a$12$51X1D3y6Lip8zF/dH7LyvOHbCrJ218jzdHC6RV6jwVaFD2KK19SHu',  NOW(), NOW(), 'menadzerosniksic1@example.com', '067623589', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Rožaje', 0, 'menadzerosrozaje@example.com', '$2a$12$kUt2.JtQ9w4kW9enyYKLDe/eaYQQZnxPDwxOK6PEqfSAZGD9BWtHW',  NOW(), NOW(), 'menadzerosrozaje1@example.com', '068562321', '1234', 't', 'f', 'f', NULL,2),
('Menadžer', 'Osnovni sud Ulcinj', 0, 'menadzerosulcinj@example.com', '$2a$12$bLX5PrJI7c2c9ceAEJxpV.QRRm.RwS2L4Gtn1099dZEfR21Pqx1OG',  NOW(), NOW(), 'menadzerosulcinj1@example.com', '068549654', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Bar', 0, 'menadzerosbar@example.com', '$2a$12$ICuyw6lpOuoA1ZpSNU35B.7oxRHs9a.OtiEeNlNPi.6p5sWZfJFUK',  NOW(), NOW(), 'menadzerosbar1@example.com', '068745', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Bijelo Polje', 0, 'menadzerosbijelopolje@example.com', '$2a$12$DxXfbGqQGARShcdWibfuievKxtUnpzXmaVO2BlVpScLU8g2HdDq5G',  NOW(), NOW(), 'menadzerosbijelopolje1@example.com', '068478965', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Pljevlja', 0, 'menadzerospljevlja@example.com', '$2a$12$Gs7LJKmco1I6dvAW1IWvCeJQCEbFTbrFKfhcMRoxjHoWRlK/xfklO',  NOW(), NOW(), 'menadzerospljevlja1@example.com', '068745892', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Viši sud Bijelo Polje', 0, 'menadzervsbp@example.com', '$2a$12$0leCPKOR4l6zvzEj5Vy/weUtursd51xZ/hAoDx4Vk9vvqqOU57Mbe',  NOW(), NOW(), 'menadzervsbp1@example.com', '068593154', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Podgorica', 0, 'menadzerospodgorica@example.com', '$2a$12$FeasSD0pRMZFET07Rlb27uTVn9OVV.ah/3HqMQgAtFJK756dntnwS',  NOW(), NOW(), 'menadzer55@example.me', '069783874', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Privredni sud Crne Gore', 0, 'menadzerpscg@example.com', '$2a$12$kIAKoC2uqBxDaYKG48iBUeIOM5BzENSik35XvDm.c.nwlMySALUzq',  NOW(), NOW(), 'menadzerpscg1@example.com', '068587146', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Sud za prekršaje Budva', 0, 'menadzerszpbd@example.com', '$2a$12$j5wE7LJTmqr8qYSTe.87E.ClbIlAoI4DZ2NRgxuZsLZRDwyuOgcxW',  NOW(), NOW(), 'menadzerszpbd1@example.com', '068457895', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Sud za prekršaje u Podgorici', 0, 'menadzerszppg@example.com', '$2a$12$vZHaoD1oDtiYiUDw5xOHZ.KOUVShavZ2bPWOUc4XRR9/t9Y1WkBou',  NOW(), NOW(), 'menadzerszppg1@example.com', '068457411', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Viši sud za prekršaje Crne Gore', 0, 'menadzervszpcg@example.com', '$2a$12$EVozgkons8DkkTAlNJSeGOkn3KnNLQbwkhYfB21zA8ensHKa7C89K',  NOW(), NOW(), 'menadzervszpcg1@example.com', '067458857', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Vrhovni sud Crne Gore', 0, 'menadzervscg@example.com', '$2a$12$M99rGJDDkSxZGZJmgUFWK.C5QdYv5eTPFg.k.LNzgzL5HQPPVH7Tm',  NOW(), NOW(), 'menadzervscg1@example.com', '069589874', '1234', 't', 'f', 'f', NULL, 2),
('Menadžer', 'Osnovni sud Kotor', 0, 'menadzeroskotor@example.com', '$2a$12$ixun8JgutpAFvew/TfVLlOIcUzVwBlK6K9nO5MvlyjBWmJWtzkEm2',  NOW(), NOW(), 'menadzeroskotor1@example.com', '069587485', '1234', 't', 'f', 'f', NULL, 2);
