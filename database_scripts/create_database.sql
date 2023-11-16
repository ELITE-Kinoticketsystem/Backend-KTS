-- To access mysql through command-line: /Applications/xampp/xamppfiles/bin/mysql -u root

Create database if not EXISTS KinoTicketSystem;
Use KinoTicketSystem;

-- Order of dropping matters because of foreign keys
DROP TABLE IF EXISTS movie_actors;
DROP TABLE IF EXISTS movie_producers;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS producers;
DROP TABLE IF EXISTS user_movies;
DROP TABLE IF EXISTS showings;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS fsk;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_types;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS payment_methods;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS price_category;
DROP TABLE IF EXISTS seats;
DROP TABLE IF EXISTS seat_category;
DROP TABLE IF EXISTS cinema_halls;
DROP TABLE IF EXISTS theatres;
DROP TABLE IF EXISTS address;




-- Create the address table
CREATE TABLE address (
    id INT PRIMARY KEY AUTO_INCREMENT,
    street VARCHAR(255) NOT NULL,
    streetnr VARCHAR(10) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    city VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL
);


-- Create the theatre table with a foreign key to the address table
CREATE TABLE theatres (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    address_id INT NOT NULL,
    FOREIGN KEY (address_id) REFERENCES address(id)
);

-- Create the cinema_halls table with a foreign key to the theatre table
CREATE TABLE cinema_halls (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    capacity INT NOT NULL,
    theatre_id INT NOT NULL,
    FOREIGN KEY (theatre_id) REFERENCES theatres(id)
);

-- Create the seat_category table
CREATE TABLE seat_category (
    id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255) NOT NULL
);


-- Create the seats table with foreign keys to the cinema_halls table and the seat_category table
CREATE TABLE seats (
    id INT PRIMARY KEY AUTO_INCREMENT,
    row_nr INT NOT NULL,
    column_nr INT NOT NULL,
    seat_category_id INT NOT NULL,
    cinema_hall_id INT NOT NULL,
    FOREIGN KEY (seat_category_id) REFERENCES seat_category(id),
    FOREIGN KEY (cinema_hall_id) REFERENCES cinema_halls(id)
);

-- Create the price_category table
CREATE TABLE price_category (
    category_name VARCHAR(255) PRIMARY KEY,
    price DECIMAL(10,2) NOT NULL
);

-- Create the tickets table with foreign keys to the seats table and the price_category table
CREATE TABLE tickets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    validated BOOLEAN NOT NULL DEFAULT FALSE,
    paid BOOLEAN NOT NULL DEFAULT FALSE,
    reserved BOOLEAN NOT NULL DEFAULT FALSE,
    price DECIMAL(10,2) NOT NULL,
    seat_id INT NOT NULL,
    price_category_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (seat_id) REFERENCES seats(id),
    FOREIGN KEY (price_category_name) REFERENCES price_category(category_name)
);

-- Create the payment_methods table
CREATE TABLE payment_methods (
    id INT PRIMARY KEY AUTO_INCREMENT,
    method_name VARCHAR(255) NOT NULL
);

-- Create the orders table with a foreign key to the tickets table and a foreign key to the payment_methods table
CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    total_price DECIMAL(10,2) NOT NULL,
    ticket_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    reservation BOOLEAN NOT NULL DEFAULT FALSE,
    booking BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (ticket_id) REFERENCES tickets(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- Create the fsk table
CREATE TABLE fsk (
    id INT PRIMARY KEY AUTO_INCREMENT,
    age INT NOT NULL
);

-- Create the movies table with a foreign key to the fsk table
CREATE TABLE movies (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    releasedDate DATE NOT NULL,
    timeInMin INT NOT NULL,
    fsk_id INT NOT NULL,
    FOREIGN KEY (fsk_id) REFERENCES fsk(id)
);

-- Create the user_types table
CREATE TABLE user_types (
    id INT PRIMARY KEY AUTO_INCREMENT,
    type_name VARCHAR(255) NOT NULL
);

-- Create the users table with a foreign key to the address table and a foreign key to the user_types table
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    password VARCHAR(255) NOT NULL,
    address_id INT NOT NULL,
    user_type_id INT NOT NULL,
    FOREIGN KEY (address_id) REFERENCES address(id),
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

-- Create the user_movies table with foreign keys to the users table and the movies table
CREATE TABLE user_movies (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    movie_id INT NOT NULL,
    list_type VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);
-- Create the events table
CREATE TABLE events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    cinema_hall_id INT NOT NULL,
    FOREIGN KEY (cinema_hall_id) REFERENCES cinema_halls(id)
);

-- Create the showings table with foreign keys to the events table and the movies table
CREATE TABLE showings (
    movie_id INT NOT NULL,
    event_id INT NOT NULL,
    PRIMARY KEY (movie_id, event_id),
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);
-- Create the actors table
CREATE TABLE actors (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
);

-- Create the producers table
CREATE TABLE producers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
);

-- Create the movie_actors table with foreign keys to the movies table and the actors table
CREATE TABLE movie_actors (
    movie_id INT NOT NULL,
    actor_id INT NOT NULL,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (actor_id) REFERENCES actors(id)
);

-- Create the movie_producers table with foreign keys to the movies table and the producers table
CREATE TABLE movie_producers (
    movie_id INT NOT NULL,
    producer_id INT NOT NULL,
    PRIMARY KEY (movie_id, producer_id),
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (producer_id) REFERENCES producers(id)
);


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
