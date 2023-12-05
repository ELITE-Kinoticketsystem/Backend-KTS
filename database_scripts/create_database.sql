DROP DATABASE IF EXISTS KinoTicketSystem;
CREATE DATABASE IF NOT EXISTS KinoTicketSystem;
USE KinoTicketSystem;


DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS event_seats;
DROP TABLE IF EXISTS event_seat_categories;
DROP TABLE IF EXISTS event_movies;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS payment_methods;
DROP TABLE IF EXISTS seats;
DROP TABLE IF EXISTS seat_categories;
DROP TABLE IF EXISTS cinema_halls;
DROP TABLE IF EXISTS theatres;
DROP TABLE IF EXISTS user_movies;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS actor_pictures;
DROP TABLE IF EXISTS producer_pictures;
DROP TABLE IF EXISTS movie_actors;
DROP TABLE IF EXISTS movie_producers;
DROP TABLE IF EXISTS producers;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS movie_genres;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS price_categories;
DROP TABLE IF EXISTS addresses;


CREATE TABLE addresses
  (
     id       BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     street   VARCHAR(100) NOT NULL,
     street_nr VARCHAR(10) NOT NULL,
     zipcode  VARCHAR(20) NOT NULL,
     city     VARCHAR(50) NOT NULL,
     country  VARCHAR(50) NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE genres
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     genre_name VARCHAR(40) NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE movies
  (
     id             BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     title          VARCHAR(80) NOT NULL,
     description    TEXT NOT NULL,
     banner_pic_url TEXT,
     cover_pic_url  TEXT,
     trailer_url    TEXT,
     rating         FLOAT,
     release_date   Date NOT NULL,
     time_in_min    INT NOT NULL,
     fsk            INT NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE movie_genres
  (
     movie_id BINARY(16) NOT NULL,
     genre_id BINARY(16) NOT NULL,
     PRIMARY KEY (movie_id, genre_id),
     FOREIGN KEY (movie_id) REFERENCES movies(id),
     FOREIGN KEY (genre_id) REFERENCES genres(id)
  );

CREATE TABLE producers
  (
     id          BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     name        VARCHAR(50) NOT NULL,
     age         DATE NOT NULL,
     description TEXT NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE actors
  (
     id          BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     name        VARCHAR(50) NOT NULL,
     age         DATE NOT NULL,
     description TEXT NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE actor_pictures
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     actor_id   BINARY(16) NOT NULL,
     pic_url    TEXT,
     PRIMARY KEY (id),
     FOREIGN KEY (actor_id) REFERENCES actors(id)
  );

CREATE TABLE producer_pictures
  (
     id            BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     producer_id   BINARY(16) NOT NULL,
     pic_url       TEXT,
     PRIMARY KEY (id),
     FOREIGN KEY (producer_id) REFERENCES producers(id)
  );

CREATE TABLE movie_producers
  (
     movie_id    BINARY(16) NOT NULL,
     producer_id BINARY(16) NOT NULL,
     PRIMARY KEY (movie_id, producer_id),
     FOREIGN KEY (movie_id) REFERENCES movies(id),
     FOREIGN KEY (producer_id) REFERENCES producers(id)
  );

CREATE TABLE movie_actors
  (
     movie_id BINARY(16) NOT NULL,
     actor_id BINARY(16) NOT NULL,
     PRIMARY KEY (movie_id, actor_id),
     FOREIGN KEY (movie_id) REFERENCES movies(id),
     FOREIGN KEY (actor_id) REFERENCES actors(id)
  );

CREATE TABLE users
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     username   VARCHAR(50),
     email      VARCHAR(80) NOT NULL,
     password   VARCHAR(50) NOT NULL,
     firstname  VARCHAR(50),
     lastname   VARCHAR(50),
     PRIMARY KEY (id)
  );

CREATE TABLE user_movies
  (
     user_id   BINARY(16) NOT NULL,
     movie_id  BINARY(16) NOT NULL,
     list_type VARCHAR(40) NOT NULL,
     PRIMARY KEY (user_id, movie_id, list_type),
     FOREIGN KEY (user_id) REFERENCES users (id),
     FOREIGN KEY (movie_id) REFERENCES movies (id)
  );

CREATE TABLE reviews
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     rating     INT NOT NULL,
     comment    TEXT NOT NULL,
     datetime   TIMESTAMP NOT NULL,
     is_spoiler BOOLEAN,
     user_id    BINARY(16) NOT NULL,
     movie_id   BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) REFERENCES users(id),
     FOREIGN KEY (movie_id) REFERENCES movies(id)
  );

CREATE TABLE theatres
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     name       VARCHAR(50) NOT NULL,
     logo_url   TEXT,
     address_id BINARY(16) NOT NULL,
     PRIMARY KEY(id),
     FOREIGN KEY (address_id) REFERENCES addresses(id)
  );

CREATE TABLE cinema_halls
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     name       VARCHAR(50) NOT NULL,
     capacity   INT NOT NULL,
     theatre_id BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (theatre_id) REFERENCES theatres(id)
  );

CREATE TABLE seat_categories
  (
     id            BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     category_name VARCHAR(30) NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE seats
  (
     id               BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     row_nr           INT NOT NULL,
     column_nr        INT NOT NULL,
     seat_category_id BINARY(16) NOT NULL,
     cinema_hall_id   BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (seat_category_id) REFERENCES seat_categories(id),
     FOREIGN KEY (cinema_hall_id) REFERENCES cinema_halls(id)
  );

CREATE TABLE payment_methods
  (
     id         BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     methodname VARCHAR(30) NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE orders
  (
     id                BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     totalprice        INT NOT NULL,
     is_paid           BOOLEAN NOT NULL,
     payment_method_id BINARY(16) NOT NULL,
     user_id           BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id),
     FOREIGN KEY (user_id) REFERENCES users(id)
  );


CREATE TABLE events
  (
     id             BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     title          VARCHAR(50) NOT NULL,
     start          DateTime NOT NULL,
     end            DateTime NOT NULL,
     description    TEXT, 
     event_type     VARCHAR(40) NOT NULL,
     cinema_hall_id BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (cinema_hall_id) REFERENCES cinema_halls(id)
  );

CREATE TABLE event_movies
  (
     event_id BINARY(16) NOT NULL,
     movie_id BINARY(16) NOT NULL,
     PRIMARY KEY (event_id, movie_id),
     FOREIGN KEY (event_id) REFERENCES events(id),
     FOREIGN KEY (movie_id) REFERENCES movies(id)
  );

CREATE TABLE price_categories
  (
     id            BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     category_name VARCHAR(40) NOT NULL,
     price         INT NOT NULL,
     PRIMARY KEY (id)
  );

CREATE TABLE event_seat_categories
  (
     event_id         BINARY(16) NOT NULL,
     seat_category_id BINARY(16) NOT NULL,
     price            INT NOT NULL,
     PRIMARY KEY (event_id, seat_category_id),
     FOREIGN KEY (event_id) REFERENCES events(id),
     FOREIGN KEY (seat_category_id) REFERENCES seat_categories(id)
  );

CREATE TABLE event_seats
  (
     id            BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     booked        BOOLEAN NOT NULL,
     blocked_until TIMESTAMP,
     user_id       BINARY(16) NOT NULL,
     seat_id       BINARY(16) NOT NULL,
     event_id      BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (user_id) REFERENCES users(id),
     FOREIGN KEY (seat_id) REFERENCES seats(id),
     FOREIGN KEY (event_id) REFERENCES events(id)
  );

CREATE TABLE tickets
  (
     id                BINARY(16) DEFAULT (Uuid_to_bin(Uuid(), 1)),
     validated         BOOLEAN NOT NULL,
     price             INT NOT NULL,
     price_category_id BINARY(16) NOT NULL,
     order_id          BINARY(16) NOT NULL,
     event_seat_id     BINARY(16) NOT NULL,
     PRIMARY KEY (id),
     FOREIGN KEY (price_category_id) REFERENCES price_categories(id),
     FOREIGN KEY (order_id) REFERENCES orders(id),
     FOREIGN KEY (event_seat_id) REFERENCES event_seats(id)
  ); 
