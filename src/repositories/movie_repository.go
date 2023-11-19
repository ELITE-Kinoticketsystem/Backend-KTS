package repositories

import (
	"context"
	"database/sql"

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
	/*
		trips := make([]*models.TripSchema, 0) // It is important to initialize the slice with 0 length so that it is serialized to [] instead of null
		for rows.Next() {
			var trip models.TripSchema
			err := rows.Scan(&trip.TripID, &trip.Name, &trip.Description, &trip.Location, &trip.StartDate, &trip.EndDate)
			if err != nil {
				log.Printf("Error while scanning trip: %v", err)
				return nil, expense_errors.EXPENSE_INTERNAL_ERROR
			}
			trips = append(trips, &trip)
		}

		return trips, nil
	*/
	return nil
}
