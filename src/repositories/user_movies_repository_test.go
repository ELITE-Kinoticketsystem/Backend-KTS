package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRemoveAllUserMovieCombinationWithMovie(t *testing.T) {

	movieId := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.user_movies WHERE user_movies.movie_id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Remove all movies corresponding to one movieId",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(movieId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while removing user_movies",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(movieId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
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
			userMovieRepo := UserMovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &movieId)

			// Call the method under test
			kts_err := userMovieRepo.RemoveAllUserMovieCombinationWithMovie(&movieId)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
