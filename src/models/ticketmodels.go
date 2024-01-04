package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type TicketDTO struct {
	ID        *uuid.UUID `sql:"primary_key" alias:"ticket.id"`
	Validated bool       `alias:"ticket.validated"`
	Price     int32      `alias:"ticket.price"`

	Seats *model.Seats
	Order *model.Orders
	Event *model.Events
}

type PatchValidateTicketResponse struct {
	Message string `json:"message"`
}
