package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
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
