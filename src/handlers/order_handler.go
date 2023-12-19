package handlers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateOrderHandler(orderController controllers.OrderControllerI, isReservation bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		createOrderDTO := models.CreateOrderDTO{}

		if err := c.ShouldBindJSON(&createOrderDTO); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		order, kts_err := orderController.CreateOrder(createOrderDTO, &eventId, userId, isReservation)

		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(200, order)
	}
}
