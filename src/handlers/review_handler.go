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

// @Summary Create review
// @Description Create review
// @Tags Reviews
// @Accept  json
// @Produce  json
// @Param review body models.CreateReviewRequest true "Review data"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.KTSErrorMessage
// @Router /reviews [post]
func CreateReviewHandler(reviewCtrl controllers.ReviewControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviewData models.CreateReviewRequest
		err := c.ShouldBindJSON(&reviewData)
		if err != nil || utils.ContainsEmptyString(reviewData.Comment, reviewData.MovieID) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		review, username, kts_err := reviewCtrl.CreateReview(reviewData, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"review": review, "username": username})
	}
}

// @Summary Delete review
// @Description Delete review
// @Tags Reviews
// @Accept  json
// @Produce  json
// @Param id path string true "Review ID"
// @Success 200
// @Failure 400 {object} models.KTSErrorMessage
// @Router /reviews/{id} [delete]
func DeleteReviewHandler(reviewCtrl controllers.ReviewControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		kts_err := reviewCtrl.DeleteReview(&id, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
