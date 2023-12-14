package repositories

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type MovieActorRepositoryI interface {
	AddMovieActor(movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError
	RemoveMovieActor(movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError
}

type MovieActorRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Combine Movie and Genre
func (mar *MovieActorRepository) AddMovieActor(movieId *uuid.UUID, actorId *uuid.UUID) *models.KTSError {
	log.Print("Adding movieActor: Start")
	// Create the insert statement
	insertQuery := table.MovieActors.INSERT(table.MovieActors.MovieID, table.MovieActors.ActorID).
		VALUES(
			utils.MysqlUuid(movieId),
			utils.MysqlUuid(actorId),
		)

	log.Print("Adding movieActor: Executing query")
	// Execute the query
	rows, err := insertQuery.Exec(mar.DatabaseManager.GetDatabaseConnection())
	log.Print("Inserting movieActor Error: ", err)
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
	rows, err := deleteQuery.Exec(mar.DatabaseManager.GetDatabaseConnection())
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
