USE KinoTicketSystem;


DELIMITER //
CREATE PROCEDURE insertMovie (IN movieTitle VARCHAR(255), IN movieDescription TEXT, IN movieReleaseDate DATE, IN movieTimeInMin INT, IN movieFskId INT, IN movieGenreId varchar(100))
BEGIN
    DECLARE m_fsk_id int;
    DECLARE m_genre_id int;
    
    Select id into m_fsk_id from fsk where age = movieFskId;
    Select id into m_genre_id from genres where name = movieGenreId;
    
    INSERT INTO movies (title, description, releasedDate, timeInMin, fsk_id, genre_id)
    VALUES (movieTitle, movieDescription, movieReleaseDate, movieTimeInMin, m_fsk_id, m_genre_id);
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE getMovieById (IN movieId int)
BEGIN
    SELECT m.id, m.title, m.description, m.releasedDate, m.timeInMin, f.age , g.name 
    from movies m 
    inner join genres g 
        on m.genre_id = g.id 
    inner join fsk f 
        on m.fsk_id = f.id 
    where m.id = movieId;
END //
DELIMITER ;