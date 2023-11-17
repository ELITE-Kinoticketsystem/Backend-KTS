
Use KinoTicketSystem;

-- TODO Fix order to delete correctly
-- for now, first run create_database.sql to drop all tables
-- and then run inser_database.sql script
TRUNCATE TABLE movie_actors;
TRUNCATE TABLE movie_producers;
TRUNCATE TABLE actors;
TRUNCATE TABLE producers;
TRUNCATE TABLE user_movies;
TRUNCATE TABLE showings;
TRUNCATE TABLE movies;
TRUNCATE TABLE fsk;
TRUNCATE TABLE genres;
TRUNCATE TABLE users;
TRUNCATE TABLE user_types;
TRUNCATE TABLE events;
TRUNCATE TABLE orders;
TRUNCATE TABLE payment_methods;
TRUNCATE TABLE tickets;
TRUNCATE TABLE price_category;
TRUNCATE TABLE seats;
TRUNCATE TABLE seat_category;
TRUNCATE TABLE cinema_halls;
TRUNCATE TABLE theatres;
TRUNCATE TABLE address;


-- Insert the seat categories
INSERT INTO seat_category (category_name) VALUES ('Standard'), ('Premium'), ('Couple'), ('Disabled');

-- Insert the user types
INSERT INTO user_types (type_name) VALUES ('customer'), ('admin'), ('casher');

-- Insert the fsk values
INSERT INTO fsk (age) VALUES (0), (6), (12), (16), (18);

-- Insert the payment methods
INSERT INTO payment_methods (method_name) VALUES ('MasterCard'), ('PayPal'), ('ApplePay'), ('Visa'), ('Cash');

-- Insert the price categories
INSERT INTO price_category (category_name, price) VALUES ('StudentDiscount', 7.00), ('ChildDiscount', 10.00), ('ElderlyDiscount', 5.00), ('regular_price', 20.00);

-- Insert demo data into address table
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('221B Baker St', 'N/A', 'NW1 6XE', 'London', 'United Kingdom');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('1600 Pennsylvania Ave NW', 'N/A', '20500', 'Washington, D.C.', 'United States');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Champs-Élysées', 'N/A', '75008', 'Paris', 'France');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Piazza del Colosseo', '1', '00184', 'Rome', 'Italy');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Tokyo Tower', '4-2-8', '105-0011', 'Minato City, Tokyo', 'Japan');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('The Dubai Mall', 'Financial Center Rd', '31166', 'Dubai', 'United Arab Emirates');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('The Bund', 'Zhongshan Dong Yi Lu', '200002', 'Shanghai', 'China');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Kremlin', 'Moscow', '103073', 'Moscow', 'Russia');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Sydney Opera House', 'Bennelong Point', 'NSW 2000', 'Sydney', 'Australia');
INSERT INTO address (street, streetnr, zipcode, city, country) VALUES ('Christ the Redeemer', 'Parque Nacional da Tijuca', '22241-330', 'Rio de Janeiro', 'Brazil');

-- Insert demo data into theatres table
INSERT INTO theatres (name, address_id) VALUES ('Theatre 1', 2);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 2', 3);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 3', 4);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 4', 5);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 5', 6);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 6', 7);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 7', 8);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 8', 9);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 9', 10);
INSERT INTO theatres (name, address_id) VALUES ('Theatre 10', 1);

-- Insert demo data into cinema_halls table
-- Theatre 1
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 7', 1);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 8', 1);
-- Theatre 2
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 2);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 2);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 2);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 2);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 2);
-- Theatre 3
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 3);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 3);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 3);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 3);
-- Theatre 4
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 4);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 4);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 4);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 4);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 4);
-- Theatre 5
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 5);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 5);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 5);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 5);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 5);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 5);
-- Theatre 6
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 6);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 6);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 6);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 6);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 6);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 6);
-- Theatre 7
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 7);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 7', 7);
-- Theatre 8
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 8);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 8);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 8);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 8);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 8);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 8);
-- Theatre 9
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 9);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 9);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 9);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 9);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 9);
-- Theatre 10
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 1', 10);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 2', 10);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 3', 10);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 4', 10);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 5', 10);
INSERT INTO cinema_halls (name, theatre_id) VALUES ('Hall 6', 10);

-- Insert demo data into seats table for first 5 cinema halls
-- Cinema Hall 1
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (1, 'D', '5', 1);

-- Cinema Hall 2
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (2, 'D', '5', 1);

-- Cinema Hall 3
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (3, 'D', '5', 1);

-- Cinema Hall 4
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (4, 'D', '5', 1);

-- Cinema Hall 5
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (5, 'D', '5', 1);

-- Cinema Hall 6
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (6, 'D', '5', 1);

-- Cinema Hall 7
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (7, 'D', '5', 1);

-- Cinema Hall 8
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (8, 'D', '5', 1);

-- Cinema Hall 9
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (9, 'D', '5', 1);

-- Cinema Hall 10
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'A', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'A', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'A', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'A', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'A', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'B', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'B', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'B', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'B', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'B', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'C', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'C', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'C', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'C', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'C', '5', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'D', '1', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'D', '2', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'D', '3', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'D', '4', 1);
INSERT INTO seats (cinema_hall_id, row_nr, column_nr, seat_category_id) VALUES (10, 'D', '5', 1);

-- Insert demo data into users table
INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user1', 'password1', 'John', 'Doe', 'johndoe@example.com', '123-456-7890', 1, 2);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user2', 'password2', 'Jane', 'Doe', 'janedoe@example.com', '123-456-7890', 2, 1);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user3', 'password3', 'Bob', 'Smith', 'bobsmith@example.com', '123-456-7890', 3, 3);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user4', 'password4', 'Alice', 'Smith', 'alicesmith@example.com', '123-456-7890', 4, 3);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user5', 'password5', 'Tom', 'Jones', 'tomjones@example.com', '123-456-7890', 5, 1);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user6', 'password6', 'Samantha', 'Jones', 'samanthajones@example.com', '123-456-7890', 6, 1);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user7', 'password7', 'David', 'Lee', 'davidlee@example.com', '123-456-7890', 7, 1);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user8', 'password8', 'Emily', 'Lee', 'emilylee@example.com', '123-456-7890', 8, 1);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user9', 'password9', 'Michael', 'Johnson', 'michaeljohnson@example.com', '123-456-7890', 9, 2);

INSERT INTO users (username, password, firstname, lastname, email, phone_number, address_id, user_type_id) 
VALUES ('user10', 'password10', 'Sarah', 'Johnson', 'sarahjohnson@example.com', '123-456-7890', 10, 1);


-- Insert demo data into genres table
INSERT INTO genres (name)
VALUES ('Action'), ('Drama'), ('Crime'),  ('Fantasy'), ('Western'), ('Romance');


-- Insert demo data into movies table
INSERT INTO movies (title, description, releasedDate, timeInMin, fsk_id, genre_id)
VALUES 
    ('The Matrix', 'A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.', '1999-03-31', 136, 2, 1),
    ('The Shawshank Redemption', 'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.', '1994-09-23', 142, 2, 2),
    ('The Godfather', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.', '1972-03-24', 175, 2, 3),
    ('The Dark Knight', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', '2008-07-18', 152, 2, 1),
    ('12 Angry Men', 'A jury holdout attempts to prevent a miscarriage of justice by forcing his colleagues to reconsider the evidence.', '1957-04-10', 96, 1, 3),
    ('Schindler''s List', 'In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.', '1994-12-15', 195, 2, 2),
    ('The Lord of the Rings: The Return of the King', 'Gandalf and Aragorn lead the World of Men against Sauron''s army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.', '2003-12-17', 201, 3, 4),
    ('Pulp Fiction', 'The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.', '1994-10-14', 154, 2, 3),
    ('The Good, the Bad and the Ugly', 'A bounty hunting scam joins two men in an uneasy alliance against a third in a race to find a fortune in gold buried in a remote cemetery.', '1966-12-23', 178, 4, 5),
    ('The Lord of the Rings: The Fellowship of the Ring', 'A meek Hobbit from the Shire and eight companions set out on a journey to destroy the powerful One Ring and save Middle-earth from the Dark Lord Sauron.', '2001-12-19', 178, 3, 4),
    ('Forrest Gump', 'The presidencies of Kennedy and Johnson, the events of Vietnam, Watergate and other historical events unfold through the perspective of an Alabama man with an IQ of 75, whose only desire is to be reunited with his childhood sweetheart.', '1994-07-06', 142, 2, 6);