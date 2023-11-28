package repositories

import (
	"errors"

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
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) DeleteEvent(id *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) GetEventsForMovieId(movieId *uuid.UUID) ([]*schemas.Event, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) GetEventsForCinemaHallId(cinemaHallId *uuid.UUID) ([]*schemas.Event, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
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
