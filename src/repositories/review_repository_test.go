package repositories

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateReview(t *testing.T) {
	testCases := []struct {
		name            string
		review          model.Reviews
		setExpectations func(mock sqlmock.Sqlmock, review model.Reviews)
		expectedError   *models.KTSError
	}{
		{
			name:   "Success",
			review: samples.GetSampleReview(),
			setExpectations: func(mock sqlmock.Sqlmock, review model.Reviews) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.reviews (id, rating, comment, datetime, is_spoiler, user_id, movie_id)\n"+
						"VALUES (?, ?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(review.ID),
					review.Rating,
					review.Comment,
					review.Datetime,
					review.IsSpoiler,
					utils.EqUUID(review.UserID),
					utils.EqUUID(review.MovieID),
				).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			expectedError: nil,
		},
		{
			name:   "Insert internal error",
			review: samples.GetSampleReview(),
			setExpectations: func(mock sqlmock.Sqlmock, review model.Reviews) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.reviews (id, rating, comment, datetime, is_spoiler, user_id, movie_id)\n"+
						"VALUES (?, ?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(review.ID),
					review.Rating,
					review.Comment,
					review.Datetime,
					review.IsSpoiler,
					utils.EqUUID(review.UserID),
					utils.EqUUID(review.MovieID),
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:   "Affected rows internal error",
			review: samples.GetSampleReview(),
			setExpectations: func(mock sqlmock.Sqlmock, review model.Reviews) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.reviews (id, rating, comment, datetime, is_spoiler, user_id, movie_id)\n"+
						"VALUES (?, ?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(review.ID),
					review.Rating,
					review.Comment,
					review.Datetime,
					review.IsSpoiler,
					utils.EqUUID(review.UserID),
					utils.EqUUID(review.MovieID),
				).WillReturnResult(
					sqlmock.NewErrorResult(sqlmock.ErrCancelled),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:   "No affected rows",
			review: samples.GetSampleReview(),
			setExpectations: func(mock sqlmock.Sqlmock, review model.Reviews) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.reviews (id, rating, comment, datetime, is_spoiler, user_id, movie_id)\n"+
						"VALUES (?, ?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(review.ID),
					review.Rating,
					review.Comment,
					review.Datetime,
					review.IsSpoiler,
					utils.EqUUID(review.UserID),
					utils.EqUUID(review.MovieID),
				).WillReturnResult(
					sqlmock.NewResult(0, 0),
				)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			reviewRepo := ReviewRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.review)

			// WHEN
			// call CreateReview with review data
			kts_err := reviewRepo.CreateReview(tc.review)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetReview(t *testing.T) {
	review := samples.GetSampleReview()
	query := "SELECT reviews.id AS \"reviews.id\",\n" +
		"reviews.rating AS \"reviews.rating\",\n" +
		"reviews.comment AS \"reviews.comment\",\n" +
		"reviews.datetime AS \"reviews.datetime\",\n" +
		"reviews.is_spoiler AS \"reviews.is_spoiler\",\n" +
		"reviews.user_id AS \"reviews.user_id\",\n" +
		"reviews.movie_id AS \"reviews.movie_id\"\n" +
		"FROM `KinoTicketSystem`.reviews\n" +
		"WHERE reviews.id = ?;"
	testCases := []struct {
		name            string
		id              uuid.UUID
		setExpectations func(mock sqlmock.Sqlmock, id uuid.UUID)
		expectedReview  *model.Reviews
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			id:   uuid.New(),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectQuery(
					query,
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnRows(
					sqlmock.NewRows([]string{"reviews.id", "reviews.rating", "reviews.comment", "reviews.datetime", "reviews.is_spoiler", "reviews.user_id", "reviews.movie_id"}).
						AddRow(
							review.ID,
							review.Rating,
							review.Comment,
							review.Datetime,
							review.IsSpoiler,
							review.UserID,
							review.MovieID,
						),
				)
			},
			expectedReview: &review,
			expectedError:  nil,
		},
		{
			name: "Internal error",
			id:   uuid.New(),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectQuery(
					query,
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedReview: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Not found",
			id:   uuid.New(),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectQuery(
					query,
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnError(
					sql.ErrNoRows,
				)
			},
			expectedReview: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			reviewRepo := ReviewRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.id)

			// WHEN
			// call GetReviewById with id
			review, kts_err := reviewRepo.GetReviewById(&tc.id)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedReview, review)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestDeleteReview(t *testing.T) {
	testCases := []struct {
		name            string
		id              uuid.UUID
		setExpectations func(mock sqlmock.Sqlmock, id uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.id = ?;",
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			expectedError: nil,
		},
		{
			name: "Delete internal error",
			id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.id = ?;",
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Rows affected internal error",
			id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.id = ?;",
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnResult(
					sqlmock.NewErrorResult(sqlmock.ErrCancelled),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "No rows affected",
			id:   uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, id uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.id = ?;",
				).WithArgs(
					utils.EqUUID(&id),
				).WillReturnResult(
					sqlmock.NewResult(0, 0),
				)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			reviewRepo := ReviewRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.id)

			// WHEN
			// call DeleteReview with id
			kts_err := reviewRepo.DeleteReview(&tc.id)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetRatingForMovie(t *testing.T) {
	movieRating := samples.GetSampleNewRating()

	query := "SELECT SUM(reviews.rating) FROM `KinoTicketSystem`.reviews WHERE reviews.movie_id = ?;"

	testCases := []struct {
		name            string
		movieId         *uuid.UUID
		setExpectations func(mock sqlmock.Sqlmock, movieId *uuid.UUID)
		expectedRating  *models.NewRating
		expectedError   *models.KTSError
	}{
		{
			name:    "Success",
			movieId: utils.NewUUID(),
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(movieId)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"SUM(reviews.rating)"},
					).AddRow(
						movieRating.Rating,
					),
				)
			},
			expectedRating: &movieRating,
			expectedError:  nil,
		},
		{
			name:    "Internal error",
			movieId: utils.NewUUID(),
			setExpectations: func(mock sqlmock.Sqlmock, movieId *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(
					utils.EqUUID(movieId),
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedRating: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			reviewRepo := ReviewRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.movieId)

			// WHEN
			// call DeleteReview with id
			rating, kts_err := reviewRepo.GetRatingForMovie(tc.movieId)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedRating, rating)
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestDeleteReviewForMovie(t *testing.T) {
	testCases := []struct {
		name            string
		movieId         uuid.UUID
		setExpectations func(mock sqlmock.Sqlmock, movieId uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name:    "Success",
			movieId: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, movieId uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.movie_id = ?;",
				).WithArgs(
					utils.EqUUID(&movieId),
				).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			expectedError: nil,
		},
		{
			name:    "Delete internal error",
			movieId: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			setExpectations: func(mock sqlmock.Sqlmock, movieId uuid.UUID) {
				mock.ExpectExec(
					"DELETE FROM `KinoTicketSystem`.reviews\n" +
						"WHERE reviews.movie_id = ?;",
				).WithArgs(
					utils.EqUUID(&movieId),
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			reviewRepo := ReviewRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.movieId)

			// WHEN
			// call DeleteReview with id
			kts_err := reviewRepo.DeleteReviewForMovie(&tc.movieId)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
