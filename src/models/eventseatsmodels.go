package models

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
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
	ID             *uuid.UUID
	RowNr          int32
	ColumnNr       int32
	Available      bool
	BlockedByOther bool
	Category       string
	Type           string
	Price          int32
}

type GetEventSeatsResponse struct {
	BlockedUntil    *time.Time                    `json:"blockedUntil"`
	CurrentUserSeat *[]GetSeatsForSeatSelectorDTO `json:"currentUserSeats"`
	Height          int32                         `json:"height"`
	Seats           *[]GetSeatsForSeatSelectorDTO `json:"seats"`
	Width           int32                         `json:"width"`
}

type PatchEventSeatResponse struct {
	BlockedUntil *time.Time `json:"blockedUntil"`
}

type GetSelectedSeatsResponse struct {
	Seats *[]GetSlectedSeatsDTO `json:"selectedSeats"`
}
