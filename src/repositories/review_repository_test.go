package repositories

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
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
			review: getSampleReview(),
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
			review: getSampleReview(),
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
			review: getSampleReview(),
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
			review: getSampleReview(),
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

func getSampleReview() model.Reviews {
	id := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	movieId := uuid.MustParse("db30d28d-506a-4637-9e9e-aef1546f9cdc")
	userId := uuid.MustParse("1264775d-b14a-43d6-a158-1bb5040f4b90")
	return model.Reviews{
		ID:        &id,
		Rating:    5,
		Comment:   "Comment",
		Datetime:  time.Now(),
		IsSpoiler: new(bool),
		MovieID:   &movieId,
		UserID:    &userId,
	}
}
