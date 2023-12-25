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
	ID        *uuid.UUID `sql:"primary_key"`
	Validated bool
	Price     int32

	Seats *model.Seats
	Order *model.Orders
	Event *model.Events
}

// type TicketDTO struct {
// 	model.Tickets

// 	Seats struct{ model.Seats }
// 	Order struct{ model.Orders }
// 	Event struct{ model.Events }
// }
