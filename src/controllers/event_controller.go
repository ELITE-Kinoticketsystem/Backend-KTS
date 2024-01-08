package controllers

import (
	"log"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type EventControllerI interface {
	CreateEvent(event *models.CreateEvtDTO) (*myid.UUID, *models.KTSError)
	GetEventsForMovie(movieId *myid.UUID, theatreId *myid.UUID) ([]*model.Events, *models.KTSError)
	GetSpecialEvents() (*[]models.GetSpecialEventsDTO, *models.KTSError)
	GetEventById(eventId *myid.UUID) (*models.GetSpecialEventsDTO, *models.KTSError)
}

type EventController struct {
	EventRepo   repositories.EventRepo
	TheatreRepo repositories.TheaterRepoI
}

func (ec *EventController) CreateEvent(eventDto *models.CreateEvtDTO) (*myid.UUID, *models.KTSError) {
	if eventDto == nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	event := eventDto.Events

	eventId, kts_err := ec.EventRepo.CreateEvent(&event)

	if kts_err != nil {
		log.Printf("Error creating event: %v", kts_err.ErrorMessage)
		return nil, kts_err
	}

	movies := eventDto.Movies

	if len(movies) == 0 {
		log.Printf("No movies provided for event: %v", eventId)
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	for _, movie := range movies {
		kts_err := ec.EventRepo.AddEventMovie(eventId, movie)
		if kts_err != nil {
			log.Printf("Error adding event movie: %v", kts_err.ErrorMessage)
			return nil, kts_err
		}
	}

	eventSeatCategories := eventDto.EventSeatCategories

	if len(eventSeatCategories) == 0 {
		log.Printf("No event seat categories provided for event: %v", eventId)
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	for _, eventSeatCategory := range eventSeatCategories {
		eventSeatCategory.EventID = *eventId
		kts_err := ec.EventRepo.CreateEventSeatCategory(&eventSeatCategory)
		if kts_err != nil {
			log.Printf("Error creating event seat category: %v", kts_err.ErrorMessage)
			return nil, kts_err
		}
	}

	kts_err = ec.createEventSeatsFast(&eventDto.CinemaHallID, eventId)
	if kts_err != nil {
		log.Printf("Error creating event seats: %v", kts_err)
		return nil, kts_err
	}

	return eventId, nil
}

func (ec *EventController) GetEventsForMovie(movieId *myid.UUID, theatreId *myid.UUID) ([]*model.Events, *models.KTSError) {
	events, err := ec.EventRepo.GetEventsForMovie(movieId, theatreId)
	if err != nil {
		log.Printf("Error getting events for movie: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return events, nil
}

func (ec *EventController) GetSpecialEvents() (*[]models.GetSpecialEventsDTO, *models.KTSError) {
	specialEvents, err := ec.EventRepo.GetSpecialEvents()
	if err != nil {
		log.Printf("Error getting special events: %v", err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return specialEvents, nil
}

func (ec *EventController) createEventSeats(cinemaHallId *myid.UUID, eventId *myid.UUID) *models.KTSError {
	seats, kts_err := ec.TheatreRepo.GetSeatsForCinemaHall(cinemaHallId)
	if kts_err != nil {
		return kts_err
	}

	for _, seat := range seats {
		eventSeatId := myid.NewUUID()
		eventSeat := &model.EventSeats{
			ID:      *eventSeatId,
			EventID: *eventId,
			SeatID:  seat.ID,
			Booked:  false,
		}
		kts_err := ec.EventRepo.CreateEventSeat(eventSeat)
		if kts_err != nil {
			return kts_err
		}
	}

	return nil
}

func (ec *EventController) createEventSeatsFast(cinemaHallId *myid.UUID, eventId *myid.UUID) *models.KTSError {
	seats, kts_err := ec.TheatreRepo.GetSeatsForCinemaHall(cinemaHallId)
	if kts_err != nil {
		return kts_err
	}

	var eventSeats []*model.EventSeats
	for _, seat := range seats {
		eventSeatId := myid.NewUUID()
		eventSeat := &model.EventSeats{
			ID:      *eventSeatId,
			EventID: *eventId,
			SeatID:  seat.ID,
			Booked:  false,
		}
		eventSeats = append(eventSeats, eventSeat)
	}
		kts_err = ec.EventRepo.CreateEventSeats(eventSeats)
		if kts_err != nil {
			return kts_err
		}

	return nil
}

func (ec *EventController) GetEventById(eventId *myid.UUID) (*models.GetSpecialEventsDTO, *models.KTSError) {
	return ec.EventRepo.GetEventById(eventId)
}
