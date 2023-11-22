package schemas

import "github.com/google/uuid"

type Theatre struct {
	Id        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	AddressId *uuid.UUID `json:"addressId"`
}

type CinemaHall struct {
	Id        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	Capacity  int        `json:"capacity"`
	TheatreId *uuid.UUID `json:"theatreId"`
}

type Seat struct {
	Id             *uuid.UUID `json:"id"`
	Row            int        `json:"row"`
	Column         int        `json:"column"`
	SeatCategoryId *uuid.UUID `json:"seatCategoryId"`
	CinemaHallId   *uuid.UUID `json:"cinemaHallId"`
}

type SeatCategory struct {
	Id           *uuid.UUID `json:"id"`
	CategoryName string     `json:"categoryName"`
}