package controllers

import (
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestCreateReview(t *testing.T) {
	user := samples.GetSampleUser()
	review := samples.GetSampleReview()
	userId := user.ID
	testCases := []struct {
		name             string
		reviewData       models.CreateReviewRequest
		setExpectations  func(userRepo mocks.MockUserRepositoryI, reviewRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest)
		expectedReview   *model.Reviews
		expectedUsername string
		expectedError    *models.KTSError
	}{
		{
			name:       "User Internal error",
			reviewData: samples.GetSampleReviewRequest(),
			setExpectations: func(userRepo mocks.MockUserRepositoryI, reviewRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				userRepo.EXPECT().GetUserById(userId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedReview:   nil,
			expectedUsername: "",
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:       "Create Internal error",
			reviewData: samples.GetSampleReviewRequest(),
			setExpectations: func(userRepo mocks.MockUserRepositoryI, reviewRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				movieId := uuid.MustParse(reviewData.MovieID)
				review := model.Reviews{
					Rating:    reviewData.Rating,
					Comment:   reviewData.Comment,
					Datetime:  reviewData.Datetime,
					IsSpoiler: &reviewData.IsSpoiler,
					MovieID:   &movieId,
					/* UserID */
				}
				userRepo.EXPECT().GetUserById(userId).Return(&user, nil)
				reviewRepo.EXPECT().CreateReview(utils.EqExceptUUIDs(review)).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedReview:   nil,
			expectedUsername: "",
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Invalid movie id",
			reviewData: models.CreateReviewRequest{
				Rating:    5,
				Comment:   "Comment",
				Datetime:  time.Now(),
				IsSpoiler: false,
				MovieID:   "invalid id",
				/* UserID */
			},
			setExpectations: func(userRepo mocks.MockUserRepositoryI, reviewRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				userRepo.EXPECT().GetUserById(userId).Return(&user, nil)
			},
			expectedReview:   nil,
			expectedUsername: "",
			expectedError:    kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:       "Success",
			reviewData: samples.GetSampleReviewRequest(),
			setExpectations: func(userRepo mocks.MockUserRepositoryI, reviewRepo mocks.MockReviewRepositoryI, reviewData models.CreateReviewRequest) {
				movieId := uuid.MustParse(reviewData.MovieID)
				review := model.Reviews{
					Rating:    reviewData.Rating,
					Comment:   reviewData.Comment,
					Datetime:  reviewData.Datetime,
					IsSpoiler: &reviewData.IsSpoiler,
					MovieID:   &movieId,
					/* UserID */
				}
				userRepo.EXPECT().GetUserById(userId).Return(&user, nil)
				reviewRepo.EXPECT().CreateReview(utils.EqExceptUUIDs(review)).Return(nil)
			},
			expectedReview:   &review,
			expectedUsername: *user.Username,
			expectedError:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock review repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			reviewRepoMock := mocks.NewMockReviewRepositoryI(mockCtrl)
			userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
			reviewController := ReviewController{
				ReviewRepo: reviewRepoMock,
				UserRepo:   userRepoMock,
			}

			// define expectations
			tc.setExpectations(*userRepoMock, *reviewRepoMock, tc.reviewData)

			// WHEN
			// call CreateReview with review data
			review, username, err := reviewController.CreateReview(tc.reviewData, userId)

			// THEN
			// check expected review, username and error
			if tc.expectedReview != nil {
				assert.NotNil(t, review, "wrong review")
				assert.True(t, utils.EqualsExceptId(*review, *tc.expectedReview), "wrong review")
			} else {
				assert.Nil(t, review, "wrong review")
			}
			assert.Equal(t, username, tc.expectedUsername, "wrong username")
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}

func TestDeleteReview(t *testing.T) {
	userId := uuid.New()
	testCases := []struct {
		name            string
		id              uuid.UUID
		setExpectations func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			id:   uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().GetReviewById(&id).Return(&model.Reviews{
					UserID: &userId,
				}, nil)
				mockRepo.EXPECT().DeleteReview(&id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Forbidden",
			id:   uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().GetReviewById(&id).Return(&model.Reviews{
					UserID: utils.NewUUID(),
				}, nil)
			},
			expectedError: kts_errors.KTS_FORBIDDEN,
		},
		{
			name: "Review internal error",
			id:   uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().GetReviewById(&id).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Delete internal error",
			id:   uuid.New(),
			setExpectations: func(mockRepo mocks.MockReviewRepositoryI, id uuid.UUID) {
				mockRepo.EXPECT().GetReviewById(&id).Return(&model.Reviews{
					UserID: &userId,
				}, nil)
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
			mockRepoMock := mocks.NewMockReviewRepositoryI(mockCtrl)
			reviewController := ReviewController{
				ReviewRepo: mockRepoMock,
			}

			// define expectations
			tc.setExpectations(*mockRepoMock, tc.id)

			// WHEN
			// call DeleteReview with review data
			err := reviewController.DeleteReview(&tc.id, &userId)

			// THEN
			// check expected error and id
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}
