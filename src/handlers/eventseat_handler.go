package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
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
		eventSeat, kts_err := eventSeatController.GetEventSeats(&eventSeatId)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, eventSeat)
	}
}

func BlockEventSeatHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventSeatId, err := uuid.Parse(c.Param("seatId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := uuid.Nil

		kts_err := eventSeatController.BlockEventSeat(&eventSeatId, &userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, eventSeatId)
	}
}
