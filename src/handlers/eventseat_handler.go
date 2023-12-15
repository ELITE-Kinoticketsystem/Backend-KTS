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

func GetEventSeatsHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventSeatId, err := uuid.Parse(c.Param("id"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		seatMap, currentUserSeats, blockedUntil, kts_err := eventSeatController.GetEventSeats(&eventSeatId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"seat_rows":        seatMap,
			"currentUserSeats": currentUserSeats,
			"blockedUntil":     blockedUntil,
		})
	}
}

func BlockEventSeatHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		eventSeatId, err := uuid.Parse(c.Param("seatId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		blockedUntil, kts_err := eventSeatController.BlockEventSeat(&eventId, &eventSeatId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"blockedUntil": blockedUntil,
		})
	}
}

