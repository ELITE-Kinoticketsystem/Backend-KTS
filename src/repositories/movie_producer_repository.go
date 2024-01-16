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

type MovieProducerRepositoryI interface {
	// Combine Movie and Actor
	AddMovieProducer(tx *sql.Tx, movieId *uuid.UUID, producerId *uuid.UUID) *models.KTSError
	RemoveMovieProducer(movieId *uuid.UUID, producerId *uuid.UUID) *models.KTSError
	RemoveAllProducerCombinationWithMovie(movieId *uuid.UUID) *models.KTSError
}

type MovieProducerRepository struct {
	managers.DatabaseManagerI
}

// Combine Movie and Genre
func (pr *MovieProducerRepository) AddMovieProducer(tx *sql.Tx, movieId *uuid.UUID, producerId *uuid.UUID) *models.KTSError {

	// Create the insert statement
	insertQuery := table.MovieProducers.INSERT(table.MovieProducers.MovieID, table.MovieProducers.ProducerID).
		VALUES(utils.MysqlUuid(movieId), utils.MysqlUuid(producerId))

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

func (pr *MovieProducerRepository) RemoveMovieProducer(movieId *uuid.UUID, producerId *uuid.UUID) *models.KTSError {

	deleteQuery := table.MovieProducers.DELETE().WHERE(
		table.MovieProducers.MovieID.EQ(utils.MysqlUuid(movieId)).AND(
			table.MovieProducers.ProducerID.EQ(utils.MysqlUuid(producerId)),
		),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(pr.GetDatabaseConnection())
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

func (pr *MovieProducerRepository) RemoveAllProducerCombinationWithMovie(movieId *uuid.UUID) *models.KTSError {
	deleteQuery := table.MovieProducers.DELETE().WHERE(
		table.MovieProducers.MovieID.EQ(utils.MysqlUuid(movieId)),
	)

	// Execute the query
	_, err := deleteQuery.Exec(pr.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}
