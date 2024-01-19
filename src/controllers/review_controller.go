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
	MovieRepo  repositories.MovieRepositoryI
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

	kts_error_update_rating := updateMovieRating(rc, &movieId)
	if kts_error_update_rating != nil {
		return nil, "", kts_error_update_rating
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

	kts_error_update_rating := updateMovieRating(rc, review.MovieID)
	if kts_error_update_rating != nil {
		return kts_error_update_rating
	}

	return nil
}

func updateMovieRating(rc ReviewController, movieId *uuid.UUID) *models.KTSError {
	ratings, kts_err_get := rc.ReviewRepo.GetRatingForMovie(movieId)
	if kts_err_get != nil {
		return kts_err_get
	}

	newMovieRating := ratings.Rating / float64(ratings.TotalRatings)

	kts_err_update := rc.MovieRepo.UpdateRating(movieId, newMovieRating)
	if kts_err_update != nil {
		return kts_err_update
	}

	return nil
}
