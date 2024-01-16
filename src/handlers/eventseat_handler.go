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

// @Summary Get event seats
// @Description Get event seats
// @Tags EventSeats
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} models.GetEventSeatsResponse
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/{eventId}/seats [get]
func GetEventSeatsHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventSeatId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		seats, currentUserSeats, blockedUntil, kts_err := eventSeatController.GetEventSeats(&eventSeatId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, models.GetEventSeatsResponse{
			Seats:           seats,
			CurrentUserSeat: currentUserSeats,
			BlockedUntil:    blockedUntil})
	}
}

// @Summary Block event seat
// @Description Block event seat
// @Tags EventSeats
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param seatId path string true "Seat ID"
// @Success 200 {object} models.PatchEventSeatResponse
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/{eventId}/seats/{seatId}/block [patch]
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

		c.JSON(http.StatusOK, models.PatchEventSeatResponse{
			BlockedUntil: blockedUntil,
		})
	}
}

// @Summary Unblock event seat
// @Description Unblock event seat
// @Tags EventSeats
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param seatId path string true "Seat ID"
// @Success 200 {object} models.PatchEventSeatResponse
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/{eventId}/seats/{seatId}/unblock [patch]
func UnblockEventSeatHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
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

		blockedUntil, kts_err := eventSeatController.UnblockEventSeat(&eventId, &eventSeatId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, models.PatchEventSeatResponse{
			BlockedUntil: blockedUntil,
		})
	}
}

func UnblockAllEventSeatsHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		kts_err := eventSeatController.UnblockAllEventSeats(&eventId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, nil)

	}
}

// @Summary Get selected seats
// @Description Get selected seats
// @Tags EventSeats
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} models.GetSelectedSeatsResponse
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /events/{eventId}/user-seats [get]
func GetSelectedSeatsHandler(eventSeatController controllers.EventSeatControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		selectedSeats, kts_err := eventSeatController.GetSelectedSeats(&eventId, userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, models.GetSelectedSeatsResponse{
			Seats: selectedSeats,
		})
	}
}
