package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id           *uuid.UUID `json:"id"`
	Title        string     `json:"title"`
	EventType    EventType  `json:"eventType"`
	Start        time.Time  `json:"start"`
	End          time.Time  `json:"end"`
	Price        int        `json:"price"` // requires conversion
	CinemaHallId *uuid.UUID `json:"cinemaHallId"`
}

type EventType string

const (
	Showing      EventType = "Showing"
	SpecialEvent EventType = "SpecialEvent"
)
