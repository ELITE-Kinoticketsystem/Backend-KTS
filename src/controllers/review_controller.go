package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ReviewControllerI interface {
	CreateReview(reviewData models.CreateReviewRequest) (*uuid.UUID, *models.KTSError)
	DeleteReview(id *uuid.UUID) *models.KTSError
}

type ReviewController struct {
	ReviewRepo repositories.ReviewRepositoryI
}

func (rc ReviewController) CreateReview(reviewData models.CreateReviewRequest) (*uuid.UUID, *models.KTSError) {
	id := uuid.New()
	movieId, err := uuid.Parse(reviewData.MovieID)
	if err != nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}
	userId, err := uuid.Parse(reviewData.UserID)
	if err != nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	review := model.Reviews{
		ID:        &id,
		Rating:    reviewData.Rating,
		Comment:   reviewData.Comment,
		Datetime:  reviewData.Datetime,
		IsSpoiler: &reviewData.IsSpoiler,
		MovieID:   &movieId,
		UserID:    &userId,
	}

	kts_error := rc.ReviewRepo.CreateReview(review)
	if kts_error != nil {
		return nil, kts_error
	}

	return &id, nil
}

func (rc ReviewController) DeleteReview(id *uuid.UUID) *models.KTSError {
	kts_error := rc.ReviewRepo.DeleteReview(id)
	if kts_error != nil {
		return kts_error
	}

	return nil
}
