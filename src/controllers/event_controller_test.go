package controllers

import (
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestEventController_CreateEvent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dbManager := mocks.NewMockDatabaseManagerI(mockCtrl)
	mockMovieRepo := mocks.NewMockMovieRepoI(mockCtrl)
	mockEventRepo := mocks.NewMockEventRepo(mockCtrl)
	mockTheatreRepo := mocks.NewMockTheaterRepoI(mockCtrl)

	movieRequest1 := &models.MovieDTO{
		Id:          nil,
		Title:       "Test Event",
		Description: "Test Description",
		ReleaseDate: time.Now(),
		TimeInMin:   120,
		Fsk:         12,
		GenreNames:  []string{"Action", "Comedy"},
	}

	movieRequest2Id := uuid.New()

	movieRequest2 := &models.MovieDTO{
		Id:          &movieRequest2Id,
		Title:       "Test Event 2",
		Description: "Test Description 2",
		ReleaseDate: time.Now(),
		TimeInMin:   120,
		Fsk:         12,
		GenreNames:  []string{"Action", "Comedy"},
	}

	seatCategories := []models.EventSeatCategoryDTO{
		{
			Name:  "Premium",
			Price: 100,
		},
		{
			Name:  "Standard",
			Price: 50,
		},
	}

	EventTypeId := uuid.New()
	CinemaHallId := uuid.New()

	eventRequest := &models.EventDTO{
		Id:                  nil,
		Title:               "Test Event",
		Start:               time.Now(),
		End:                 time.Now(),
		EventTypeId:         &EventTypeId,
		CinemaHallId:        &CinemaHallId,
		Movies:              []models.MovieDTO{*movieRequest1, *movieRequest2},
		EventSeatCategories: seatCategories,
	}

	createEventFunc := mockEventRepo.EXPECT().CreateEvent(gomock.Any())
	createEventFunc.DoAndReturn(func(event *schemas.Event) (*schemas.Event, error) {
		// check if all values are set
		if event.Id == nil {
			t.Errorf("Event Id is nil")
		}
		if event.Title != eventRequest.Title {
			t.Errorf("Event Title is not equal")
		}
		if event.Start != eventRequest.Start {
			t.Errorf("Event Start is not equal")
		}
		if event.End != eventRequest.End {
			t.Errorf("Event End is not equal")
		}
		if event.EventTypeId != eventRequest.EventTypeId {
			t.Errorf("Event EventTypeId is not equal")
		}
		if event.CinemaHallId != eventRequest.CinemaHallId {
			t.Errorf("Event CinemaHallId is not equal")
		}
		return event, nil
	})

	

	eventController := &EventController{
		EventRepo:       mockEventRepo,
		MovieRepo:       mockMovieRepo,
		DatabaseManager: dbManager,
		TheaterRepo:     mockTheatreRepo,
	}

	err := eventController.CreateEvent(eventRequest)

	if err != nil {
		t.Errorf("Error while creating event: %v", err)
	}
}
