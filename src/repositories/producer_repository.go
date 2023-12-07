package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	jet_mysql "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type ProducerRepositoryI interface {
	GetProducers() (*[]model.Producers, *models.KTSError)
	GetProducerById(id *uuid.UUID) (*model.Producers, *models.KTSError)
	CreateProducer(producer *model.Producers) *models.KTSError
	UpdateProducer(producer *model.Producers) *models.KTSError
	DeleteProducer(id *uuid.UUID) *models.KTSError
}

type ProducerRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (mr *MovieRepository) GetProducers() (*[]model.Producers, *models.KTSError) {
	var producers []model.Producers

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Producers.AllColumns,
	).FROM(
		table.Producers,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &producers)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(producers) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &producers, nil
}

func (mr *MovieRepository) GetProducerById(id *uuid.UUID) (*model.Producers, *models.KTSError) {
	var producer model.Producers

	binary_id, _ := id.MarshalBinary()

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Producers.AllColumns,
	).FROM(
		table.Producers,
	).WHERE(
		table.Producers.ID.EQ(jet_mysql.String(string(binary_id))),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &producer)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &producer, nil
}

func (mr *MovieRepository) CreateProducer(producer *model.Producers) *models.KTSError {
	// Create the query
	stmt := table.Producers.INSERT(
		table.Producers.AllColumns,
	).VALUES(
		table.Producers.ID.SET(jet_mysql.String(producer.ID.String())),
		table.Producers.Name.SET(jet_mysql.String(producer.Name)),
		table.Producers.Birthdate.SET(jet_mysql.Date(producer.Birthdate.Year(), producer.Birthdate.Month(), producer.Birthdate.Day())),
		table.Producers.Description.SET(jet_mysql.String(producer.Description)),
		table.Producers.PicURL.SET(jet_mysql.String((*producer.PicURL))),
	)

	// Execute the query
	rows, err := stmt.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (mr *MovieRepository) UpdateProducer(producer *model.Producers) *models.KTSError {
	// Create the query
	stmt := table.Producers.UPDATE(
		table.Producers.AllColumns,
	).SET(
		table.Producers.ID.SET(jet_mysql.String(producer.ID.String())),
		table.Producers.Name.SET(jet_mysql.String(producer.Name)),
		table.Producers.Birthdate.SET(jet_mysql.Date(producer.Birthdate.Year(), producer.Birthdate.Month(), producer.Birthdate.Day())),
		table.Producers.Description.SET(jet_mysql.String(producer.Description)),
		table.Producers.PicURL.SET(jet_mysql.String((*producer.PicURL))),
	).WHERE(
		table.Producers.ID.EQ(jet_mysql.String(producer.ID.String())),
	)

	// Execute the query
	rows, err := stmt.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (mr *MovieRepository) DeleteProducer(id *uuid.UUID) *models.KTSError {
	// Create the query
	stmt := table.Producers.DELETE().WHERE(table.Producers.ID.EQ(jet_mysql.String(id.String())))

	// Execute the query
	rows, err := stmt.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
