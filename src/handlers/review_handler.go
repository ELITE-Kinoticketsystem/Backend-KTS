package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateReviewHandler(reviewCtrl controllers.ReviewControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviewData models.CreateReviewRequest
		err := c.ShouldBindJSON(&reviewData)
		if err != nil || utils.ContainsEmptyString(reviewData.Comment, reviewData.UserID, reviewData.MovieID) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		reviewId, kts_err := reviewCtrl.CreateReview(reviewData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, reviewId)
	}
}

func DeleteReviewHandler(reviewCtrl controllers.ReviewControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := reviewCtrl.DeleteReview(&id)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.Status(http.StatusOK)
	}
}
