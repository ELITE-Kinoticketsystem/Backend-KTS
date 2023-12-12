package controllers

import (
	"reflect"
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestEventController_CreateEvent(t *testing.T) {

	eventRequest := &models.CreateEvtDTO{
		Events: model.Events{
			Title:        "Test Event",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  nil,
			EventType:    "Test event type",
			CinemaHallID: utils.NewUUID(),
		},
		Movies: []*uuid.UUID{
			utils.NewUUID(),
		},
		EventSeatCategories: []model.EventSeatCategories{
			{
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
	}

	tests := []struct {
		name            string
		expectFuncs     func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T)
		expectedError   *models.KTSError
		expectedEventId *uuid.UUID
		eventRequest    *models.CreateEvtDTO
	}{
		{
			name: "Create Event",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(nil)
				mockEventRepo.EXPECT().CreateEventSeatCategory(&eventRequest.EventSeatCategories[0]).Return(nil)
				mockTheatreRepo.EXPECT().GetSeatsForCinemaHall(eventRequest.Events.CinemaHallID).Return([]model.Seats{
					{
						ID:             utils.NewUUID(),
						RowNr:          1,
						ColumnNr:       1,
						CinemaHallID:   eventRequest.Events.CinemaHallID,
						SeatCategoryID: eventRequest.EventSeatCategories[0].SeatCategoryID,
					},
				}, nil)
				mockEventRepo.EXPECT().CreateEventSeat(gomock.Any()).Return(nil)
			},
			expectedError:   nil,
			expectedEventId: eventRequest.Events.ID,
			eventRequest:    eventRequest,
		},
		{
			name: "Create Event returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Add Event Movie returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Create Event Seat Category returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(nil)
				mockEventRepo.EXPECT().CreateEventSeatCategory(&eventRequest.EventSeatCategories[0]).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Get Seats For Cinema Hall returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(nil)
				mockEventRepo.EXPECT().CreateEventSeatCategory(&eventRequest.EventSeatCategories[0]).Return(nil)
				mockTheatreRepo.EXPECT().GetSeatsForCinemaHall(eventRequest.Events.CinemaHallID).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Create Event Seat returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(nil)
				mockEventRepo.EXPECT().CreateEventSeatCategory(&eventRequest.EventSeatCategories[0]).Return(nil)
				mockTheatreRepo.EXPECT().GetSeatsForCinemaHall(eventRequest.Events.CinemaHallID).Return([]model.Seats{
					{
						ID:             utils.NewUUID(),
						RowNr:          1,
						ColumnNr:       1,
						CinemaHallID:   eventRequest.Events.CinemaHallID,
						SeatCategoryID: eventRequest.EventSeatCategories[0].SeatCategoryID,
					},
				}, nil)
				mockEventRepo.EXPECT().CreateEventSeat(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Nil eventDto",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
			},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
			expectedEventId: nil,
			eventRequest:    nil,
		},
		{
			name: "Movies nil or empty",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
			},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
			expectedEventId: nil,
			eventRequest: &models.CreateEvtDTO{
				Events: eventRequest.Events,
			},
		},
		{
			name: "EventSeatCategories nil or empty",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(nil)
			},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
			expectedEventId: nil,
			eventRequest: &models.CreateEvtDTO{
				Events: eventRequest.Events,
				Movies: eventRequest.Movies,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockMovieRepo := mocks.NewMockMovieRepoI(mockCtrl)
			mockEventRepo := mocks.NewMockEventRepo(mockCtrl)
			mockTheatreRepo := mocks.NewMockTheaterRepoI(mockCtrl)

			tt.expectFuncs(mockEventRepo, mockMovieRepo, mockTheatreRepo, t)

			eventController := &EventController{
				EventRepo:   mockEventRepo,
				MovieRepo:   mockMovieRepo,
				TheatreRepo: mockTheatreRepo,
			}

			// // when
			event, ktsErr := eventController.CreateEvent(tt.eventRequest)

			// then
			if ktsErr != tt.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError.ErrorMessage, ktsErr.ErrorMessage)
			}

			if tt.expectedEventId != event {
				t.Errorf("Expected event id: %v, but got: %v", tt.expectedEventId, event)
			}
		})
	}
}

func ExpectAddMovieGenreWorks(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().AddMovieGenre(gomock.Any(), gomock.Any()).DoAndReturn(func(movieId *uuid.UUID, genreId *uuid.UUID) error {
		if movieId == nil {
			t.Errorf("Movie Id is nil")
		}
		if genreId == nil {
			t.Errorf("Genre Id is nil")
		}
		return nil
	}).AnyTimes()
}

func TestEventController_GetEventsForMovie(t *testing.T) {
	movieId := utils.NewUUID()

	expectedEvents := []*model.Events{{
		ID:           utils.NewUUID(),
		Title:        "Test Event 1",
		Start:        time.Now(),
		End:          time.Now().Add(time.Hour),
		Description:  nil,
		EventType:    "Test event type 1",
		CinemaHallID: utils.NewUUID(),
	},
		{
			ID:           utils.NewUUID(),
			Title:        "Test Event 2",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  nil,
			EventType:    "Test event type 2",
			CinemaHallID: utils.NewUUID(),
		},
	}

	tests := []struct {
		name           string
		expectFuncs    func(mockEventRepo *mocks.MockEventRepo, t *testing.T)
		expectedError  *models.KTSError
		expectedEvents []*model.Events
	}{
		{
			name: "Get Events for Movie",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetEventsForMovie(movieId).Return(expectedEvents, nil)
			},
			expectedError:  nil,
			expectedEvents: expectedEvents,
		},
		{
			name: "Get Events for Movie returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetEventsForMovie(movieId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
			expectedEvents: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventRepo := mocks.NewMockEventRepo(mockCtrl)

			tt.expectFuncs(mockEventRepo, t)

			eventController := &EventController{
				EventRepo: mockEventRepo,
			}

			// when
			events, ktsErr := eventController.GetEventsForMovie(movieId)

			// then
			if ktsErr != tt.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError.ErrorMessage, ktsErr.ErrorMessage)
			}

			if !reflect.DeepEqual(events, tt.expectedEvents) {
				t.Errorf("Expected events: %v, but got: %v", tt.expectedEvents, events)
			}
		})
	}
}


func TestEventController_GetSpecialEvents(t *testing.T) {
	tests := []struct {
		name           string
		expectFuncs    func(mockEventRepo *mocks.MockEventRepo, t *testing.T)
		expectedResult *[]models.GetSpecialEventsDTO
		expectedError  *models.KTSError
	}{
		{
			name: "Get Special Events",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetSpecialEvents().Return(&[]models.GetSpecialEventsDTO{}, nil)
			},
			expectedResult: &[]models.GetSpecialEventsDTO{},
			expectedError:  nil,
		},
		{
			name: "Get Special Events returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetSpecialEvents().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResult: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventRepo := mocks.NewMockEventRepo(mockCtrl)

			tt.expectFuncs(mockEventRepo, t)

			eventController := &EventController{
				EventRepo: mockEventRepo,
			}

			// when
			result, err := eventController.GetSpecialEvents()

			// then
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Expected result: %v, but got: %v", tt.expectedResult, result)
			}

			if err != tt.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError, err)
			}
		})
	}
}