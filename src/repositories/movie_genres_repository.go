package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	jet_mysql "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type MovieGenreRepositoryI interface {
	// Combine Movie and Genre
	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
	RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
}

type MovieGenreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Combine Movie and Genre
func (mr *MovieRepository) AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError {

	binary_movie_id, _ := movieId.MarshalBinary()
	binary_genre_id, _ := genreId.MarshalBinary()

	// Create the insert statement
	insertQuery := table.MovieGenres.INSERT(table.MovieGenres.MovieID, table.MovieGenres.GenreID).
		VALUES(jet_mysql.String(string(binary_movie_id)), jet_mysql.String(string(binary_genre_id)))

	// Execute the query
	_, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (mr *MovieRepository) RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError {

	// Create the delete statement
	deleteQuery := table.MovieGenres.DELETE().WHERE(
		table.MovieGenres.MovieID.EQ(
			jet_mysql.CAST(jet_mysql.String(movieId.String())).AS_BINARY(),
		).AND(
			table.MovieGenres.GenreID.EQ(
				jet_mysql.CAST(jet_mysql.String(genreId.String())).AS_BINARY(),
			),
		),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
