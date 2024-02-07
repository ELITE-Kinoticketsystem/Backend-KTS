package repositories

import (
	"database/sql"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type MovieActorRepositoryI interface {
	AddMovieActor(tx *sql.Tx, movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError
	RemoveMovieActor(movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError
	RemoveAllActorCombinationWithMovie(movieId *uuid.UUID) *models.KTSError
}

type MovieActorRepository struct {
	managers.DatabaseManagerI
}

// Combine Movie and Genre
func (mar *MovieActorRepository) AddMovieActor(tx *sql.Tx, movieId *uuid.UUID, actorId *uuid.UUID) 
				*models.KTSError {
	// Create the insert statement
	insertQuery := table.MovieActors.INSERT(
		table.MovieActors.MovieID, table.MovieActors.ActorID
	).VALUES(
		utils.MysqlUuid(movieId),utils.MysqlUuid(actorId),
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
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

func (mar *MovieActorRepository) RemoveMovieActor(movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError {

	deleteQuery := table.MovieGenres.DELETE().WHERE(
		table.MovieActors.MovieID.EQ(utils.MysqlUuid(movieId)).AND(
			table.MovieActors.ActorID.EQ(utils.MysqlUuid(actorId)),
		),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(mar.GetDatabaseConnection())
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

func (mar *MovieActorRepository) RemoveAllActorCombinationWithMovie(movieId *uuid.UUID) *models.KTSError {
	deleteQuery := table.MovieActors.DELETE().WHERE(
		table.MovieActors.MovieID.EQ(utils.MysqlUuid(movieId)),
	)

	// Execute the query
	_, err := deleteQuery.Exec(mar.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}
