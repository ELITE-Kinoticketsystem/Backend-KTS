package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetProducers(t *testing.T) {
	sampleProducers := utils.GetSampleProducers()

	query := "SELECT producers.id AS \"producers.id\", producers.name AS \"producers.name\", producers.birthdate AS \"producers.birthdate\", producers.description AS \"producers.description\", producers.pic_url AS \"producers.pic_url\", producer_pictures.id AS \"producer_pictures.id\", producer_pictures.producer_id AS \"producer_pictures.producer_id\", producer_pictures.pic_url AS \"producer_pictures.pic_url\" FROM `KinoTicketSystem`.producers LEFT JOIN `KinoTicketSystem`.producer_pictures ON (producer_pictures.producer_id = producers.id);"

	testCases := []struct {
		name             string
		setExpectations  func(mock sqlmock.Sqlmock)
		expectedProducer *[]models.GetProducersDTO
		expectedError    *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url", "producer_pictures.id", "producer_pictures.producer_id", "producer_pictures.pic_url"}),
				)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple genres",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url", "producer_pictures.id", "producer_pictures.producer_id", "producer_pictures.pic_url"},
					).AddRow(
						(*sampleProducers)[0].ID, (*sampleProducers)[0].Name, (*sampleProducers)[0].Birthdate, (*sampleProducers)[0].Description, (*sampleProducers)[0].PicURL, (*sampleProducers)[0].Pictures[0].ID, (*sampleProducers)[0].Pictures[0].ProducerID, (*sampleProducers)[0].Pictures[0].PicURL,
					).AddRow(
						(*sampleProducers)[0].ID, (*sampleProducers)[0].Name, (*sampleProducers)[0].Birthdate, (*sampleProducers)[0].Description, (*sampleProducers)[0].PicURL, (*sampleProducers)[0].Pictures[1].ID, (*sampleProducers)[0].Pictures[1].ProducerID, (*sampleProducers)[0].Pictures[1].PicURL,
					).AddRow(
						(*sampleProducers)[1].ID, (*sampleProducers)[1].Name, (*sampleProducers)[1].Birthdate, (*sampleProducers)[1].Description, (*sampleProducers)[1].PicURL, (*sampleProducers)[1].Pictures[0].ID, (*sampleProducers)[1].Pictures[0].ProducerID, (*sampleProducers)[1].Pictures[0].PicURL,
					).AddRow(
						(*sampleProducers)[1].ID, (*sampleProducers)[1].Name, (*sampleProducers)[1].Birthdate, (*sampleProducers)[1].Description, (*sampleProducers)[1].PicURL, (*sampleProducers)[1].Pictures[1].ID, (*sampleProducers)[1].Pictures[1].ProducerID, (*sampleProducers)[1].Pictures[1].PicURL,
					),
				)
			},
			expectedProducer: sampleProducers,
			expectedError:    nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the genreRepository with the mock database connection
			producerRepo := ProducerRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			producers, kts_err := producerRepo.GetProducers()

			// Verify the results
			assert.Equal(t, tc.expectedProducer, producers)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetProducerById(t *testing.T) {
	sampleProducers := utils.GetSampleProducerDTO()

	query := "SELECT producers.id AS \"producers.id\", producers.name AS \"producers.name\", producers.birthdate AS \"producers.birthdate\", producers.description AS \"producers.description\", producers.pic_url AS \"producers.pic_url\", producer_pictures.id AS \"producer_pictures.id\", producer_pictures.producer_id AS \"producer_pictures.producer_id\", producer_pictures.pic_url AS \"producer_pictures.pic_url\", movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.producers LEFT JOIN `KinoTicketSystem`.producer_pictures ON (producer_pictures.producer_id = producers.id) LEFT JOIN `KinoTicketSystem`.movie_producers ON (movie_producers.producer_id = producers.id) LEFT JOIN `KinoTicketSystem`.movies ON (movies.id = movie_producers.movie_id) WHERE actors.id = ?;"

	producer_id := sampleProducers.ID

	testCases := []struct {
		name             string
		setExpectations  func(mock sqlmock.Sqlmock, producerId *uuid.UUID)
		expectedProducer *models.ProducerDTO
		expectedError    *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(producer_id)).WillReturnRows(
					sqlmock.NewRows([]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url", "producer_pictures.id", "producer_pictures.producer_id", "producer_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
				)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single producer",
			setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(producer_id)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url", "producer_pictures.id", "producer_pictures.producer_id", "producer_pictures.pic_url", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						&sampleProducers.ID, &sampleProducers.Name, &sampleProducers.Birthdate, &sampleProducers.Description, &sampleProducers.PicURL, &sampleProducers.Pictures[0].ID, &sampleProducers.Pictures[0].ProducerID, &sampleProducers.Pictures[0].PicURL, &sampleProducers.Movies[0].ID, &sampleProducers.Movies[0].Title, &sampleProducers.Movies[0].Description, &sampleProducers.Movies[0].BannerPicURL, &sampleProducers.Movies[0].CoverPicURL, &sampleProducers.Movies[0].TrailerURL, &sampleProducers.Movies[0].Rating, &sampleProducers.Movies[0].ReleaseDate, &sampleProducers.Movies[0].TimeInMin, &sampleProducers.Movies[0].Fsk,
					).AddRow(
						&sampleProducers.ID, &sampleProducers.Name, &sampleProducers.Birthdate, &sampleProducers.Description, &sampleProducers.PicURL, &sampleProducers.Pictures[1].ID, &sampleProducers.Pictures[1].ProducerID, &sampleProducers.Pictures[1].PicURL, &sampleProducers.Movies[1].ID, &sampleProducers.Movies[1].Title, &sampleProducers.Movies[1].Description, &sampleProducers.Movies[1].BannerPicURL, &sampleProducers.Movies[1].CoverPicURL, &sampleProducers.Movies[1].TrailerURL, &sampleProducers.Movies[1].Rating, &sampleProducers.Movies[1].ReleaseDate, &sampleProducers.Movies[1].TimeInMin, &sampleProducers.Movies[1].Fsk,
					).AddRow(
						&sampleProducers.ID, &sampleProducers.Name, &sampleProducers.Birthdate, &sampleProducers.Description, &sampleProducers.PicURL, &sampleProducers.Pictures[0].ID, &sampleProducers.Pictures[0].ProducerID, &sampleProducers.Pictures[0].PicURL, &sampleProducers.Movies[0].ID, &sampleProducers.Movies[0].Title, &sampleProducers.Movies[0].Description, &sampleProducers.Movies[0].BannerPicURL, &sampleProducers.Movies[0].CoverPicURL, &sampleProducers.Movies[0].TrailerURL, &sampleProducers.Movies[0].Rating, &sampleProducers.Movies[0].ReleaseDate, &sampleProducers.Movies[0].TimeInMin, &sampleProducers.Movies[0].Fsk,
					).AddRow(
						&sampleProducers.ID, &sampleProducers.Name, &sampleProducers.Birthdate, &sampleProducers.Description, &sampleProducers.PicURL, &sampleProducers.Pictures[1].ID, &sampleProducers.Pictures[1].ProducerID, &sampleProducers.Pictures[1].PicURL, &sampleProducers.Movies[1].ID, &sampleProducers.Movies[1].Title, &sampleProducers.Movies[1].Description, &sampleProducers.Movies[1].BannerPicURL, &sampleProducers.Movies[1].CoverPicURL, &sampleProducers.Movies[1].TrailerURL, &sampleProducers.Movies[1].Rating, &sampleProducers.Movies[1].ReleaseDate, &sampleProducers.Movies[1].TimeInMin, &sampleProducers.Movies[1].Fsk,
					),
				)
			},
			expectedProducer: sampleProducers,
			expectedError:    nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(producer_id)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the genreRepository with the mock database connection
			producerRepo := ProducerRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, producer_id)

			// Call the method under test
			producer, kts_err := producerRepo.GetProducerById(producer_id)

			// Verify the results
			assert.Equal(t, tc.expectedProducer, producer)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreateProducer(t *testing.T) {
	sampleProducer := utils.GetSampleProducer()

	query := "INSERT INTO `KinoTicketSystem`.producers (id, name, birthdate, description, pic_url) VALUES (?, ?, ?, ?, ?);"

	testCases := []struct {
		name               string
		setExpectations    func(mock sqlmock.Sqlmock, producer *model.Producers)
		expectedProducerId bool
		expectedError      *models.KTSError
	}{
		{
			name: "Successful creation",
			setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedProducerId: true,
			expectedError:      nil,
		},
		{
			name: "Error while creating producer",
			setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedProducerId: false,
			expectedError:      kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedProducerId: false,
			expectedError:      kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Producer not found",
			setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedProducerId: false,
			expectedError:      kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the genreRepository with the mock database connection
			producerRepo := ProducerRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleProducer)

			// Call the method under test
			producerId, kts_err := producerRepo.CreateProducer(sampleProducer)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedProducerId && producerId == nil {
				t.Error("Expected actor ID, got nil")
			}

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUpdateProducer(t *testing.T) {
	// TODO

	// sampleProducer := utils.GetSampleProducer()

	// query := "UPDATE `KinoTicketSystem`.producers SET name = ?, birthdate = CAST(? AS DATE), description = ?, pic_url = ? WHERE producers.id = ?;"

	// testCases := []struct {
	// 	name            string
	// 	setExpectations func(mock sqlmock.Sqlmock, producer *model.Producers)
	// 	expectedError   *models.KTSError
	// }{
	// 	{
	// 		name: "Successful update",
	// 		setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
	// 			mock.ExpectExec(query).WithArgs(sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL, utils.EqUUID(sampleProducer.ID)).WillReturnResult(sqlmock.NewResult(0, 0))
	// 		},
	// 		expectedError: nil,
	// 	},
	// 	// {
	// 	// 	name: "Error while updating producer",
	// 	// 	setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
	// 	// 		mock.ExpectExec(query).WithArgs(sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL, sampleProducer.ID).WillReturnError(sqlmock.ErrCancelled)
	// 	// 	},
	// 	// 	expectedError: kts_errors.KTS_INTERNAL_ERROR,
	// 	// },
	// 	// {
	// 	// 	name: "Error while converting rows affected",
	// 	// 	setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
	// 	// 		mock.ExpectExec(query).WithArgs(sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL, sampleProducer.ID).WillReturnResult(
	// 	// 			sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
	// 	// 		)
	// 	// 	},
	// 	// 	expectedError: kts_errors.KTS_INTERNAL_ERROR,
	// 	// },
	// 	// {
	// 	// 	name: "Producer not found",
	// 	// 	setExpectations: func(mock sqlmock.Sqlmock, producer *model.Producers) {
	// 	// 		mock.ExpectExec(query).WithArgs(sampleProducer.Name, sampleProducer.Birthdate, sampleProducer.Description, sampleProducer.PicURL, sampleProducer.ID).WillReturnResult(sqlmock.NewResult(1, 0))
	// 	// 	},
	// 	// 	expectedError: kts_errors.KTS_NOT_FOUND,
	// 	// },
	// }

	// for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		// Create a new mock database connection
	// 		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	// 		if err != nil {
	// 			t.Fatalf("Failed to create mock database connection: %v", err)
	// 		}
	// 		defer db.Close()

	// 		// Create a new instance of the genreRepository with the mock database connection
	// 		producerRepo := ProducerRepository{
	// 			DatabaseManager: &managers.DatabaseManager{
	// 				Connection: db,
	// 			},
	// 		}

	// 		tc.setExpectations(mock, sampleProducer)

	// 		// Call the method under test
	// 		kts_err := producerRepo.UpdateProducer(sampleProducer)

	// 		// Verify the results
	// 		assert.Equal(t, tc.expectedError, kts_err)

	// 		// Verify that all expectations were met
	// 		if err = mock.ExpectationsWereMet(); err != nil {
	// 			t.Errorf("There were unfulfilled expectations: %s", err)
	// 		}
	// 	})
	// }
	t.Skip("Not implemented")
}

func TestDeleteProducer(t *testing.T) {
	// TODO

	// sampleProducer := utils.GetSampleProducer()

	// producerId := sampleProducer.ID

	// query := "DELETE FROM `KinoTicketSystem`.producers WHERE producers.id = ?;"

	// testCases := []struct {
	// 	name            string
	// 	setExpectations func(mock sqlmock.Sqlmock, producerId *uuid.UUID)
	// 	expectedError   *models.KTSError
	// }{
	// 	{
	// 		name: "Successful deletion",
	// 		setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
	// 			mock.ExpectExec(query).WithArgs(utils.EqUUID(producerId)).WillReturnResult(sqlmock.NewResult(1, 1))
	// 		},
	// 		expectedError: nil,
	// 	},
	// 	{
	// 		name: "Error while deleting producer",
	// 		setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
	// 			mock.ExpectExec(query).WithArgs(utils.EqUUID(producerId)).WillReturnError(sqlmock.ErrCancelled)
	// 		},
	// 		expectedError: kts_errors.KTS_INTERNAL_ERROR,
	// 	},
	// 	{
	// 		name: "Error while converting rows affected",
	// 		setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
	// 			mock.ExpectExec(query).WithArgs(utils.EqUUID(producerId)).WillReturnResult(
	// 				sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
	// 			)
	// 		},
	// 		expectedError: kts_errors.KTS_INTERNAL_ERROR,
	// 	},
	// 	{
	// 		name: "Producer not found",
	// 		setExpectations: func(mock sqlmock.Sqlmock, producerId *uuid.UUID) {
	// 			mock.ExpectExec(query).WithArgs(utils.EqUUID(producerId)).WillReturnResult(sqlmock.NewResult(1, 0))
	// 		},
	// 		expectedError: kts_errors.KTS_NOT_FOUND,
	// 	},
	// }

	// for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		// Create a new mock database connection
	// 		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	// 		if err != nil {
	// 			t.Fatalf("Failed to create mock database connection: %v", err)
	// 		}
	// 		defer db.Close()

	// 		// Create a new instance of the genreRepository with the mock database connection
	// 		producerRepo := ProducerRepository{
	// 			DatabaseManager: &managers.DatabaseManager{
	// 				Connection: db,
	// 			},
	// 		}

	// 		tc.setExpectations(mock, producerId)

	// 		// Call the method under test
	// 		kts_err := producerRepo.DeleteProducer(producerId)

	// 		// Verify the results
	// 		assert.Equal(t, tc.expectedError, kts_err)

	// 		// Verify that all expectations were met
	// 		if err = mock.ExpectationsWereMet(); err != nil {
	// 			t.Errorf("There were unfulfilled expectations: %s", err)
	// 		}
	// 	})
	// }
	t.Skip("Not implemented")
}
