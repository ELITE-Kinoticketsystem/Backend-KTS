package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

func GetSampleReviewRequest() models.CreateReviewRequest {
	datetime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return models.CreateReviewRequest{
		Rating:    5,
		Comment:   "Comment",
		Datetime:  datetime,
		IsSpoiler: false,
		MovieID:   "db30d28d-506a-4637-9e9e-aef1546f9cdc",
	}
}

func GetSampleReview() model.Reviews {
	id := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	movieId := uuid.MustParse("db30d28d-506a-4637-9e9e-aef1546f9cdc")
	userId := uuid.MustParse("47cf7525-01df-45b7-a3a9-d3cb25ae939f")
	datetime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return model.Reviews{
		ID:        &id,
		Rating:    5,
		Comment:   "Comment",
		Datetime:  datetime,
		IsSpoiler: new(bool),
		MovieID:   &movieId,
		UserID:    &userId,
	}
}

func GetSampleNewRating() models.NewRating {
	
	return models.NewRating{
		Rating: 3.4,
	}
}
