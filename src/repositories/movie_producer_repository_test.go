package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddMovieProducer(t *testing.T) {

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	query := "INSERT INTO `KinoTicketSystem`.movie_producers (movie_id, producer_id) VALUES (?, ?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Create movieProducer",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while creating movieProducer",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "movieProducer not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
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

			// Create a new instance of the movieProducerRepository with the mock database connection
			movieProducerRepo := MovieProducerRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieProducerRepo.AddMovieProducer(&uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// RemovemovieProducer

func TestRemoveMovieProducer(t *testing.T) {

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_producers WHERE (movie_producers.movie_id = ?) AND (movie_producers.producer_id = ?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Add movie_producer",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while adding movie_producer",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "MovieProducer not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, producerId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
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

			// Create a new instance of the MovieRepository with the mock database connection
			movieProducerRepo := MovieProducerRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieProducerRepo.RemoveMovieProducer(&uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
