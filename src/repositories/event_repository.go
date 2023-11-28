package repositories

import (
	"errors"
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type EventRepo interface {
	CreateEvent(event *schemas.Event) error
	DeleteEvent(*uuid.UUID) error
	GetEventsForMovieId(movieId *uuid.UUID) ([]*schemas.Event, error)
	GetSpecialEvents() ([]*models.EventDTO, error)

	CreatePriceCategory(priceCategory *schemas.PriceCategory) (*schemas.PriceCategory, error)

	AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) error
	DeleteEventMovies(eventId *uuid.UUID) error

	CreateEventSeatCategory(eventSeatCategory *schemas.EventSeatCategory) (*schemas.EventSeatCategory, error)
	DeleteEventSeatCategoryByEventId(eventId *uuid.UUID) error
	CreateEventSeat(eventSeat *schemas.EventSeat) (*schemas.EventSeat, error)
	DeleteEventSeatsByEventId(eventId *uuid.UUID) error
}

type EventRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

const (
	event        = "event"
	specialEvent = "special_event"
)

func (er *EventRepository) CreateEvent(event *schemas.Event) error {
	_, err := er.DatabaseManager.ExecuteStatement("INSERT INTO events (id, title, start, end, event_type_id, cinema_hall_id) VALUES (?, ?, ?, ?, ?, ?)", event.Id, event.Title, event.Start, event.End, event.EventTypeId, event.CinemaHallId)
	if err != nil {
		log.Printf("Error while inserting event: %v", err)
		return err
	}

	return nil
}

func (er *EventRepository) DeleteEvent(id *uuid.UUID) error {
	query := "DELETE FROM events WHERE id=?"
	result, err := er.DatabaseManager.ExecuteStatement(query, id)
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

func (er *EventRepository) GetEventsForMovieId(movieId *uuid.UUID) ([]*schemas.Event, error) {
	// Add to query that only events that are in the future are returned
	query := "SELECT * FROM events WHERE id IN (SELECT event_id FROM event_movie WHERE movie_id=?) AND start > NOW() ORDER BY start"

	rows, err := er.DatabaseManager.ExecuteQuery(query, movieId)
	if err != nil {
		log.Printf("Error while getting events for movie id: %v", err)
		return nil, err
	}

	events := make([]*schemas.Event, 0)
	for rows.Next() {
		event := schemas.Event{}
		err := rows.Scan(&event.Id, &event.Title, &event.Start, &event.End, &event.EventTypeId, &event.CinemaHallId)
		if err != nil {
			log.Printf("Error while scanning events for movie id: %v", err)
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}

func (er *EventRepository) GetEventsForCinemaHallId(cinemaHallId *uuid.UUID) ([]*schemas.Event, error) {
	query := "SELECT * FROM events WHERE cinema_hall_id=?"
	rows, err := er.DatabaseManager.ExecuteQuery(query, cinemaHallId)
	if err != nil {
		log.Printf("Error while getting events for cinema hall id: %v", err)
		return nil, err
	}

	events := make([]*schemas.Event, 0)
	for rows.Next() {
		event := schemas.Event{}
		err := rows.Scan(&event.Id, &event.Title, &event.Start, &event.End, &event.EventTypeId, &event.CinemaHallId)
		if err != nil {
			log.Printf("Error while scanning events for cinema hall id: %v", err)
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}

// type EventDTO struct {
// 	Id                  *uuid.UUID             `json:"id"`
// 	Title               string                 `json:"title"`
// 	Start               time.Time              `json:"start"`
// 	End                 time.Time              `json:"end"`
// 	EventTypeId         *uuid.UUID             `json:"eventTypeID"`
// 	CinemaHallId        *uuid.UUID             `json:"cinemaHallID"`
// 	Movies              []MovieDTO             `json:"movie"`
// 	EventSeatCategories []EventSeatCategoryDTO `json:"eventSeatCategories"`
// }

// type EventSeatCategoryDTO struct {
// 	Name  string `json:"name"`
// 	Price int    `json:"price"`
// }

// type MovieDTO struct {
// 	Id          *uuid.UUID `json:"id"`
// 	Title       string     `json:"title"`
// 	Description string     `json:"description"`
// 	ReleaseDate time.Time  `json:"releaseDate"`
// 	TimeInMin   int        `json:"timeInMin"`
// 	Fsk         int        `json:"fsk"`
// 	GenreNames  []string   `json:"genreName"`
// }

func (er *EventRepository) GetSpecialEvents() ([]*models.EventDTO, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) CreatePriceCategory(priceCategory *schemas.PriceCategory) (*schemas.PriceCategory, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) DeleteEventMovies(eventId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) CreateEventSeatCategory(eventSeatCategory *schemas.EventSeatCategory) (*schemas.EventSeatCategory, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) CreateEventSeat(eventSeat *schemas.EventSeat) (*schemas.EventSeat, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) DeleteEventMovie(eventId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) DeleteEventSeatCategoryByEventId(eventId *uuid.UUID) error {
	return errors.New("not implemented")
}

func (er *EventRepository) DeleteEventSeatsByEventId(eventId *uuid.UUID) error {
	return errors.New("not implemented")
}
