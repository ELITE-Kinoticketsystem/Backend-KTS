package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	jet_mysql "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type ProducerRepositoryI interface {
	GetProducers() (*[]models.GetProducersDTO, *models.KTSError)
	GetProducerById(id *uuid.UUID) (*models.ProducerDTO, *models.KTSError)
	CreateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError)
	CreateProducerPicture(producerPicture *model.ProducerPictures) (*uuid.UUID, *models.KTSError)
	UpdateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError)
	DeleteProducer(id *uuid.UUID) *models.KTSError
}

type ProducerRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (pr *ProducerRepository) GetProducers() (*[]models.GetProducersDTO, *models.KTSError) {
	var producers []models.GetProducersDTO

	// Create the query
	stmt := jet_mysql.SELECT(
		table.Producers.AllColumns,
		table.ProducerPictures.AllColumns,
	).
		FROM(
			table.Producers.
				LEFT_JOIN(table.ProducerPictures, table.ProducerPictures.ProducerID.EQ(table.Producers.ID)),
		)

	// Execute the query
	err := stmt.Query(pr.DatabaseManager.GetDatabaseConnection(), &producers)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(producers) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &producers, nil
}

func (pr *ProducerRepository) GetProducerById(id *uuid.UUID) (*models.ProducerDTO, *models.KTSError) {
	var producer models.ProducerDTO
	// Create the query
	stmt := jet_mysql.SELECT(
		table.Producers.AllColumns,
		table.ProducerPictures.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(
			table.Producers.
				LEFT_JOIN(table.ProducerPictures, table.ProducerPictures.ProducerID.EQ(table.Producers.ID)).
				LEFT_JOIN(table.MovieProducers, table.MovieProducers.ProducerID.EQ(table.Producers.ID)).
				LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.MovieProducers.MovieID)),
		).
		WHERE(
			table.Actors.ID.EQ(utils.MysqlUuid(id)),
		)

	// Execute the query
	err := stmt.Query(pr.DatabaseManager.GetDatabaseConnection(), &producer)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &producer, nil
}

func (pr *ProducerRepository) CreateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError) {
	producer.ID = utils.NewUUID()

	// Create the query
	stmt := table.Producers.INSERT(
		table.Producers.AllColumns,
	).VALUES(
		producer.ID,
		producer.Name,
		producer.Birthdate,
		producer.Description,
		producer.PicURL,
	)

	// Execute the query
	rows, err := stmt.Exec(pr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return producer.ID, nil
}

func (pr *ProducerRepository) CreateProducerPicture(producerPicture *model.ProducerPictures) (*uuid.UUID, *models.KTSError) {
	producerPicture.ID = utils.NewUUID()

	insertStmt := table.ProducerPictures.INSERT(table.ProducerPictures.AllColumns).VALUES(
		utils.MysqlUuid(producerPicture.ID),
		utils.MysqlUuid(producerPicture.ProducerID),
		producerPicture.PicURL,
	)

	rows, err := insertStmt.Exec(pr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return producerPicture.ProducerID, nil
}

func (pr *ProducerRepository) UpdateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError) {

	// binary_id, _ := producer.ID.MarshalBinary()

	// // Create the query
	// stmt := table.Producers.UPDATE(
	// 	table.Producers.AllColumns,
	// ).SET(
	// 	table.Producers.Name.SET(jet_mysql.String(producer.Name)),
	// 	table.Producers.Birthdate.SET(jet_mysql.Date(producer.Birthdate.Year(), producer.Birthdate.Month(), producer.Birthdate.Day())),
	// 	table.Producers.Description.SET(jet_mysql.String(producer.Description)),
	// 	table.Producers.PicURL.SET(jet_mysql.String((*producer.PicURL))),
	// ).WHERE(
	// 	table.Producers.ID.EQ(jet_mysql.String(string(binary_id))),
	// )

	// // Execute the query
	// rows, err := stmt.Exec(pr.DatabaseManager.GetDatabaseConnection())
	// if err != nil {
	// 	return kts_errors.KTS_INTERNAL_ERROR
	// }

	// rowsAffected, err := rows.RowsAffected()
	// if err != nil {
	// 	return kts_errors.KTS_INTERNAL_ERROR
	// }

	// if rowsAffected == 0 {
	// 	return kts_errors.KTS_NOT_FOUND
	// }

	return nil, nil
}

func (pr *ProducerRepository) DeleteProducer(id *uuid.UUID) *models.KTSError {
	// // Create the query
	// stmt := table.Producers.DELETE().WHERE(table.Producers.ID.EQ(jet_mysql.String(id.String())))

	// // Execute the query
	// rows, err := stmt.Exec(pr.DatabaseManager.GetDatabaseConnection())
	// if err != nil {
	// 	return kts_errors.KTS_INTERNAL_ERROR
	// }

	// rowsAffected, err := rows.RowsAffected()
	// if err != nil {
	// 	return kts_errors.KTS_INTERNAL_ERROR
	// }

	// if rowsAffected == 0 {
	// 	return kts_errors.KTS_NOT_FOUND
	// }

	return nil
}
