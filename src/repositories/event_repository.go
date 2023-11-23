package repositories

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type EventRepositoryI interface {
	CreateEvent(event *schemas.Event) error
	UpdateEvent(event *schemas.Event) error
	DeleteEvent(event *schemas.Event) error
}

type EventRepository struct {
	DatabaseMgr managers.DatabaseManagerI
}

// THE STRUCTS REPRESENT THE TABLES IN THE DATABASE
// type Event struct {
// 	Id                  *uuid.UUID `json:"id"`
// 	Title               string     `json:"title"`
// 	Start               time.Time  `json:"start"`
// 	End                 time.Time  `json:"end"`
// 	EventTypeId         *uuid.UUID `json:"eventTypeId"`
// 	CinemaHallId        *uuid.UUID `json:"cinemaHallId"`
// 	EventSeatCategoryId *uuid.UUID `json:"eventSeatCategoryId"`
// }

// type EventType struct {
// 	Id       *uuid.UUID `json:"id"`
// 	TypeName string     `json:"typeName"`
// }

// type EventMovie struct {
// 	EventId *uuid.UUID `json:"eventId"`
// 	MovieId *uuid.UUID `json:"movieId"`
// }

// type Ticket struct {
// 	Id              *uuid.UUID `json:"id"`
// 	Validated       bool       `json:"validated"`
// 	Price           int        `json:"price"` // requires conversion
// 	PriceCategoryId *uuid.UUID `json:"priceCategoryId"`
// 	OrderId         *uuid.UUID `json:"orderId"`
// 	EventSeatId     *uuid.UUID `json:"eventSeatId"`
// }

// type PriceCategory struct {
// 	CategoryName string `json:"categoryName"`
// 	Price        int    `json:"price"`
// }

// type EventSeatCategory struct {
// 	EventId        *uuid.UUID
// 	SeatCategoryId *uuid.UUID
// 	Price          int
// }

// type SeatCategory struct {
// 	Id       *uuid.UUID
// 	Category string
// }

// type EventSeat struct {
// 	Id           *uuid.UUID
// 	Booked       bool
// 	BlockedUntil time.Time
// 	UserId       *uuid.UUID
// 	SeatId       *uuid.UUID
// 	EventId      *uuid.UUID
// }

func (er *EventRepository) CreateEvent(event *schemas.Event) error {
	_, err := er.DatabaseMgr.ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES ($1, $2, $3, $4, $5, $6)", event.Id, event.Title, event.Start, event.End, event.EventTypeId, event.CinemaHallId)
	if err != nil {
		log.Printf("Error while inserting event: %v", err)
		return err
	}

	return nil
}

func (er *EventRepository) UpdateEvent(event *schemas.Event) error {
	result, err := er.DatabaseMgr.ExecuteStatement("UPDATE events SET title=$1, start=$2, end=$3, event_type_id=$4, cinema_hall_id=$5 WHERE id=$6", event.Title, event.Start, event.End, event.EventTypeId, event.CinemaHallId, event.Id)
	if err != nil {
		log.Printf("Error while updating event: %v", err)
		return err
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("No event with id %v found", event.Id)
		return err
	}

	return nil
}

func (er *EventRepository) DeleteEvent(id *uuid.UUID) error {
	query := "DELETE FROM events WHERE id=$1"
	result, err := er.DatabaseMgr.ExecuteStatement(query, id)
	if err != nil {
		log.Printf("Error while deleting event: %v", err)
		return err
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		log.Printf("No event with id %v found", id)
		return err
	}
	return nil
}

func (er *EventRepository) GetEvent(id *uuid.UUID) (*schemas.Event, error) {
	query := "SELECT * FROM events where id=$1"
	row := er.DatabaseMgr.ExecuteQueryRow(query, id)

	event := schemas.Event{}
	err := row.Scan(&event.Id, &event.Title, &event.Start, &event.End, &event.EventTypeId, &event.CinemaHallId)
	if err != nil {
		log.Printf("Error while getting event: %v", err)
		return nil, err
	}

	return &event, nil
}
