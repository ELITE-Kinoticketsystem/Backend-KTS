package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

type MovieRepo interface {
	GetMovies(ctx context.Context) []*models.Movie

	// CRUD
	GetMovieById(ctx context.Context, movieId *uuid.UUID) *models.Movie
	CreateMovie(ctx context.Context, movie *models.Movie) *models.Movie
	UpdateMovie(ctx context.Context, movie *models.Movie) *models.Movie
	DeleteMovie(ctx context.Context, movieId *uuid.UUID) *models.Movie

	// Validation
	ValidateIfMovieExists(ctx context.Context, movieId *uuid.UUID) bool

	GetActorsFromMovie(ctx context.Context, movieId *uuid.UUID) []*models.Actor
	GetProducerFromMovie(ctx context.Context, movieId *uuid.UUID) *models.Producer
}

type MovieRepository struct {
	DatabaseMgr managers.DatabaseManagerI
}

func (mr *MovieRepository) GetMovies(ctx context.Context) []*models.Movie {
	return nil
}

func (mr *MovieRepository) GetMovieById(ctx context.Context, movieId *uuid.UUID) *models.Movie {
	// Maybe use procedure
	query := `SELECT m.id, m.title, m.description, m.releasedDate, m.timeInMin, f.age , g.name from movies m inner join genres g on m.genre_id = g.id inner join fsk f on m.fsk_id = f.id where m.id = ?;` + movieId.String()
	row := mr.DatabaseMgr.ExecuteQueryRow(query)
	// return rowToTripSchema(row)
	return rowToMovieSchema(row)
}

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
