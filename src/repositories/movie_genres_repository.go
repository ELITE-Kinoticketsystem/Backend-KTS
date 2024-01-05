package repositories

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type MovieGenreRepositoryI interface {
	// Combine Movie and Genre
	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
	RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
	RemoveAllGenreCombinationWithMovie(movieId *uuid.UUID) *models.KTSError
	RemoveAllMovieCombinationWithGenre(genreId *uuid.UUID) *models.KTSError
}

type MovieGenreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Combine Movie and Genre
func (mgr *MovieGenreRepository) AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError {
	// Create the insert statement
	insertQuery := table.MovieGenres.INSERT(table.MovieGenres.MovieID, table.MovieGenres.GenreID).
		VALUES(utils.MysqlUuid(movieId), utils.MysqlUuid(genreId))

	// Execute the query
	rows, err := insertQuery.Exec(mgr.DatabaseManager.GetDatabaseConnection())
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

func (mgr *MovieGenreRepository) RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError {

	deleteQuery := table.MovieGenres.DELETE().WHERE(
		table.MovieGenres.MovieID.EQ(utils.MysqlUuid(movieId)).AND(
			table.MovieGenres.GenreID.EQ(utils.MysqlUuid(genreId)),
		),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(mgr.DatabaseManager.GetDatabaseConnection())
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

func (mgr *MovieGenreRepository) RemoveAllGenreCombinationWithMovie(movieId *uuid.UUID) *models.KTSError {
	deleteQuery := table.MovieGenres.DELETE().WHERE(
		table.MovieGenres.MovieID.EQ(utils.MysqlUuid(movieId)),
	)

	// Execute the query
	_, err := deleteQuery.Exec(mgr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil

}

func (mgr *MovieGenreRepository) RemoveAllMovieCombinationWithGenre(genreId *uuid.UUID) *models.KTSError {
	deleteQuery := table.MovieGenres.DELETE().WHERE(
		table.MovieGenres.GenreID.EQ(utils.MysqlUuid(genreId)),
	)

	// Execute the query
	_, err := deleteQuery.Exec(mgr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil

}