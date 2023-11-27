package controllers

import (
	"log"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type EventControllerI interface {
	CreateEvent(event *models.EventDTO) (*schemas.Event, *models.KTSError)
	DeleteEvent(eventId *uuid.UUID) *models.KTSError
	GetEventsForMovie(movieId *uuid.UUID) ([]*schemas.Event, *models.KTSError)
}

type EventController struct {
	EventRepo   repositories.EventRepo
	MovieRepo   repositories.MovieRepoI
	TheatreRepo repositories.TheaterRepoI
}

func (ec *EventController) CreateEvent(eventRequest *models.EventDTO) (*schemas.Event, *models.KTSError) {
	eventId := uuid.New()
	event := &schemas.Event{
		Id:           &eventId,
		Title:        eventRequest.Title,
		Start:        eventRequest.Start,
		End:          eventRequest.End,
		EventTypeId:  eventRequest.EventTypeId,
		CinemaHallId: eventRequest.CinemaHallId,
	}
	err := ec.EventRepo.CreateEvent(event)
	if err != nil {
		log.Printf("Error creating event: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	createdEventId := event.Id

	err = ec.createMovies(eventRequest.Movies, createdEventId)
	if err != nil {
		log.Printf("Error creating movies: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	err = ec.createEventSeatCategories(eventRequest.EventSeatCategories, createdEventId)
	if err != nil {
		log.Printf("Error creating event seat categories: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	err = ec.createEventSeats(eventRequest.CinemaHallId, createdEventId)
	if err != nil {
		log.Printf("Error creating event seats: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return event, nil
}

func (ec *EventController) DeleteEvent(eventId *uuid.UUID) *models.KTSError {

	err := ec.EventRepo.DeleteEvent(eventId)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	err = ec.EventRepo.DeleteEventMovies(eventId)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	err = ec.EventRepo.DeleteEventSeatCategoryByEventId(eventId)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	err = ec.EventRepo.DeleteEventSeatsByEventId(eventId)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (ec *EventController) GetEventsForMovie(movieId *uuid.UUID) ([]*schemas.Event, *models.KTSError) {
	events, err := ec.EventRepo.GetEventsForMovieId(movieId)
	if err != nil {
		log.Printf("Error getting events for movie: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return events, nil
}

func (ec *EventController) createMovies(movies []models.MovieDTO, eventId *uuid.UUID) error {
	// TODO: still not completely though trough
	// Currently when the id field is not set, we create a new movie
	// How do we know we dont already have this movie in the database?

	for _, movie := range movies {
		var movieId uuid.UUID
		if movie.Id == nil {
			createdMovie, err := ec.createNewMovie(&movie)
			if err != nil {
				return err
			}
			movieId = *createdMovie.Id
		} else {
			movieId = *movie.Id
		}
		err := ec.EventRepo.AddEventMovie(eventId, &movieId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ec *EventController) createNewMovie(movieRequest *models.MovieDTO) (*schemas.Movie, error) {
	movieId := uuid.New()
	genreNames := movieRequest.GenreNames

	movie := &schemas.Movie{
		Id:          &movieId,
		Title:       movieRequest.Title,
		Description: movieRequest.Description,
		ReleaseDate: movieRequest.ReleaseDate,
		TimeInMin:   movieRequest.TimeInMin,
		Fsk:         movieRequest.Fsk,
	}

	err := ec.MovieRepo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}

	createdMovieId := movie.Id
	if createdMovieId == nil {
		return nil, err
	}

	err = ec.createGenres(genreNames, createdMovieId)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (ec *EventController) createGenres(genreNames []string, movieId *uuid.UUID) error {
	for _, genreName := range genreNames {
		genre, err := ec.MovieRepo.GetGenreByName(genreName)
		if err != nil {
			return err
		}
		var genreId uuid.UUID
		if genre == nil {
			genreId = uuid.New()
			genre := &schemas.Genre{
				Id:        &genreId,
				GenreName: genreName,
			}
			err := ec.MovieRepo.CreateGenre(genre)
			if err != nil {
				return err
			}
		}

		err = ec.MovieRepo.AddMovieGenre(movieId, &genreId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ec *EventController) createEventSeatCategories(eventSeatCategories []models.EventSeatCategoryDTO, eventId *uuid.UUID) error {
	for _, eventSeatCategory := range eventSeatCategories {
		seatCategoryId := uuid.New()
		eventSeatCategory := &schemas.EventSeatCategory{
			Price:          eventSeatCategory.Price,
			EventId:        eventId,
			SeatCategoryId: &seatCategoryId,
		}
		_, err := ec.EventRepo.CreateEventSeatCategory(eventSeatCategory)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ec *EventController) createEventSeats(cinemaHallId *uuid.UUID, eventId *uuid.UUID) error {
	seats, err := ec.TheatreRepo.GetSeatsForCinemaHall(cinemaHallId)
	if err != nil {
		return err
	}

	for _, seat := range seats {
		eventSeatId := uuid.New()
		eventSeat := &schemas.EventSeat{
			Id:           &eventSeatId,
			Booked:       false,
			BlockedUntil: time.Time{},
			UserId:       nil,
			SeatId:       seat.Id,
			EventId:      eventId,
		}
		_, err := ec.EventRepo.CreateEventSeat(eventSeat)
		if err != nil {
			return err
		}
	}

	return nil
}
