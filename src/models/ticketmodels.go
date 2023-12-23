package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

// type testStruct struct {
// 	model.Tickets

// 	EventSeats model.EventSeats
// }

// type TicketDTOValidate struct {
// 	ID        *uuid.UUID
// 	Validated bool
// 	Price     int32
// 	IsPaid    bool
// }

type TicketDTO struct {
	ID        *uuid.UUID
	Validated bool
	Price     int32

	Seats *model.Seats
	Order *model.Orders
	Event *model.Events
}