package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type CreateOrderDTO struct {
	EventSeatPriceCategories []struct {
		EventSeatId     *uuid.UUID `binding:"required"`
		PriceCategoryId *uuid.UUID `binding:"required"`
	}

	PaymentMethodID *uuid.UUID `binding:"required"`
}

type GetOrderDTO struct {
	Order      model.Orders
	Event      model.Events
	CinemaHall model.CinemaHalls
	Theatre    model.Theatres
	Movies     []model.Movies
	Tickets    []struct {
		Ticket        model.Tickets
		PriceCategory model.PriceCategories
		Seat          model.Seats
		SeatCategory  model.SeatCategories
	}
}

type OrderConfirmation struct {
	OrderId        *uuid.UUID
	EventTitle     string
	CinemaHallName string
	EventStart     string
	TotalPrice     int32
	TheatreName    string
	Tickets        []struct {
		PriceCategory string
		SeatCategory  string
		Row           int32
		Column        int32
	}
}
