package controllers

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type EventControllerI interface {
	CreateEvent(event *schemas.Event) error
}

type EventController struct {
	DatabaseManager managers.DatabaseManagerI
	EventRepo       repositories.EventRepo
	MovieRepo       repositories.MovieRepoI
	TheaterRepo     repositories.TheaterRepoI
}

func (ec *EventController) CreateEvent(eventRequest *models.EventDTO) error {
	eventId := uuid.New()
	event := &schemas.Event{
		Id:           &eventId,
		Title:        eventRequest.Title,
		Start:        eventRequest.Start,
		End:          eventRequest.End,
		EventTypeId:  eventRequest.EventTypeId,
		CinemaHallId: eventRequest.CinemaHallId,
	}
	createdEvent, err := ec.EventRepo.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil

	createdEventId := createdEvent.Id

	err = ec.createMovies(eventRequest.Movies, createdEventId)
	if err != nil {
		return err
	}

	err = ec.createEventSeatCategories(eventRequest.EventSeatCategories, createdEventId)
	if err != nil {
		return err
	}

	err = ec.createEventSeats(eventRequest.CinemaHallId, createdEventId)
	if err != nil {
		return err
	}

	return nil
}

func (ec *EventController) createMovies(movies []models.MovieDTO, eventId *uuid.UUID) error {
	for _, movie := range movies {
		if movie.Id == nil {
			movieId := uuid.New()
			genreNames := movie.GenreNames

			movie := &schemas.Movie{
				Id:          &movieId,
				Title:       movie.Title,
				Description: movie.Description,
				ReleaseDate: movie.ReleaseDate,
				TimeInMin:   movie.TimeInMin,
				Fsk:         movie.Fsk,
			}

			createdMovie, err := ec.MovieRepo.CreateMovie(movie)
			if err != nil {
				return err
			}

			createdMovieId := createdMovie.Id

			err = ec.EventRepo.AddEventMovie(eventId, createdMovieId)
			if err != nil {
				return err
			}

			err = ec.createGenres(genreNames, createdMovieId)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
			ec.MovieRepo.CreateGenre(genre)
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
	seats, err := ec.TheaterRepo.GetSeatsForCinemaHall(cinemaHallId)
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
