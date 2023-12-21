package controllers

import (
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestCreateReview(t *testing.T) {
	testCases := []struct {
		name            string
		reviewData      models.CreateReviewRequest
		setExpectations func(mockRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest)
		expectedError   *models.KTSError
	}{
		{
			name:       "Internal error",
			reviewData: getSampleReviewRequest(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				movieId := uuid.MustParse(reviewData.MovieID)
				userId := uuid.MustParse(reviewData.UserID)
				review := model.Reviews{
					Rating:    reviewData.Rating,
					Comment:   reviewData.Comment,
					Datetime:  reviewData.Datetime,
					IsSpoiler: &reviewData.IsSpoiler,
					MovieID:   &movieId,
					UserID:    &userId,
				}
				mockRepo.EXPECT().CreateReview(utils.EqExceptId(review)).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Invalid movie id",
			reviewData: models.CreateReviewRequest{
				Rating:    5,
				Comment:   "Comment",
				Datetime:  time.Now(),
				IsSpoiler: false,
				MovieID:   "invalid id",
				UserID:    "fa40ef37-dba9-4b0d-9fc4-0edc7588200e",
			},
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
		{
			name: "Invalid user id",
			reviewData: models.CreateReviewRequest{
				Rating:    5,
				Comment:   "Comment",
				Datetime:  time.Now(),
				IsSpoiler: false,
				MovieID:   "7236556f-5e78-4e94-8910-3072c2f5cd5b",
				UserID:    "invalid id",
			},
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:       "Success",
			reviewData: getSampleReviewRequest(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				movieId := uuid.MustParse(reviewData.MovieID)
				userId := uuid.MustParse(reviewData.UserID)
				review := model.Reviews{
					Rating:    reviewData.Rating,
					Comment:   reviewData.Comment,
					Datetime:  reviewData.Datetime,
					IsSpoiler: &reviewData.IsSpoiler,
					MovieID:   &movieId,
					UserID:    &userId,
				}
				mockRepo.EXPECT().CreateReview(utils.EqExceptId(review)).Return(nil)
			},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock review repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			reviewRepoMock := mocks.NewMockReviewRepositoryI(mockCtrl)
			reviewController := ReviewController{
				ReviewRepo: reviewRepoMock,
			}

			// define expectations
			tc.setExpectations(*reviewRepoMock, tc.reviewData)

			// WHEN
			// call CreateReview with review data
			_, err := reviewController.CreateReview(tc.reviewData)

			// THEN
			// check expected error and id
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}

func TestDeleteReview(t *testing.T) {
	testCases := []struct {
		name            string
		id              uuid.UUID
		setExpectations func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			id: uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().DeleteReview(&id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			id: uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().DeleteReview(&id).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock review repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			reviewRepoMock := mocks.NewMockReviewRepositoryI(mockCtrl)
			reviewController := ReviewController{
				ReviewRepo: reviewRepoMock,
			}

			// define expectations
			tc.setExpectations(*reviewRepoMock, tc.id)

			// WHEN
			// call DeleteReview with review data
			err := reviewController.DeleteReview(&tc.id)

			// THEN
			// check expected error and id
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}

func getSampleReviewRequest() models.CreateReviewRequest {
	return models.CreateReviewRequest{
		Rating:    5,
		Comment:   "Comment",
		Datetime:  time.Now(),
		IsSpoiler: false,
		MovieID:   "7236556f-5e78-4e94-8910-3072c2f5cd5b",
		UserID:    "fa40ef37-dba9-4b0d-9fc4-0edc7588200e",
	}
}
