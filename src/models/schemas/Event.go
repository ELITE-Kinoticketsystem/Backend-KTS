package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id                  *uuid.UUID `json:"id"`
	Title               string     `json:"title"`
	Start               time.Time  `json:"start"`
	End                 time.Time  `json:"end"`
	EventTypeId         *uuid.UUID `json:"eventTypeId"`
	CinemaHallId        *uuid.UUID `json:"cinemaHallId"`
	EventSeatCategoryId *uuid.UUID `json:"eventMovieId"`
}

type EventType struct {
	Id       *uuid.UUID `json:"id"`
	TypeName string     `json:"typeName"`
}

type EventMovie struct {
	EventId *uuid.UUID `json:"eventId"`
	MovieId *uuid.UUID `json:"movieId"`
}

type Ticket struct {
	Id              *uuid.UUID `json:"id"`
	Validated       bool       `json:"validated"`
	Price           int        `json:"price"` // requires conversion
	PriceCategoryId *uuid.UUID `json:"priceCategoryId"`
	OrderId         *uuid.UUID `json:"orderId"`
	EventSeatId     *uuid.UUID `json:"eventSeatId"`
}

type PriceCategory struct {
	CategoryName string `json:"categoryName"`
	Price        int    `json:"price"`
}

type EventSeatCategory struct {
	EventId        *uuid.UUID
	SeatCategoryId *uuid.UUID
	Price          int
}

type EventSeat struct {
	Id           *uuid.UUID
	Booked       bool
	BlockedUntil time.Time
	UserId       *uuid.UUID
	SeatId       *uuid.UUID
	EventId      *uuid.UUID
}
