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

// @Summary Create event
// @Description Create event
// @Tags Events
// @Accept  json
// @Produce  json
// @Param event body models.CreateEvtDTO true "Event data"
// @Success 200 {array} uuid.UUID
// @Failure 500 {object} string
// @Failure 400 {object} models.KTSErrorMessage
// @Router /events [post]
func CreateEventHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var eventData models.CreateEvtDTO
		if err := c.ShouldBindJSON(&eventData); utils.ContainsEmptyString(eventData.Events.Title) || err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		eventId, err := eventController.CreateEvent(&eventData)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusCreated, eventId)
	}
}

// @Summary Get events for movie
// @Description Get events for movie
// @Tags Events
// @Accept  json
// @Produce  json
// @Param id path string true "Movie ID"
// @Success 200 {array} model.Events
// @Failure 500 {object} models.KTSErrorMessage
// @Router /movies/{id}/events [get]
func GetEventsForMovieHandler(eventController controllers.EventControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		theatreId, err := uuid.Parse(c.Param("theatreId"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		events, kts_err := eventController.GetEventsForMovie(&movieId, &theatreId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, events)
	}
}

// @Summary Get special events
// @Description Get special events
// @Tags Events
// @Accept  json
// @Produce  json
// @Success 200 {array} models.GetSpecialEventsDTO
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/special [get]
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

// @Summary Get event by id
// @Description Get event by id
// @Tags Events
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} models.GetSpecialEventsDTO
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/{eventId} [get]
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
