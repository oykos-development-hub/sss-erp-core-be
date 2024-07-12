INSERT INTO suppliers (title, abbreviation, official_id, address, description, folder_id, created_at, updated_at, entity, tax_percentage, parent_id) VALUES
--institutions
--ostaviti poresku upravu da ima id 1 zbog migracije sifarnika poreske uprave
( 'Poreska uprava Crne Gore', 'PUCG', '00000000', 'Bulevar Šarla de Gola br. 2', '', 0, '2024-04-22 08:44:05.098391', '2024-06-18 13:43:26.165354', 'institution', 0, NULL),
( 'Centar za alternativno rješavanje sporova', 'CZARS', '03250987', 'Serdara Jola Piletića bb', '', 0, '2024-04-22 11:52:53.2502', '2024-05-26 16:52:01.303355', 'institution', 0, NULL),
( 'Komisija za zaštitu prava u postupcima javnih nabavki', 'KZZPUPJN', '', 'Novaka Miloševa 28', '020 510-402', 0, '2024-04-22 12:02:01.82697', '2024-04-22 12:02:01.82697', 'institution', 0, NULL),
( 'Agencija za mirno rješavanje radnih sporova', 'AZMRRS', '', 'Serdara Jola Piletića bb', '+382 20 676523', 0, '2024-04-22 11:51:39.873929', '2024-04-22 11:54:57.913617', 'institution', 0, NULL),

--suppliers
('Namos', 'S1', '123456789', 'Address 1', 'Namos Dajkovic Company', 101, '2023-08-09 13:47:25.472372', '2024-03-07 12:04:40.443107', 'supplier', 0, NULL),
('EPCG', 'S2', '987654321', 'Address 2', 'Elektroprivreda', 102, '2023-08-09 13:47:25.472372', '2024-04-03 13:40:04.126075', 'supplier', 0, NULL),
('Telekom', 'S3', '456789123', 'Address 3', 'Telekom Crna Gora', 103, '2023-08-09 13:47:25.472372', '2024-04-16 14:26:40.293693', 'supplier', 0, NULL),
('Kastex DOO', 'S4', '02816667', 'Podgorica bb', '', 0, '0001-01-01 00:00:00', '2024-05-26 17:35:25.264943', 'supplier', 0, NULL),
('Ljetopis DOO','S5', '03450876', 'Iva Vizina', '', 0, '2023-10-24 13:09:57.868696', '2024-05-26 16:54:09.357015', 'supplier', 0, NULL),
('Montefarm DOO', 'S6', '03350986', 'Vojina Katnića', '', 0, '2023-10-24 13:11:27.538157', '2024-05-26 16:53:34.531831', 'supplier', 0, NULL),
( 'Ina Crna Gora', 'ICG', '028111112', 'Podgorica bb', '', 0, '2023-11-20 08:52:19.504703', '2024-05-26 16:51:24.538323', 'supplier', 0, NULL),
( 'Multicom DOO', 'MLTC', '02759535', 'Podgorica bb', '', 0, '0001-01-01 00:00:00', '2024-05-26 16:50:34.679737', 'supplier', 0, NULL),
( 'Cleaning DOO', 'S7', '02845556', 'Podgorica', '', 0, '0001-01-01 00:00:00', '2024-05-26 16:52:39.552174', 'supplier', 0, NULL),
( 'Mineralna voda Rada', 'S8', '02744456', 'Bijelo Polje', 'Mineralna voda', 0, '0001-01-01 00:00:00', '2023-12-08 12:01:12.856304', 'supplier', 0, NULL),
( 'Oykos development DOO', 'OD', '03292053', 'Vojina Katnića br. 47', '', 0, '2023-11-24 09:36:11.655058', '2024-05-26 16:53:05.75583', 'supplier', 0, NULL),
( 'Kontiki travel', 'S9', '1717731', 'Iva Vizina 20', 'Turistička agencija', 0, '2023-12-08 09:22:18.719684', '2023-12-08 09:22:18.719685', 'supplier', 0, NULL),
( 'Dnevne novine Dan', 'S10', '8372171', 'Iva Vizina 34', 'Dnevne novine', 0, '2023-12-11 12:23:17.140523', '2023-12-11 12:23:17.140523', 'supplier', 0, NULL),
( 'G tech DOO', 'S11', '02813334', 'Đoka Miraševića bb', '', 0, '2023-12-19 08:49:27.448075', '2024-06-10 09:24:21.356267', 'supplier', 0, NULL),
( 'Tenero group DOO', 'S12', '02247089', 'Mutezira Karađuzovića bb', '', 0, '2024-05-29 09:31:12.832807', '2024-05-29 09:31:12.832807', 'supplier', 0, NULL),

--donators
( 'Ambasada SAD Crna Gora', 'SAD', 'CJJ', 'Vukašina Markovića', 'Ambasada SAD Crna Gora', 0, '2023-12-21 10:18:18.161462', '2023-12-21 10:18:18.161462', 'donation', 0, NULL),

--lawyers
( 'Boris Vujović', 'BV', '0246810', 'Slobode bb', 'Advokat Boris Vujović', 0, '2024-04-22 13:29:07.175348', '2024-04-26 09:50:28.99373', 'lawyer', 0, NULL),

--subjects
( 'Balša Brković', 'BB', '098765410', 'Zlatica bb', 'Advokat Balša Brković bb', 0, '2024-04-22 13:30:23.291091', '2024-04-22 13:30:23.291091', 'subjects', 0, NULL),
( 'Aleksandar Bošković', 'AB', '02811110', 'Vasa Raičkovića 4b', '', 0, '2024-06-07 08:39:46.56458', '2024-04-22 13:30:37.151303', 'subjects', 0, NULL),

--employees
( 'David Komnenović', 'DK', '1109000211019', 'Nikšić', '', 0, '2024-04-26 12:16:18.993666', '2024-05-26 19:20:52.536446', 'employee', 0, NULL),

--executors
( 'Filip Lalovic', 'FL', '11111', 'Tolosi', '', 0, '2024-06-05 13:01:47.555238', '2024-06-05 13:01:47.555238', 'executor', 0, NULL),

--municipalities
('Nikšić', 'NK', '02912349', 'Njegoševa bb', '', 0, '2024-04-09 13:26:07.231769', '2024-05-26 19:23:06.595644', 'municipalities', 13, NULL),
('Berane', 'BE', '12234', '29. novembar', '', 0, '2024-04-10 08:17:18.605125', '2024-04-25 08:20:53.210662', 'municipalities', 13, NULL),
('Budva', 'BD', '', '', '', 0, '2024-04-12 11:57:40.862917', '2024-04-12 11:57:40.862917', 'municipalities', 10, NULL),
('Podgorica', 'PG', '02019710', 'Njegoševa br. 13', '', 0, '2024-04-12 11:57:57.413109', '2024-06-12 20:50:01.972991', 'municipalities', 15, NULL),
('Cetinje', 'CT', '', '', '', 0, '2024-04-12 11:58:05.191097', '2024-04-12 11:58:05.191098', 'municipalities', 15, NULL),


--banks
( 'Erste banka', 'EB', '02811114', 'Arsenija Boljevića 2A', '', 0, '2024-06-13 08:26:02.989796', '2024-06-13 23:13:18.206226', 'bank', 0, NULL),
( 'Crnogorska komercijalna banka', 'CKB', '02812220', 'Moskovska bb', '', 0, '2024-06-13 23:08:23.505189', '2024-06-13 23:08:23.505189', 'bank', 0, NULL),
( 'Hipotekarna banka', 'HB', '02811112', 'Josipa Broza Tita bb', '', 0, '2024-06-13 23:09:14.866028', '2024-06-13 23:09:14.866028', 'bank', 0, NULL),
( 'NLB banka', 'NLB', '02811113', 'Bulevar Stanka Dragojevića 46', '', 0, '2024-06-13 23:10:41.038802', '2024-06-13 23:10:41.038802', 'bank', 0, NULL),
( 'Addiko bank', 'AB', '02811115', 'Bul. Džordža Vašingtona 98', '', 0, '2024-06-13 23:16:04.11153', '2024-06-13 23:16:04.111531', 'bank', 0, NULL),
( 'Lovćen banka', 'LB', '02811116', 'Bulevar knjaza Danila Petrovića 13/32', '', 0, '2024-06-13 23:24:23.923912', '2024-06-13 23:24:23.923912', 'bank', 0, NULL);
