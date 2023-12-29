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
		var eventData models.CreateEvtDTO
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

func GetEventsForMovieHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := uuid.MustParse(c.Param("id"))
		events, err := eventController.GetEventsForMovie(&movieId)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, events)
	}
}

func GetSpecialEventsHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		events, err := eventController.GetSpecialEvents()
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, events)
	}
}

func GetEventByIdHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("eventId"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		event, kts_err := eventController.GetEventById(&id)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, event)
	}
}
