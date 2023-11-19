package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

type MovieRepo interface {
	GetMovies() []*models.Movie

	// CRUD
	GetMovieById(movieId *uuid.UUID) *models.Movie
	CreateMovie(movie *models.Movie) *models.Movie
	UpdateMovie(movie *models.Movie) *models.Movie
	DeleteMovie(movieId *uuid.UUID) *models.Movie

	// Validation
	ValidateIfMovieExists(movieId *uuid.UUID) bool

	GetActorsFromMovie(movieId *uuid.UUID) []*models.Actor
	GetProducerFromMovie(movieId *uuid.UUID) []*models.Producer
}

type MovieRepository struct {
	DatabaseMgr managers.DatabaseManagerI
}

func (mr *MovieRepository) GetMovies() []*models.Movie {
	query := `SELECT m.id, m.title, m.description, m.releasedDate, m.timeInMin, f.age , g.name from movies m inner join genres g on m.genre_id = g.id inner join fsk f on m.fsk_id = f.id;`
	rows, err := mr.DatabaseMgr.ExecuteQuery(query)
	if err != nil {
		log.Printf("Error while querying movies: %v", err)
		return nil
	}
	defer rows.Close()

	return rowsToMovieSchema(rows)
}

func (mr *MovieRepository) GetMovieById(movieId *uuid.UUID) *models.Movie {
	// Maybe use procedure
	query := `SELECT m.id, m.title, m.description, m.releasedDate, m.timeInMin, f.age , g.name from movies m inner join genres g on m.genre_id = g.id inner join fsk f on m.fsk_id = f.id where m.id = ?;` + movieId.String()
	row := mr.DatabaseMgr.ExecuteQueryRow(query)
	return rowToMovieSchema(row)
}

func (mr *MovieRepository) CreateMovie(movie *models.Movie) {
	// TODO need procedure because of fsk and genre id's
	insert_values := fmt.Sprintf("%s, %s, %s, %d, %s, %s", movie.Title, movie.Description, movie.ReleaseDate, movie.TimeInMin, movie.FSK, movie.Genre)
	query := "INSERT INTO movies (title, description, releasedDate, timeInMin, fsk_id, genre_id) VALUES (" + insert_values + ");"
	result, err := mr.DatabaseMgr.ExecuteStatement(query)
	if err != nil {
		log.Printf("Error while inserting movie: %v", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("Error while inserting movie: %v", err)
		return
	}
	return
}

func (mr *MovieRepository) UpdateMovie(movie *models.Movie) {
	// Update movie
	// TODO need procedure because of fsk and genre id's
	updateString := fmt.Sprintf("UPDATE movies SET title = %s, description = %s, releasedDate = %s, timeInMin = %d, fsk_id = %s, genre_id = %s WHERE id = %s", movie.Title, movie.Description, movie.ReleaseDate, movie.TimeInMin, movie.FSK, movie.Genre, movie.Id)
	result, err := mr.DatabaseMgr.ExecuteStatement(updateString)
	if err != nil {
		log.Printf("Error while updating movie: %v", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("Error while updating movie: %v", err)
		return
	}

	return
}

func (mr *MovieRepository) DeleteMovie(movieId *uuid.UUID) {
	result, err := mr.DatabaseMgr.ExecuteStatement(fmt.Sprintf("DELETE FROM movies WHERE id = %s", movieId))
	if err != nil {
		log.Printf("Error while deleting movie: %v", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("Error while deleting movie: %v", err)
		return
	}

	return
}

func (mr *MovieRepository) ValidateIfMovieExists(movieId *uuid.UUID) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM movies WHERE id = %s", movieId)
	row := mr.DatabaseMgr.ExecuteQueryRow(query)

	var count int
	if err := row.Scan(&count); err != nil {
		log.Printf("Error while scanning movie: %v", err)
		return false
	}

	return count > 0
}

func (mr *MovieRepository) GetActorsFromMovie(movieId *uuid.UUID) []*models.Actor {
	query := fmt.Sprintf("select m.title, a.name from movie_actors as ma inner join movies m on ma.movie_id = m.id inner join actors a on ma.actor_id = a.id where m.id = %s;", movieId)
	rows, err := mr.DatabaseMgr.ExecuteQuery(query)
	if err != nil {
		log.Printf("Error while querying actors: %v", err)
		return nil
	}
	defer rows.Close()

	return rowsToActorSchema(rows)
}

func (mr *MovieRepository) GetProducersFromMovie(movieId *uuid.UUID) []*models.Producer {
	query := fmt.Sprintf("select m.title, a.name from movie_producers as mp inner join movies m on mp.movie_id = m.id inner join producers p on mp.actor_id = p.id where m.id = %s;", movieId)
	rows, err := mr.DatabaseMgr.ExecuteQuery(query)
	if err != nil {
		log.Printf("Error while querying producers: %v", err)
		return nil
	}
	defer rows.Close()

	return rowsToProducerSchema(rows)
}

// rowsToMovieSchema converts a set of rows to a slice of MovieSchema
func rowsToMovieSchema(rows *sql.Rows) []*models.Movie {
	movies := make([]*models.Movie, 0) // It is important to initialize the slice with 0 length so that it is serialized to [] instead of null
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.TimeInMin, &movie.FSK, &movie.Genre)
		if err != nil {
			log.Printf("Error while scanning movie: %v", err)
			return nil
		}
		movies = append(movies, &movie)
	}

	return movies
}

// rowToMovieSchema converts a row to a MovieSchema
func rowToMovieSchema(row *sql.Row) *models.Movie {
	movie := models.Movie{}
	if err := row.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.TimeInMin, &movie.FSK, &movie.Genre); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Printf("Error while scanning movie: %v", err)
		return nil
	}

	return &movie
}

// Method will be moved to actor repository
func rowsToActorSchema(rows *sql.Rows) []*models.Actor {
	actors := make([]*models.Actor, 0)
	for rows.Next() {
		var actor models.Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.Age)
		if err != nil {
			log.Printf("Error while scanning movie: %v", err)
			return nil
		}
		actors = append(actors, &actor)
	}

	return actors
}

// Method will be moved to actor repository
func rowsToProducerSchema(rows *sql.Rows) []*models.Producer {
	producers := make([]*models.Producer, 0)
	for rows.Next() {
		var producer models.Producer
		err := rows.Scan(&producer.ID, &producer.Name, &producer.Age)
		if err != nil {
			log.Printf("Error while scanning movie: %v", err)
			return nil
		}
		producers = append(producers, &producer)
	}

	return producers
}
