package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type CreateOrderDTO struct {
	EventSeatPriceCategories []struct {
		EventSeatId     *uuid.UUID
		PriceCategoryId *uuid.UUID
	}

	PaymentMethodID *uuid.UUID
}

type GetOrderDTO struct {
	Order   model.Orders
	Tickets []struct {
		Ticket model.Tickets
		Seat   model.Seats
	}
}
