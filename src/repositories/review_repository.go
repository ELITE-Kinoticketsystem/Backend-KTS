package repositories

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type ReviewRepositoryI interface {
	CreateReview(review model.Reviews) *models.KTSError
	DeleteReview(id *uuid.UUID) *models.KTSError
}

type ReviewRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (rr *ReviewRepository) CreateReview(review model.Reviews) *models.KTSError {
	stmt := table.Reviews.INSERT(
		table.Reviews.ID,
		table.Reviews.Rating,
		table.Reviews.Comment,
		table.Reviews.Datetime,
		table.Reviews.IsSpoiler,
		table.Reviews.UserID,
		table.Reviews.MovieID,
	).VALUES(
		utils.MysqlUuid(review.ID),
		review.Rating,
		review.Comment,
		review.Datetime,
		review.IsSpoiler != nil && *review.IsSpoiler,
		utils.MysqlUuid(review.UserID),
		utils.MysqlUuid(review.MovieID),
	)

	result, err := stmt.Exec(rr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}
	return nil
}

func (rr *ReviewRepository) DeleteReview(id *uuid.UUID) *models.KTSError {
	stmt := table.Reviews.DELETE().WHERE(table.Reviews.ID.EQ(utils.MysqlUuid(id)))
	result, err := stmt.Exec(rr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
