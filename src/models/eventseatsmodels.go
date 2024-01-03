package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type GetEventSeatsDTO struct {
	EventSeat         model.EventSeats
	Seat              model.Seats
	SeatCategory      model.SeatCategories
	EventSeatCategory model.EventSeatCategories
}

type GetSeatsForSeatSelectorDTO struct {
	ID            *uuid.UUID
	RowNr         int32
	ColumnNr      int32
	Available     bool
	BookedByOther bool
	Category      string
	Type          string
	Price         int32
}
