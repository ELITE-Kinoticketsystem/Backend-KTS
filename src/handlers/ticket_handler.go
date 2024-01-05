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

// @Summary Get Ticket By Id
// @Description Get Ticket By Id
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.TicketDTO
// @Failure 500 {object} models.KTSErrorMessage
// @Router /tickets/{id} [get]
func GetTicketByIdHandler(ticketController controllers.TicketControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketId, err := uuid.Parse(c.Param("ticketId"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		ticket, kts_err := ticketController.GetTicketById(&ticketId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, ticket)
	}
}

// @Summary Validate Ticket
// @Description Validate Ticket
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.PatchValidateTicketResponse
// @Failure 409 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /tickets/{id} [patch]
func ValidateTicketHandler(ticketController controllers.TicketControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketId, err := uuid.Parse(c.Param("ticketId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := ticketController.ValidateTicket(&ticketId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, models.PatchValidateTicketResponse{Message: "Ticket successfully validated"})
	}
}
