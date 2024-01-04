package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ReviewControllerI interface {
	CreateReview(reviewData models.CreateReviewRequest, userId *uuid.UUID) (*model.Reviews, string, *models.KTSError)
	DeleteReview(id *uuid.UUID, userId *uuid.UUID) *models.KTSError
}

type ReviewController struct {
	ReviewRepo repositories.ReviewRepositoryI
	UserRepo   repositories.UserRepositoryI
}

func (rc ReviewController) CreateReview(reviewData models.CreateReviewRequest, userId *uuid.UUID) (*model.Reviews, string, *models.KTSError) {
	user, kts_err := rc.UserRepo.GetUserById(userId)
	if kts_err != nil {
		return nil, "", kts_err
	}

	id := uuid.New()
	movieId, err := uuid.Parse(reviewData.MovieID)
	if err != nil {
		return nil, "", kts_errors.KTS_BAD_REQUEST
	}

	review := model.Reviews{
		ID:        &id,
		Rating:    reviewData.Rating,
		Comment:   reviewData.Comment,
		Datetime:  reviewData.Datetime,
		IsSpoiler: &reviewData.IsSpoiler,
		MovieID:   &movieId,
		UserID:    userId,
	}

	kts_error := rc.ReviewRepo.CreateReview(review)
	if kts_error != nil {
		return nil, "", kts_error
	}

	return &review, *user.Username, nil
}

func (rc ReviewController) DeleteReview(id *uuid.UUID, userId *uuid.UUID) *models.KTSError {
	review, err := rc.ReviewRepo.GetReviewById(id)
	if err != nil {
		return err
	}
	if *review.UserID != *userId {
		return kts_errors.KTS_FORBIDDEN
	}
	kts_error := rc.ReviewRepo.DeleteReview(id)
	if kts_error != nil {
		return kts_error
	}

	return nil
}
