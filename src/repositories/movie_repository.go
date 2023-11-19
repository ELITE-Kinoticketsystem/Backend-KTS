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
	GetProducerFromMovie(movieId *uuid.UUID) *models.Producer
}

type MovieRepository struct {
	DatabaseMgr managers.DatabaseManagerI
}

func (mr *MovieRepository) GetMovies() []*models.Movie {
	query := `SELECT m.id, m.title, m.description, m.releasedDate, m.timeInMin, f.age , g.name from movies m inner join genres g on m.genre_id = g.id inner join fsk f on m.fsk_id = f.id;`
	rows, err := mr.DatabaseMgr.ExecuteQuery(query)
	if err != nil {
		log.Printf("Error while querying trips: %v", err)
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
		log.Printf("Error while inserting trip: %v", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("Error while inserting trip: %v", err)
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
		log.Printf("Error while updating trip: %v", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("Error while updating trip: %v", err)
		return
	}

	return
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
