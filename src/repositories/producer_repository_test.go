package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetProducers(t *testing.T) {
	sampleProducers := utils.GetSampleProducers()

	query := "SELECT producers.id AS \"producers.id\", producers.name AS \"producers.name\", producers.birthdate AS \"producers.birthdate\", producers.description AS \"producers.description\", producers.pic_url AS \"producers.pic_url\" FROM `KinoTicketSystem`.producers;"

	testCases := []struct {
		name             string
		setExpectations  func(mock sqlmock.Sqlmock)
		expectedProducer *[]model.Producers
		expectedError    *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url"}),
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
						[]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url"},
					).AddRow(
						(*sampleProducers)[0].ID, (*sampleProducers)[0].Name, (*sampleProducers)[0].Birthdate, (*sampleProducers)[0].Description, (*sampleProducers)[0].PicURL,
					).AddRow(
						(*sampleProducers)[1].ID, (*sampleProducers)[1].Name, (*sampleProducers)[1].Birthdate, (*sampleProducers)[1].Description, (*sampleProducers)[1].PicURL,
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
	sampleProducers := utils.GetSampleProducer()

	query := "SELECT producers.id AS \"producers.id\", producers.name AS \"producers.name\", producers.birthdate AS \"producers.birthdate\", producers.description AS \"producers.description\", producers.pic_url AS \"producers.pic_url\" FROM `KinoTicketSystem`.producers WHERE producers.id = ?;"

	producer_id := sampleProducers.ID

	testCases := []struct {
		name             string
		setExpectations  func(mock sqlmock.Sqlmock)
		expectedProducer *model.Producers
		expectedError    *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(producer_id).WillReturnRows(
					sqlmock.NewRows([]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url"}),
				)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single producer",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(producer_id).WillReturnRows(
					sqlmock.NewRows(
						[]string{"producers.id", "producers.name", "producers.birthdate", "producers.description", "producers.pic_url"},
					).AddRow(
						sampleProducers.ID, sampleProducers.Name, sampleProducers.Birthdate, sampleProducers.Description, sampleProducers.PicURL,
					),
				)
			},
			expectedProducer: &sampleProducers,
			expectedError:    nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sampleProducers.ID).WillReturnError(sqlmock.ErrCancelled)
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
