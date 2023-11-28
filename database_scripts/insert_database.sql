USE KinoTicketSystem_v2;


Delete from tickets;
Delete from event_seats;
Delete from event_seat_categories;
Delete from event_movies;
Delete from events;
Delete from orders;
Delete from payment_methods;
Delete from seats;
Delete from seat_categories;
Delete from cinema_halls;
Delete from theatres;
Delete from user_movies;
Delete from reviews;
Delete from users;
Delete from actor_pictures;
Delete from producer_pictures;
Delete from movie_actors;
Delete from movie_producers;
Delete from producers;
Delete from actors;
Delete from movie_genres;
Delete from movies;
Delete from genres;
Delete from price_categories;
Delete from addresses;


INSERT INTO addresses (id, street, street_nr, zipcode, city, country) 
VALUES
    (UUID_TO_BIN('6ba7b810-9dad-11d1-80b4-00c04fd430c8'), '123 Main St', '1', '12345', 'New York', 'USA'), 
    (UUID_TO_BIN('6ba7b811-9dad-11d1-80b4-00c04fd430c8'), '456 Elm St', '2', '23456', 'Los Angeles', 'USA'), 
    (UUID_TO_BIN('6ba7b812-9dad-11d1-80b4-00c04fd430c8'), '789 Oak St', '3', '34567', 'Chicago', 'USA'), 
    (UUID_TO_BIN('6ba7b813-9dad-11d1-80b4-00c04fd430c8'), '321 Pine St', '4', '45678', 'San Francisco', 'USA'), 
    (UUID_TO_BIN('6ba7b814-9dad-11d1-80b4-00c04fd430c8'), '654 Maple St', '5', '56789', 'Seattle', 'USA'), 
    (UUID_TO_BIN('6ba7b815-9dad-11d1-80b4-00c04fd430c8'), '987 Cedar St', '6', '67890', 'Miami', 'USA'), 
    (UUID_TO_BIN('6ba7b816-9dad-11d1-80b4-00c04fd430c8'), '654 Birch St', '7', '78901', 'Dallas', 'USA'), 
    (UUID_TO_BIN('6ba7b817-9dad-11d1-80b4-00c04fd430c8'), '321 Walnut St', '8', '89012', 'Boston', 'USA'), 
    (UUID_TO_BIN('6ba7b818-9dad-11d1-80b4-00c04fd430c8'), '987 Cherry St', '9', '90123', 'Denver', 'USA'), 
    (UUID_TO_BIN('6ba7b819-9dad-11d1-80b4-00c04fd430c8'), '654 Grape St', '10', '01234', 'Phoenix', 'USA'); 


-- Insert statement for PriceCategory table
INSERT INTO price_categories (id, category_name, price)
VALUES
    (UUID_TO_BIN('6ba7b81c-9dad-11d1-80b4-00c04fd430c0'), 'StudentDiscount', 7.00), 
    (UUID_TO_BIN('6ba7b81d-9dad-11d1-80b4-00c04fd430c1'), 'ChildDiscount', 10.00),
    (UUID_TO_BIN('6ba7b81e-9dad-11d1-80b4-00c04fd430c2'), 'ElderlyDiscount', 5.00), 
    (UUID_TO_BIN('6ba7b81f-9dad-11d1-80b4-00c04fd430c3'), 'regular_price', 20.00); 

-- Insert statement for Genre table
INSERT INTO genres (id, genre_name)
VALUES
    (UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4'), 'Action'),
    (UUID_TO_BIN('6ba7b821-9dad-11d1-80b4-00c04fd430c5'), 'Drama'),
    (UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6'), 'Crime'),
    (UUID_TO_BIN('6ba7b823-9dad-11d1-80b4-00c04fd430c7'), 'Fantasy'),
    (UUID_TO_BIN('6ba7b824-9dad-11d1-80b4-00c04fd430c8'), 'Western'),
    (UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9'), 'Romance');

-- Insert statement for Movie table
INSERT INTO movies (id, title, description, release_date, time_in_min, fsk)
VALUES
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'The Shawshank Redemption', 'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.', '1994-10-14', 142, 16),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'The Godfather', 'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.', '1972-03-24', 175, 16),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'The Dark Knight', 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', '2008-07-18', 152, 12),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'Pulp Fiction', 'The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.', '1994-10-14', 154, 18),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'Fight Club', 'An insomniac office worker and a devil-may-care soapmaker form an underground fight club that evolves into something much, much more.', '1999-10-15', 139, 18),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'Forrest Gump', 'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold through the perspective of an Alabama man with an IQ of 75, whose only desire is to be reunited with his childhood sweetheart.', '1994-07-06', 142, 12),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'The Matrix', 'A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.', '1999-03-31', 136, 16),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'Goodfellas', 'The story of Henry Hill and his life in the mob, covering his relationship with his wife Karen Hill and his mob partners Jimmy Conway and Tommy DeVito in the Italian-American crime syndicate.', '1990-09-19', 146, 16),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'The Lord of the Rings: The Fellowship of the Ring', 'A meek Hobbit from the Shire and eight companions set out on a journey to destroy the powerful One Ring and save Middle-earth from the Dark Lord Sauron.', '2001-12-19', 178, 12),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'Inception', 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O.', '2010-07-16', 148, 12);

INSERT Into movie_genres (movie_id, genre_id)  
VALUES
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b821-9dad-11d1-80b4-00c04fd430c5')),

    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9')),

    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6')),
    
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b824-9dad-11d1-80b4-00c04fd430c8')),
    
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9')),

    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b821-9dad-11d1-80b4-00c04fd430c5')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b824-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6') ),

    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),

    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b821-9dad-11d1-80b4-00c04fd430c5')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9')),

    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b821-9dad-11d1-80b4-00c04fd430c5')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b822-9dad-11d1-80b4-00c04fd430c6')),

    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'),UUID_TO_BIN('6ba7b820-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b825-9dad-11d1-80b4-00c04fd430c9'))
;
-- Insert statement for Producer table
INSERT INTO producers (id, name, age, description)
VALUES
    (UUID_TO_BIN('6ba7b830-9dad-11d1-80b4-00c04fd430c0'), 'John Doe', 25, "John Doe is a producer."),
    (UUID_TO_BIN('6ba7b831-9dad-11d1-80b4-00c04fd430c1'), 'Jane Smith', 30, "Jane Smith is a producer."),
    (UUID_TO_BIN('6ba7b832-9dad-11d1-80b4-00c04fd430c2'), 'Michael Johnson', 35, "Michael Johnson is a producer."),
    (UUID_TO_BIN('6ba7b833-9dad-11d1-80b4-00c04fd430c3'), 'Emily Davis', 28, "Emily Davis is a producer."),
    (UUID_TO_BIN('6ba7b834-9dad-11d1-80b4-00c04fd430c4'), 'David Wilson', 32, "David Wilson is a producer."),
    (UUID_TO_BIN('6ba7b835-9dad-11d1-80b4-00c04fd430c5'), 'Sarah Anderson', 29, "Sarah Anderson is a producer."),
    (UUID_TO_BIN('6ba7b836-9dad-11d1-80b4-00c04fd430c6'), 'Christopher Taylor', 31, "Christopher Taylor is a producer."),
    (UUID_TO_BIN('6ba7b837-9dad-11d1-80b4-00c04fd430c7'), 'Jessica Martinez', 27, "Jessica Martinez is a producer."),
    (UUID_TO_BIN('6ba7b838-9dad-11d1-80b4-00c04fd430c8'), 'Matthew Brown', 33, "Matthew Brown is a producer."),
    (UUID_TO_BIN('6ba7b839-9dad-11d1-80b4-00c04fd430c9'), 'Olivia Garcia', 26, "Olivia Garcia is a producer.");


-- Insert statement for Actors table
INSERT INTO actors (id, name, age, description)
VALUES
    (UUID_TO_BIN('6ba7b840-9dad-11d1-80b4-00c04fd430c0'), 'Tom Hanks', 65, "Tom Hanks is an actor."),
    (UUID_TO_BIN('6ba7b841-9dad-11d1-80b4-00c04fd430c1'), 'Leonardo DiCaprio', 46, "Leonardo DiCaprio is an actor."),
    (UUID_TO_BIN('6ba7b842-9dad-11d1-80b4-00c04fd430c2'), 'Brad Pitt', 57, "Brad Pitt is an actor."),
    (UUID_TO_BIN('6ba7b843-9dad-11d1-80b4-00c04fd430c3'), 'Meryl Streep', 72, "Meryl Streep is an actor."),
    (UUID_TO_BIN('6ba7b844-9dad-11d1-80b4-00c04fd430c4'), 'Robert Downey Jr.', 56, "Robert Downey Jr. is an actor."),
    (UUID_TO_BIN('6ba7b845-9dad-11d1-80b4-00c04fd430c5'), 'Scarlett Johansson', 36, "Scarlett Johansson is an actor."),
    (UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6'), 'Denzel Washington', 66, "Denzel Washington is an actor."),
    (UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7'), 'Jennifer Lawrence', 31, "Jennifer Lawrence is an actor."),
    (UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8'), 'Johnny Depp', 58, "Johnny Depp is an actor."),
    (UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9'), 'Emma Stone', 33, "Emma Stone is an actor.");


-- Insert statement for MovieProducers table
INSERT INTO movie_producers (movie_id, producer_id)
VALUES
    -- Movie 1
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b830-9dad-11d1-80b4-00c04fd430c0')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b831-9dad-11d1-80b4-00c04fd430c1')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b832-9dad-11d1-80b4-00c04fd430c2')),
    -- Movie 2
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b833-9dad-11d1-80b4-00c04fd430c3')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b834-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b835-9dad-11d1-80b4-00c04fd430c5')),
    -- Movie 3
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b836-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b837-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b838-9dad-11d1-80b4-00c04fd430c8')),
    -- Movie 4
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b839-9dad-11d1-80b4-00c04fd430c9')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b830-9dad-11d1-80b4-00c04fd430c0')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b831-9dad-11d1-80b4-00c04fd430c1')),
    -- Movie 5
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b832-9dad-11d1-80b4-00c04fd430c2')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b833-9dad-11d1-80b4-00c04fd430c3')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b834-9dad-11d1-80b4-00c04fd430c4')),
    -- Movie 6
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b835-9dad-11d1-80b4-00c04fd430c5')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b836-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b837-9dad-11d1-80b4-00c04fd430c7')),
    -- Movie 7
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b838-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b839-9dad-11d1-80b4-00c04fd430c9')),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b830-9dad-11d1-80b4-00c04fd430c0')),
    -- Movie 8
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b831-9dad-11d1-80b4-00c04fd430c1')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b832-9dad-11d1-80b4-00c04fd430c2')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b833-9dad-11d1-80b4-00c04fd430c3')),
    -- Movie 9
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b834-9dad-11d1-80b4-00c04fd430c4')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b835-9dad-11d1-80b4-00c04fd430c5')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b836-9dad-11d1-80b4-00c04fd430c6')),
    -- Movie 10
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b837-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b838-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b839-9dad-11d1-80b4-00c04fd430c9'));



-- Insert statement for MovieActors table
INSERT INTO movie_actors (movie_id, actor_id)
VALUES
    -- Movie 1
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 2
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b842-9dad-11d1-80b4-00c04fd430c2')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 3
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 4
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 5
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b842-9dad-11d1-80b4-00c04fd430c2')),
    (UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 6
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 7
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 8
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b842-9dad-11d1-80b4-00c04fd430c2')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 9
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9')),
    -- Movie 10
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b846-9dad-11d1-80b4-00c04fd430c6')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b847-9dad-11d1-80b4-00c04fd430c7')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b848-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), UUID_TO_BIN('6ba7b849-9dad-11d1-80b4-00c04fd430c9'));



-- Insert statement for Users table
INSERT INTO users (id, username, email, password, firstname, lastname)
VALUES
    -- User 1
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), 'DoeJ', 'john.doe@example.com', 'pwd_john', 'John', 'Doe'),
    -- User 2
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), 'SmithJ', 'jane.smith@example.com', 'pwd_jane', 'Jane', 'Smith'),
    -- User 3
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), 'JohnsonM', 'michael.johnson@example.com', 'pwd_michael', 'Michael', 'Johnson'),
    -- User 4
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), 'BrownE', 'emily.brown@example.com', 'pwd_emily', 'Emily', 'Brown'),
    -- User 5
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), 'WilsonD', 'david.wilson@example.com', 'pwd_david', 'David', 'Wilson'),
    -- User 6
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), 'TaylorS', 'sophia.taylor@example.com', 'pwd_sophia', 'Sophia', 'Taylor'),
    -- User 7
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), 'AndersonD', 'daniel.anderson@example.com', 'pwd_daniel', 'Daniel', 'Anderson'),
    -- User 8
    (UUID_TO_BIN('fd4a096f-93f5-4f2a-86c6-69a2d20365ff'), 'ThomasO', 'olivia.thomas@example.com', 'pwd_olivia', 'Olivia', 'Thomas'),
    -- User 9
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), 'RobinsonE', 'ethan.robinson@example.com', 'pwd_ethan', 'Ethan', 'Robinson'),
    -- User 10
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), 'HarrisA', 'ava.harris@example.com', 'pwd_ava', 'Ava', 'Harris');



-- Insert statements for user_movies table
INSERT INTO user_movies (user_id, movie_id, list_type)
VALUES
    -- User 1
    -- watched_movies
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('af7c1fe6-d669-414e-b066-e9733f0de7a8'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 2
    -- watched_movies
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('08c71152-c552-42e7-b094-f510ff44e9cb'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 3
    -- watched_movies
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('c558a80a-f319-4c10-95d4-4282ef745b4b'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 4
    -- watched_movies
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('1ad1fccc-d279-46a0-8980-1d91afd6ba67'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 5
    -- watched_movies
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('5108babc-bf35-44d5-a9ba-de08badfa80a'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 6
    -- watched_movies
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('2d790a4d-7c9c-4e23-9c9c-5749c5fa7fdb'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 7
    -- watched_movies
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('8304e5ff-6324-4863-ac51-8fcbc6812b13'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 8
    -- watched_movies
    (UUID_TO_BIN('fd4a096f-93f5-4f2a-86c6-69a2d20365ff'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('fd4a096f-93f5-4f2a-86c6-69a2d20365ff'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('fd4a096f-93f5-4f2a-86c6-69a2d20365ff'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('fd4a096f-93f5-4f2a-86c6-69a2d20365ff'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),

    -- User 9
    -- watched_movies
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites'),
    (UUID_TO_BIN('96fdc209-0551-4d67-b9ad-0e9067a44bc4'), UUID_TO_BIN('6ba7b82f-9dad-11d1-80b4-00c04fd430c9'), 'favorites'),

    -- User 10
    -- watched_movies
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b826-9dad-11d1-80b4-00c04fd430c0'), 'watched_movies'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b827-9dad-11d1-80b4-00c04fd430c1'), 'watched_movies'),

    -- watchlist
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b828-9dad-11d1-80b4-00c04fd430c2'), 'watchlist'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b829-9dad-11d1-80b4-00c04fd430c3'), 'watchlist'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b82a-9dad-11d1-80b4-00c04fd430c4'), 'watchlist'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b82b-9dad-11d1-80b4-00c04fd430c5'), 'watchlist'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b82c-9dad-11d1-80b4-00c04fd430c6'), 'watchlist'),

    -- favorites
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b82d-9dad-11d1-80b4-00c04fd430c7'), 'favorites'),
    (UUID_TO_BIN('0a7d6250-0be5-4036-8f23-33dc1762bed0'), UUID_TO_BIN('6ba7b82e-9dad-11d1-80b4-00c04fd430c8'), 'favorites')
;



INSERT INTO theatres (id, name, address_id)
VALUES 
    (UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33'), 'Theatre 1', UUID_TO_BIN('6ba7b810-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f'), 'Theatre 2', UUID_TO_BIN('6ba7b811-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f'), 'Theatre 3', UUID_TO_BIN('6ba7b812-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf'), 'Theatre 4', UUID_TO_BIN('6ba7b813-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69'), 'Theatre 5', UUID_TO_BIN('6ba7b814-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a'), 'Theatre 6', UUID_TO_BIN('6ba7b815-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b'), 'Theatre 7', UUID_TO_BIN('6ba7b816-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24'), 'Theatre 8', UUID_TO_BIN('6ba7b817-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa'), 'Theatre 9', UUID_TO_BIN('6ba7b818-9dad-11d1-80b4-00c04fd430c8')),
    (UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22'), 'Theatre 10', UUID_TO_BIN('6ba7b819-9dad-11d1-80b4-00c04fd430c8'));



Insert into seat_categories (id, category_name)
values 
    (UUID_TO_BIN('b50dcee9-cd61-4fb1-a541-e0c1a4beb5d1'), 'Standard'),
    (UUID_TO_BIN('3882da2f-f73a-475e-88f6-36e8b1d91757'), 'Premium'),
    (UUID_TO_BIN('fca6845c-92d2-4e8f-98c7-88d2832f0311'), 'Couple'),
    (UUID_TO_BIN('2b31501a-0e74-4515-b1e6-e5a4c8518924'), 'Disabled')
;

-- Insert statements for CinemaHall table
-- Theatre 1
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d01'), 'Cinema Hall 1', 100, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f01'), 'Cinema Hall 2', 150, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c01'), 'Cinema Hall 3', 200, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e01'), 'Cinema Hall 4', 120, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a01'), 'Cinema Hall 5', 180, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c01'), 'Cinema Hall 6', 150, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e01'), 'Cinema Hall 7', 100, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a01'), 'Cinema Hall 8', 200, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c01'), 'Cinema Hall 9', 120, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e01'), 'Cinema Hall 10', 180, UUID_TO_BIN('cd452f0c-99f8-4176-b5be-5fccb19c0b33'));

-- Theatre 2
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d02'), 'Cinema Hall 1', 100, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f02'), 'Cinema Hall 2', 150, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c02'), 'Cinema Hall 3', 200, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e02'), 'Cinema Hall 4', 120, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a02'), 'Cinema Hall 5', 180, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c02'), 'Cinema Hall 6', 150, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e02'), 'Cinema Hall 7', 100, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a02'), 'Cinema Hall 8', 200, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c02'), 'Cinema Hall 9', 120, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e02'), 'Cinema Hall 10', 180, UUID_TO_BIN('503a5764-feaa-43d5-ad7e-f523091fbd8f'));

-- Theatre 3
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d03'), 'Cinema Hall 1', 100, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f03'), 'Cinema Hall 2', 150, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c03'), 'Cinema Hall 3', 200, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e03'), 'Cinema Hall 4', 120, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a03'), 'Cinema Hall 5', 180, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c03'), 'Cinema Hall 6', 150, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e03'), 'Cinema Hall 7', 100, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a03'), 'Cinema Hall 8', 200, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c03'), 'Cinema Hall 9', 120, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e03'), 'Cinema Hall 10', 180, UUID_TO_BIN('193e8a03-1581-46e9-b0a4-17c55fa2649f'));

-- Theatre 4
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d04'), 'Cinema Hall 1', 100, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f04'), 'Cinema Hall 2', 150, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c04'), 'Cinema Hall 3', 200, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e04'), 'Cinema Hall 4', 120, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a04'), 'Cinema Hall 5', 180, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c04'), 'Cinema Hall 6', 150, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e04'), 'Cinema Hall 7', 100, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a04'), 'Cinema Hall 8', 200, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c04'), 'Cinema Hall 9', 120, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e04'), 'Cinema Hall 10', 180, UUID_TO_BIN('112d73b4-79e5-4be8-b9ae-d0840f00d4cf'));

-- Theatre 5
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d05'), 'Cinema Hall 1', 100, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f05'), 'Cinema Hall 2', 150, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c05'), 'Cinema Hall 3', 200, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e05'), 'Cinema Hall 4', 120, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a05'), 'Cinema Hall 5', 180, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c05'), 'Cinema Hall 6', 150, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e05'), 'Cinema Hall 7', 100, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a05'), 'Cinema Hall 8', 200, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c05'), 'Cinema Hall 9', 120, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e05'), 'Cinema Hall 10', 180, UUID_TO_BIN('2729976a-fd39-40de-88f1-cb954a8a1e69'));

-- Theatre 6
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d06'), 'Cinema Hall 1', 100, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f06'), 'Cinema Hall 2', 150, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c06'), 'Cinema Hall 3', 200, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e06'), 'Cinema Hall 4', 120, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a06'), 'Cinema Hall 5', 180, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c06'), 'Cinema Hall 6', 150, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e06'), 'Cinema Hall 7', 100, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a06'), 'Cinema Hall 8', 200, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c06'), 'Cinema Hall 9', 120, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e06'), 'Cinema Hall 10', 180, UUID_TO_BIN('a90161ad-0ebf-4705-8884-155a1062898a'));

-- Theatre 7
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d07'), 'Cinema Hall 1', 100, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f07'), 'Cinema Hall 2', 150, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c07'), 'Cinema Hall 3', 200, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e07'), 'Cinema Hall 4', 120, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a07'), 'Cinema Hall 5', 180, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c07'), 'Cinema Hall 6', 150, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e07'), 'Cinema Hall 7', 100, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a07'), 'Cinema Hall 8', 200, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c07'), 'Cinema Hall 9', 120, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e07'), 'Cinema Hall 10', 180, UUID_TO_BIN('72d334b3-af8c-469f-912c-e1818b10a73b'));

-- Theatre 8
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d08'), 'Cinema Hall 1', 100, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f08'), 'Cinema Hall 2', 150, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c08'), 'Cinema Hall 3', 200, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e08'), 'Cinema Hall 4', 120, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a08'), 'Cinema Hall 5', 180, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c08'), 'Cinema Hall 6', 150, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e08'), 'Cinema Hall 7', 100, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a08'), 'Cinema Hall 8', 200, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c08'), 'Cinema Hall 9', 120, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e08'), 'Cinema Hall 10', 180, UUID_TO_BIN('3fcb9f36-4cb1-451e-9562-4c4d915a2c24'));

-- Theatre 9
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d09'), 'Cinema Hall 1', 100, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f09'), 'Cinema Hall 2', 150, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c09'), 'Cinema Hall 3', 200, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e09'), 'Cinema Hall 4', 120, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a09'), 'Cinema Hall 5', 180, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c09'), 'Cinema Hall 6', 150, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e09'), 'Cinema Hall 7', 100, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a09'), 'Cinema Hall 8', 200, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c09'), 'Cinema Hall 9', 120, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e09'), 'Cinema Hall 10', 180, UUID_TO_BIN('dfc68e6e-3025-4fef-946a-9ac1385234fa'));

-- Theatre 10
INSERT INTO cinema_halls (Id, name, capacity, theatre_id)
VALUES 
    (UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d10'), 'Cinema Hall 1', 100, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('a5f6d7c8-9b0e-4d3c-8f7a-6b5c4d3e2f10'), 'Cinema Hall 2', 150, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('e9f8d7c6-b5a4-3e2d-1c0b-9e8f7a6d5c10'), 'Cinema Hall 3', 200, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('1c2d3e4f-5a6b-7c8d-9e0f-1a2b3c4d5e10'), 'Cinema Hall 4', 120, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('9e8f7a6b-5c4d-3e2f-1a0b-9c8d7e6f5a10'), 'Cinema Hall 5', 180, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c10'), 'Cinema Hall 6', 150, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('9c8d7e6f-5a4b-3c2d-1e0f-9b8a7c6d5e10'), 'Cinema Hall 7', 100, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('1e2f3a4b-5c6d-7e8f-9a0b-1c2d3e4f5a10'), 'Cinema Hall 8', 200, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('9a8b7c6d-5e4f-3a2b-1c0d-9e8f7a6b5c10'), 'Cinema Hall 9', 120, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22')),
    (UUID_TO_BIN('1c0d2e3f-5a6b-7c8d-9e0f-1a2b3c4d5e10'), 'Cinema Hall 10', 180, UUID_TO_BIN('f22248fb-d5b2-4829-ae96-75ab59f3ff22'));


-- Insert statements for the first cinemaHall 2x2 seats for the first theatre
INSERT INTO seats (id, row_nr, column_nr, seat_category_id, cinema_hall_id)
VALUES 
    (UUID_TO_BIN('29653bcb-a57e-4ac8-a87b-c5219ba1f5cf'), 1, 1, UUID_TO_BIN('b50dcee9-cd61-4fb1-a541-e0c1a4beb5d1'), UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d01')),
    (UUID_TO_BIN('c9370954-eab1-4e99-9e97-6dc845c1b433'), 1, 2, UUID_TO_BIN('b50dcee9-cd61-4fb1-a541-e0c1a4beb5d1'), UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d01')),
    (UUID_TO_BIN('1778c52c-f2f7-4ee9-b4b4-fb15844f7343'), 2, 1, UUID_TO_BIN('b50dcee9-cd61-4fb1-a541-e0c1a4beb5d1'), UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d01')),
    (UUID_TO_BIN('fe1b8a0c-67fe-49e9-a98f-1555e702711e'), 2, 2, UUID_TO_BIN('b50dcee9-cd61-4fb1-a541-e0c1a4beb5d1'), UUID_TO_BIN('c3e8a9e5-9c2e-4e6d-9a0d-7e3a4b6c5d01'));


Insert into payment_methods (id, methodname) 
values 
    (UUID_TO_BIN('2b1f7fb2-881c-4f2d-bd46-77a48d2846c8'), 'PayPal'),
    (UUID_TO_BIN('108a70de-d64e-4b3a-a0cf-cc6190daba6e'), 'CreditCard'),
    (UUID_TO_BIN('b4a5fdd3-b467-4c61-81b4-83a243d001e2'), 'Cash'),
    (UUID_TO_BIN('4a000f56-6c84-4e6c-8601-e34419c202d1'), 'BankTransfer'),
    (UUID_TO_BIN('ce39a16d-d375-45ba-863c-0acac841f555'), 'DebitCard')
;
