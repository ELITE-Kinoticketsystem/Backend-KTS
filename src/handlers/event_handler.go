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

func CreateEventHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var eventData models.EventDTO
		if err := c.ShouldBindJSON(&eventData); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		event, err := eventController.CreateEvent(&eventData)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusCreated, event)
	}
}

func DeleteEventHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId := uuid.MustParse(c.Param("eventId"))
		err := eventController.DeleteEvent(&eventId)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
