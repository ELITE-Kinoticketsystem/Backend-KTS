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

// AddMovieGenre
func TestAddMovieGenre(t *testing.T) {

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	query := "INSERT INTO `KinoTicketSystem`.movie_genres (movie_id, genre_id) VALUES (?, ?);\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Create movieGenre",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while creating movieGenre",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "movieGenre not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
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

			// Create a new instance of the MovieGenreRepository with the mock database connection
			movieGenreRepo := MovieGenreRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}
			mock.ExpectBegin()
			tx, _ := db.Begin()

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieGenreRepo.AddMovieGenre(tx, &uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// RemoveMovieGenre

func TestRemoveMovieGenre(t *testing.T) {

	uuid1 := uuid.New()
	uuid2 := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_genres\nWHERE (movie_genres.movie_id = ?) AND (movie_genres.genre_id = ?);\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Add movie_genre",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while adding movie_genre",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&uuid1), utils.EqUUID(&uuid2)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID, genreId *uuid.UUID) {
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
			movieGenreRepo := MovieGenreRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &uuid1, &uuid2)

			// Call the method under test
			kts_err := movieGenreRepo.RemoveMovieGenre(&uuid1, &uuid2)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestRemoveAllGenreCombinationWithMovie(t *testing.T) {

	movieId := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_genres WHERE movie_genres.movie_id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Remove all genre corresponding to one movie",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(movieId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while removing movie_genres",
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
			movieGenreRepo := MovieGenreRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &movieId)

			// Call the method under test
			kts_err := movieGenreRepo.RemoveAllGenreCombinationWithMovie(&movieId)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestRemoveAllMovieCombinationWithGenre(t *testing.T) {

	genreId := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.movie_genres WHERE movie_genres.genre_id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genreId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Remove all movies corresponding to one genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(genreId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while removing movie_genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreId *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(genreId)).WillReturnError(sqlmock.ErrCancelled)
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
			movieGenreRepo := MovieGenreRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &genreId)

			// Call the method under test
			kts_err := movieGenreRepo.RemoveAllMovieCombinationWithGenre(&genreId)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
