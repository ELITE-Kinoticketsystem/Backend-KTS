package main

type CinemaHall struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Seats    []Seat `json:"seats"`
}


type Seat struct {
	ID           int          `json:"id"`
	Row          int          `json:"row"`
	Column       int          `json:"column"`
	SeatCategory SeatCategory `json:"seatCategory"`
}

type SeatCategory string

const (
	PremiumSeat  SeatCategory = "PremiumSeat"
	DisabledSeat SeatCategory = "DisabledSeat"
	CoupleSeat   SeatCategory = "CoupleSeat"
	NormalSeat   SeatCategory = "NormalSeat"
)