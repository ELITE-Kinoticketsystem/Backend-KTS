package models

import "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"

type GetEventSeatsDTO struct {
	EventSeat    model.EventSeats
	Seat         model.Seats
	SeatCategory model.SeatCategories
}
