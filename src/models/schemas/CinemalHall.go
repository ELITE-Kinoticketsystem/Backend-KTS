package schemas

import "github.com/google/uuid"

type CinemaHall struct {
	Id        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	Capacity  int        `json:"capacity"`
	TheatreId *uuid.UUID `json:"theatreId"`
}

type Seat struct {
	Id           *uuid.UUID   `json:"id"`
	Row          int          `json:"row"`
	Column       int          `json:"column"`
	SeatCategory SeatCategory `json:"seatCategory"`
	CinemaHallId *uuid.UUID   `json:"cinemaHallId"`
}

type SeatCategory string

const (
	PremiumSeat  SeatCategory = "PremiumSeat"
	DisabledSeat SeatCategory = "DisabledSeat"
	CoupleSeat   SeatCategory = "CoupleSeat"
	NormalSeat   SeatCategory = "NormalSeat"
)
