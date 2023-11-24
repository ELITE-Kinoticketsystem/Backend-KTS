package models

import (
	"time"

	"github.com/google/uuid"
)

type EventDTO struct {
	Id                  *uuid.UUID             `json:"id"`
	Title               string                 `json:"title"`
	Start               time.Time              `json:"start"`
	End                 time.Time              `json:"end"`
	EventTypeId         *uuid.UUID             `json:"eventTypeID"`
	CinemaHallId        *uuid.UUID             `json:"cinemaHallID"`
	Movies              []MovieDTO             `json:"movie"`
	EventSeatCategories []EventSeatCategoryDTO `json:"eventSeatCategories"`
}

type EventSeatCategoryDTO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
