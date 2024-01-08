package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestAddMovieActor(t *testing.T) {

	uuid1 := myid.New()
	uuid2 := myid.New()

	query := "INSERT INTO `KinoTicketSystem`.movie_actors (movie_id, actor_id) VALUES (?, ?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Create movieActor",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while creating movieActor",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "movieActor not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(sqlmock.NewResult(1, 0))
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

			// Create a new instance of the movieActorRepository with the mock database connection
			movieActorRepo := MovieActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieActorRepo.AddMovieActor(&uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// RemovemovieActor

func TestRemoveMovieActor(t *testing.T) {

	uuid1 := myid.New()
	uuid2 := myid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_genres WHERE (movie_actors.movie_id = ?) AND (movie_actors.actor_id = ?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Add movie_actor",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while adding movie_actor",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "MovieActor not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID, actorId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(uuid1), utils.EqUUID(uuid2)).WillReturnResult(sqlmock.NewResult(1, 0))
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
			movieActorRepo := MovieActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieActorRepo.RemoveMovieActor(&uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestRemoveAllActorCombinationWithMovie(t *testing.T) {

	movieId := myid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_actors WHERE movie_actors.movie_id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *myid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Remove all actors corresponding to one movie",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while removing movie_actor",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnError(sqlmock.ErrCancelled)
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
			movieActorRepo := MovieActorRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &movieId)

			// Call the method under test
			kts_err := movieActorRepo.RemoveAllActorCombinationWithMovie(&movieId)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
