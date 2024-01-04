package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

func GetSampleReviewRequest() models.CreateReviewRequest {
	return models.CreateReviewRequest{
		Rating:    5,
		Comment:   "Comment",
		Datetime:  time.Now(),
		IsSpoiler: false,
		MovieID:   "7236556f-5e78-4e94-8910-3072c2f5cd5b",
	}
}

func GetSampleReview() model.Reviews {
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
