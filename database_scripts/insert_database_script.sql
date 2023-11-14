-- Address 
INSERT INTO address (street, streetnr , zipcode, city , country )
VALUES
("Ap #543-3786 Quis, Rd.","2","74554", "Waren", "India"),
("P.O. Box 672, 2411 Tristique Road","4","22714", "Dortmund", "Sweden"),
("2113 Nam Street", "2", "51936", "Heide", "Mexico"),
("2581 Consequat St.","8", "94372", "Grimma", "Turkey"),
("Ap #215-2603 Et Street", "12", "28424", "Riesa", "France"),
("P.O. Box 809, 7446 Non Road", "14","78391", "Panketal", "United Kingdom"),
("Ap #180-2543 Pede St.","12", "33441", "Rendsburg", "Turkey"),
("Ap #379-7061 Vitae Rd.","20","80975", "Freital", "Chile"),
("341-4618 Sagittis. Rd.", "18","42864", "Bremerhaven", "Ireland"),
("4671 Amet Avenue", "16", "74576", "Rödermark", "Austria"),
("Ap #449-6689 Urna Ave", "10", "07194", "Wandlitz", "South Africa"),
("Ap #440-3900 Orci, Street", "20", "40100", "Villingen-Schwenningen", "Chile"),
("Ap #339-7551 Rutrum. Rd.","8", "07941", "Wadgassen", "Chile"),
("243-9956 Non St.","14","48351", "Hamburg", "Netherlands"),
("6534 Et St.","6", "13833", "Grimma", "Norway"),
("Ap #709-8162 Nec, St.", "12","02790", "Hamburg", "Chile"),
("469-9731 Imperdiet Av.","10", "66760", "Andernach", "Spain"),
("9282 Ac Rd.", "16","73451", "Pirna", "Russian Federation"),
("3670 Mauris Street","6", "10861", "Norderstedt", "China"),
("Ap #882-6314 Lacinia St.","8","48486", "Neunkirchen", "Nigeria");


-- theatres
INSERT INTO theatres (name, address_id)
VALUES
("Nibh Donec Est LLP", 19), ("Placerat Orci Inc.",8), ("Nostra Per Ltd", 21),
("Eu Erat Semper Incorporated", 19), ("Pharetra Felis Eget LLP",11),
("Aliquam Enim Limited", 3), ("Euismod Enim Etiam Limited", 21), ("Nullam Scelerisque Corporation", 4), ("Eu Euismod PC", 19), ("Donec Est Ltd", 15), ("Lobortis Risus In PC",2), ("Risus Limited",21),
("Vitae Erat Associates", 17), ("Sed Dictum Associates", 22), ("Lectus Pede Industries",6), ("Arcu Aliquam Company", 20), ("Pretium Aliquet Corporation", 22),
("Eros Limited", 10), ("Ultrices LLP",1),
("Egestas Urna Industries", 23);


-- cinema_halls
INSERT INTO cinema_halls (name, capacity , theatres_id)
VALUES
("Nibh Donec Est LLP", 75,2),
("Placerat Orci Inc.", 67,9),
("'Nostra Per Ltd" ,56, 9) ,
("Eu Erat Semper Incorporated", 55,6),
("Pharetra Felis Eget LLP", 63,9),
("Aliquam Enim Limited", 66,1),
("Euismod Enim Etiam Limited",55,7),
("Nullam Scelerisque Corporation", 57,6),
("Eu Euismod PC", 72,4),
("Donec Est Ltd", 74,9),
("Lobortis Risus In PC", 65,3),
("Risus Limited", 51,12),
("Vitae Erat Associates", 53,20),
("Sed Dictum Associates", 66,5),
("Lectus Pede Industries", 53,10),
("Arcu Aliquam Company", 51,1),
("Pretium Aliquet Corporation", 51,12),
("Eros Limited", 56,13),
("Ultrices LLP", 62,14),
("Egestas Urna Industries", 61,6);

-- seats
INSERT INTO seats (row_nr, seat_nr, category, cinema_halls_id)
VALUES
(9,2, "Premium", 2),
(2,2, "couple", 9),
(1,5, "Standard", 9),
(8,7, "disabled", 6),
(3,5, "couple", 9),
(0,0, "disabled",1),
(3,2, "Standard", 7),
(3,7, "disabled", 6),
(1,2, "couple", 4),
(9,10, "Premium", 9),
(5,8, "couple", 3),
(1,5, "Standard", 12),
(4,1, "Premium", 20),
(9,9, "disabled", 5),
(6,6, "disabled", 10),
(2,7, "couple", 1),
(3,3,"Premium", 12),
(4,2, "Standard", 13),
(1,7, "Premium", 14),
(3,6, "Premium", 6);

-- Tickets
INSERT INTO tickets (price, date_stamp, validated, paid, reserved, seat_row_nr, seat_seat_nr)
VALUES
(15.71, "2023-07-19 22:26:30", 1,1,1,8,2),
(5.92, "2024-01-27 17:40:37", 1,1,1,5,10),
(7.64, "2024-10-25 22:56:19", 2,2,2,3,8),
(6.43, "2024-01-30 21:03:20", 2,2,2,3,3),
(19.6, "2023-05-18 02:21:37", 1,1,1,5,10),
(12.27, "2023-05-02 03:07:34", 1,1,1,10,6),
(15.24, "2024-03-20 17:39:38" ,2,2,2,5,9),
(15.29, "2023-04-18 14:14:54", 2,2,2,2,4),
(19-16, "2022-12-25 08:34:47" ,1,1,1,9,6),
(15.33, "2024-06-10 19:45:16", 1, 1, 1,5,9),
(8.6, "2023-01-24 21:56:34", 2,2,2,6,7),
(12.8, "2023-07-17 01:50:52" , 2,2,2,3,6),
(12.37, "2022-11-30 07:46:26", 1,1, 1,9,3),
(15.1, "2023-05-25 05:55:21", 1,1,1,6,2),
(5.46, "2024-03-02 01:51:47", 2,2,2,7,3),
(14.33, "2024-08-24 22:01:16", 2,2,2,2,8),
(11.23, "2024-01-27 19:59:46", 1,1,1,4,5),
(19.01, "2023-06-30 18:27:47",1,1,1,9,2),
(8.98, "2024-06-22 12:26:49", 2,2,2,5,10),
(6.30, "2023-01-03 13:53:33" ,2,2,2,3,2);



-- movies
INSERT INTO movies (name, description, releaseDate, timeInMinutes, fsk)
VALUES
("GJM44LSK4IW", "et ultrices", "2021-09-04",137,"16"),
("OEKZONIR7JP", "volutpat ornare, facilisis eget, ipsum. Donec sollicitudin adipiscing ligula. Aenean gravida nunc sed pede.", "2021-10-08",110, "16"),
("SLR12WGV6SX", "Nunc pulvinar arcu et pede. Nunc", "2020-11-18", 112, "18"),
("BFI27VMS3KD", "auctor odio a purus. Duis elementum, dui quis accumsan convallis, ante lectus","2021-10-24",102,"16"),
("DTN41XOR5UM", "Sed diam lorem, auctor quis, tristique ac, eleifend vitae, erat. Vivamus nisi. Mauris nulla. Integer urna.", "2021-05-22",142,"16"),
("HGA23HTJZYV", "egestas ligula. Nullam feugiat placerat velit. Quisque varius. Nam porttitor scelerisque neque. Nullam nisl.", "2021-08-21", 110, "18"),
("ZFD420DA8SY", "lorem, vehicula et, rutrum eu, ultrices sit amet, risus. Donec nibh enim, gravida sit amet, dapibus id, blandit at, nisi. Cum sociis natoque penatibus et magnis","2021-08-24" , 113, "16"),
("'QEX91RRS1UB", "mollis. Phasellus libero mauris, aliquam eu, accumsan sed, facilisis vitae, orci. Phasellus dapibus quam quis diam. Pellentesque habitant morbi tristique", "2021-02-04", 134, "6"),
("JFQ51FJS4BP", "et ultrices posuere cubilia Curae Donec tincidunt. Donec vitae erat vel", "2021-08-23", 143, "18"),
("AFP55DIF5RZ", "lorem ut aliquam iaculis, lacus pede sagittis augue, eu tempor erat neque non quam. Pellentesque habitant", "2021-10-28", 138, "16"),
("HHH91EEI1JW", "mollis dui, in sodales elit erat vitae risus. Duis a mi fringilla mi lacinia mattis. Integer eu lacus. Quisque imperdiet, erat nonummy ultricies ornare, elit elitfermentum risus.", "2021-01-08", 133, "16"),
("ROI85LOS3ZU", "non justo. Proin non massa non ante bibendum ullamcorper. Duis cursus, diam at pretium aliquet, metus urna convallis erat, ", "2021-04-16", 129, "12"),
("NTU24HFN7EM", "sollicitudin commodo ipsum. Suspendisse non leo. Vivamus nibh dolor, nonummy ac, feugiat non, lobortis quis, pede. Suspendisse dui. Fusce diam nunc, ullamcorper eu, euismod ac.", "2020-12-05",127,"16"),
("FBE77GCU4NL", "dapibus quam quis diam. Pellentesque habitant morbi tristique senectus", "2021-01-08", 122,"6"),
("TSJ31RBA1XE", "elit. Etiam laoreet, libero et tristique pellentesque, tellus sem mollis dui, in sodales elit erat vitae risus. Duis a mi fringilla mi lacinia mattis. Integer eu lacus.", "2021-10-05", 131,"6"),
("JIN98SAP5FL","leo. Vivamus nibh dolor, nonummy ac, feugiat non, lobortis quis, pede.","2021-01-30", 103, "12"),
("MNK11EBO2BD", "In scelerisque scelerisque dui. Suspendisse ac", "2020-10-12", 110, "12"),
("BIO21RXM3MY", "turpis non enim. Mauris quis turpis","2020-12-22",104, "16"),
("TQI91IIW1RB", "elit sed consequat auctor, nunc nulla vulputate dui, nec tempus mauris erat", "2020-12-26", 111, "12"),
("WHB66HIN3AN", "tempor augue ac ipsum. Phasellus vitae mauris sit amet lorem semper auctor. Mauris vel turpis. Aliquam adipiscing lobortis risus. In mi pede, nonummy ut, ", "2021-02-10", 137,"6") ;


-- user
INSERT INTO user (username, firstname,email, password, address_id)
VALUES
("Olson", "Vance", "h-vance@ic loud net", "BPM28IEN5RX", 6) , 
("Fitzpatrick", "Germane", "d_germane7289@yahoo.net", "QMV67ZJF9BQ", 11),
("Moses", "Noah", "fry,noah@google.org", "ERY48MBY6DP" ,6),
("Pickett", "Iliana", "k_iliana@aol.com", "IEE42ST04TV", 9) ,
("James", "Vielka", "vielka-pollard@protonmail,couk", "RPX03UTX1ER" ,7) ,
("Mcpherson", "Steel", "s_burgess7052@icloud.ca", "WVT960AN6ZR", 11), 
("Walton", "Lucas", "lucas.merril1705@icloud.couk", "KYQ24FV09B0", 6),
("Cervantes", "Denton", "osborne_denton1428@aol.com", "HWA44UZC3MB", 14), 
("Kline", "Indigo", "1-indigo@aol.net", "ZWE66QEW4ST", 8),
("Haney", "Ursula", "ursulawitt@protonmail.net", "XUM26LIS6RU" ,20),
("Bowen", "Brynn", "bblackburn@google. couk", "NLJ40TBT3QJ",2),
("Mcdowell", "Lawrence", "1_harrington2577@yahoo. ca", "LCP64NAH8UG", 11), 
("Sanford", "Teegan", "rteegan@google. net", "OSP60TEN6GF", 19),
("Barton", "Alma", "a.mcgee@protonmail.couk", "JNQ41BDV7OV", 13),
("Buchanan", "William", "h_william7006@google.com", "TEY53CES2IU", 3),
("Conley", "Murphy", "carter_murphy@icloud.edu", "OKS25FWM7KR" ,1),
("Sweeney", "Maia", "m.palmer@icloud.org", "REY66KWQ8IS", 18),
("Kane", "Dante", "d_sandoval@aol. edu", "FWH57HSX6YZ" ,8),
("Garza", "Edan", "edan-cox271@out look. org", "NOC73LRV3JL", 14),
("Cardenas", "Devin", "molina_devin@aol.com", "KEP36ASY2XB", 1) ;

-- user_watchlist
-- not yet

-- user_cinemas
-- not yet

-- events
INSERT INTO events (start_date, end_date, cinema_halls_id, cinema_halls_theatres_id)
VALUES
("2024-02-07 10:25:21", "2023-04-09 16:54:17",2,14),
("2024-08-20 19:26:24", "2024-03-08 08:51:00", 15,4),
("2024-10-22 10:57:29", "2024-08-16 22:23:46", 12, 18),
("2024-02-15 02:52:43", "2024-08-19 20:47:41", 12, 9),
("2024-05-11 05:05:24", "2024-08-03 04:36:06", 2, 0),
("2023-03-25 18:32:42", "2024-06-25 05:12:15",3,9),
("2022-12-18 08:39:21", "2023-08-22 22:21:19",5, 16),
("2023-07-03 13:54:28", "2023-08-29 00:15:05", 4, 4),
("2024-04-12 02:11:26", "2023-12-23 05:11:44",17,4),
("2022-11-22 00:40:52","2023-05-02 15:43:45", 3,20),
("2024-08-15 13:05:49", "2024-11-10 17:49:59", 7,8),
("2024-03-06 15:39:07", "2023-08-08 07:58:27", 17,6),
("2023-03-29 11:16:21", "2023-11-26 09:26:19", 4,12),
("2023-02-03 07:25:05", "2024-07-01 17:24:57", 8, 20),
("2024-08-05 12:05:24", "2024-10-13 22:28:46", 10, 2),
("2023-04-15 21:57:22", "2023-01-29 08:19:38", 17,11),
("2023-02-23 17:24:00", "2023-03-19 00:56:14", 4,13),
("2023-08-31 20:11:50", "2024-10-22 16:09:41", 9, 14),
("2023-11-27 10:17:20", "2024-05-15 09:06:24", 8, 16),
("2023-11-19 00:06:20", "2023-07-30 10:58:38", 9,14);