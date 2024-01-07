package models

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

type GetEventSeatsDTO struct {
	EventSeat         model.EventSeats
	Seat              model.Seats
	SeatCategory      model.SeatCategories
	EventSeatCategory model.EventSeatCategories
}

type GetSlectedSeatsDTO struct {
	EventSeat         model.EventSeats
	Seat              model.Seats
	SeatCategory      model.SeatCategories
	EventSeatCategory model.EventSeatCategories
	CinemaHall        model.CinemaHalls
	Theatre           model.Theatres
}

type GetSeatsForSeatSelectorDTO struct {
	ID             *myid.UUID
	RowNr          int32
	ColumnNr       int32
	Available      bool
	BlockedByOther bool
	Category       string
	Type           string
	Price          int32
}

type GetEventSeatsResponse struct {
	BlockedUntil    *time.Time                      `json:"blockedUntil"`
	CurrentUserSeat *[]GetSeatsForSeatSelectorDTO   `json:"currentUserSeats"`
	SeatRows        *[][]GetSeatsForSeatSelectorDTO `json:"seat_rows"`
}

type PatchEventSeatResponse struct {
	BlockedUntil *time.Time `json:"blockedUntil"`
}

type GetSelectedSeatsResponse struct {
	Seats *[]GetSlectedSeatsDTO `json:"selectedSeats"`
}
