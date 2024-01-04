package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		c.JSON(http.StatusOK, gin.H{"message": "Ticket successfully validated"})
	}
}
