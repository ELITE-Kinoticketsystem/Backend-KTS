package repositories

import (
	"errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

type EventRepo interface {
	CreateEvent(event *model.Events) (*uuid.UUID, error)
	DeleteEvent(*uuid.UUID) error
	GetEventsForMovieId(movieId *uuid.UUID) ([]*model.Events, error)
	GetSpecialEvents() ([]*models.EventDTO, error)

	CreatePriceCategory(priceCategory *model.PriceCategories) error

	AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) error
	DeleteEventMovies(eventId *uuid.UUID) error

	CreateEventSeatCategory(eventSeatCategory *model.EventSeatCategories) error
	DeleteEventSeatCategoryByEventId(eventId *uuid.UUID) error
	CreateEventSeat(eventSeat *model.EventSeats) error
	DeleteEventSeatsByEventId(eventId *uuid.UUID) error
}

type EventRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (er *EventRepository) CreateEvent(event *model.Events) error {

}

func (er *EventRepository) DeleteEvent(id *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) GetEventsForMovieId(movieId *uuid.UUID) ([]*model.Events, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) GetEventsForCinemaHallId(cinemaHallId *uuid.UUID) ([]*model.Events, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) GetSpecialEvents() ([]*models.EventDTO, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}

func (er *EventRepository) CreatePriceCategory(priceCategory *model.PriceCategories) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) DeleteEventMovies(eventId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) CreateEventSeatCategory(eventSeatCategory *model.EventSeatCategories) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (er *EventRepository) CreateEventSeat(eventSeat *model.EventSeats) error {
	// TODO: implement
	return errors.New("not implemented")
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
