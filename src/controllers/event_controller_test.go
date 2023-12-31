package controllers

import (
	"reflect"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
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
		expectFuncs     func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T)
		expectedError   *models.KTSError
		expectedEventId *uuid.UUID
		eventRequest    *models.CreateEvtDTO
	}{
		{
			name: "Create Event",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Add Event Movie returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				mockEventRepo.EXPECT().CreateEvent(&eventRequest.Events).Return(eventRequest.Events.ID, nil)
				mockEventRepo.EXPECT().AddEventMovie(eventRequest.Events.ID, eventRequest.Movies[0]).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
			expectedEventId: nil,
			eventRequest:    eventRequest,
		},
		{
			name: "Create Event Seat Category returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
			},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
			expectedEventId: nil,
			eventRequest:    nil,
		},
		{
			name: "Movies nil or empty",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
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

			mockEventRepo := mocks.NewMockEventRepo(mockCtrl)
			mockTheatreRepo := mocks.NewMockTheaterRepoI(mockCtrl)

			tt.expectFuncs(mockEventRepo, mockTheatreRepo, t)

			eventController := &EventController{
				EventRepo:   mockEventRepo,
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

func TestEventController_GetEventsForMovie(t *testing.T) {
	movieId := utils.NewUUID()
	theatreId := utils.NewUUID()

	expectedEvents := samples.GetModelEvents()

	tests := []struct {
		name           string
		expectFuncs    func(mockEventRepo *mocks.MockEventRepo, t *testing.T)
		expectedError  *models.KTSError
		expectedEvents []*model.Events
	}{
		{
			name: "Get Events for Movie",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetEventsForMovie(movieId, theatreId).Return(expectedEvents, nil)
			},
			expectedError:  nil,
			expectedEvents: expectedEvents,
		},
		{
			name: "Get Events for Movie returns error",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				mockEventRepo.EXPECT().GetEventsForMovie(movieId, theatreId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			events, ktsErr := eventController.GetEventsForMovie(movieId, theatreId)

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
func TestEventController_GetEventById(t *testing.T) {
	eventId := utils.NewUUID()

	expectedEvent := &models.GetSpecialEventsDTO{
		Events: model.Events{
			ID:          eventId,
			Title:       "Test Event",
			Description: utils.GetStringPointer("Test Description"),
		},
	}

	tests := []struct {
		name             string
		eventId          *uuid.UUID
		expectedEventDTO *models.GetSpecialEventsDTO
		expectedError    *models.KTSError
	}{
		{
			name:             "Get event by id",
			eventId:          eventId,
			expectedEventDTO: expectedEvent,
			expectedError:    nil,
		},
		{
			name:             "Get event by id fails",
			eventId:          eventId,
			expectedEventDTO: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventRepo := mocks.NewMockEventRepo(mockCtrl)

			mockEventRepo.EXPECT().GetEventById(tt.eventId).Return(tt.expectedEventDTO, tt.expectedError)

			eventController := &EventController{
				EventRepo: mockEventRepo,
			}

			// when
			eventDTO, err := eventController.GetEventById(tt.eventId)

			// then
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError, err)
			}

			if !reflect.DeepEqual(eventDTO, tt.expectedEventDTO) {
				t.Errorf("Expected event DTO: %v, but got: %v", tt.expectedEventDTO, eventDTO)
			}
		})
	}
}
