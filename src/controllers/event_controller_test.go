package controllers

import (
	"errors"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestEventController_CreateEvent(t *testing.T) {

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

	tests := []struct {
		name             string
		expectFuncs      func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T)
		expectedErrorMsg string
	}{
		{
			name: "ExpectCreateEventWorks",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieWorks(mockEventRepo, t)
				ExpectCreateEventSeatCategory(mockEventRepo, t)
				ExpectGetSeatsForCinemaHallWorks(mockTheatreRepo, t)
				ExpectCreateEventSeatWorks(mockEventRepo, t)
			},
			expectedErrorMsg: "",
		}, {
			name: "ExpectGetGenreByNameReturnsNilSoGenreHasToBeCreated",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExpectGetGenreByNameReturnsNil(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectCreateGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieWorks(mockEventRepo, t)
				ExpectCreateEventSeatCategory(mockEventRepo, t)
				ExpectGetSeatsForCinemaHallWorks(mockTheatreRepo, t)
				ExpectCreateEventSeatWorks(mockEventRepo, t)
			},
			expectedErrorMsg: "",
		},
		{
			name: "CreateEventReturnsNil",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventReturnsNil(mockEventRepo, eventRequest, t)
			},
			expectedErrorMsg: "Event creation failed",
		},
		{
			name: "CreateEventReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventReturnsError(mockEventRepo, eventRequest, t)
			},
			expectedErrorMsg: "Event creation failed",
		},
		{
			name: "ExpectCreateMovieReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieReturnsError(mockMovieRepo, t)
			},
			expectedErrorMsg: "Movie creation failed",
		},
		{
			name: "ExpectGetGenreByNameReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameReturnsError(mockMovieRepo, t)
			},
			expectedErrorMsg: "Genre creation failed",
		},
		{
			name: "ExcpectGetGenreByNameReturnsNilSoGenreHasToBeCreatedButCreateGenreReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExpectGetGenreByNameReturnsNil(mockMovieRepo, t)
				ExpectCreateGenreReturnsError(mockMovieRepo, t)
			},
			expectedErrorMsg: "Genre creation failed",
		},
		{
			name: "ExpectAddMovieGenreReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreReturnsError(mockMovieRepo, t)
			},
			expectedErrorMsg: "Movie genre creation failed",
		},
		{
			name: "ExpectAddEventMovieReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieReturnsError(mockEventRepo, t)
			},
			expectedErrorMsg: "Event movie creation failed",
		},
		{
			name: "ExpectCreateEventSeatCategoryReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {

				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieWorks(mockEventRepo, t)
				ExpectCreateEventSeatCategoryReturnsError(mockEventRepo, t)
			},
			expectedErrorMsg: "Event seat category creation failed",
		},
		{
			name: "ExpectGetSeatsForCinemaHallReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {

				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieWorks(mockEventRepo, t)
				ExpectCreateEventSeatCategory(mockEventRepo, t)
				ExpectGetSeatsForCinemaHallReturnsError(mockTheatreRepo, t)
			},
			expectedErrorMsg: "Event seat creation failed",
		},
		{
			name: "ExpectCreateEventSeatWorksReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, mockMovieRepo *mocks.MockMovieRepoI, mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {

				ExpectCreateEventWorks(mockEventRepo, eventRequest, t)
				ExpectCreateMovieWorks(mockMovieRepo, t)
				ExcpectGetGenreByNameWorks(mockMovieRepo, t)
				ExpectAddMovieGenreWorks(mockMovieRepo, t)
				ExpectAddEventMovieWorks(mockEventRepo, t)
				ExpectCreateEventSeatCategory(mockEventRepo, t)
				ExpectGetSeatsForCinemaHallWorks(mockTheatreRepo, t)
				ExpectCreateEventSeatReturnsError(mockEventRepo, t)
			},
			expectedErrorMsg: "Event seat creation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
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
			event, err := eventController.CreateEvent(eventRequest)

			// then
			if tt.expectedErrorMsg != "" {
				if err == nil {
					t.Errorf("Expected error: %v, but got nil", tt.expectedErrorMsg)
				}
				if event != nil {
					t.Errorf("Expected event to be nil, but got: %v", event)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if event == nil {
					t.Errorf("Expected event to be not nil, but got nil")
				}
			}
		})
	}
}

func TestEventController_DeleteEvent(t *testing.T) {
	eventId := uuid.New()

	tests := []struct {
		name          string
		expectFuncs   func(mockEventRepo *mocks.MockEventRepo, t *testing.T)
		expectedError bool
	}{
		{
			name: "ExpectDeleteEventWorks",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectDeleteEventWorks(mockEventRepo, t)
				ExpectDeleteEventMoviesWorks(mockEventRepo, t)
				ExpectDeleteEventSeatCategoryByEventIdWorks(mockEventRepo, t)
				ExpectDeleteEventSeatsByEventIdWorks(mockEventRepo, t)
			},
			expectedError: false,
		},
		{
			name: "ExpectDeleteEventReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectDeleteEventReturnsError(mockEventRepo, t)
			},
			expectedError: true,
		},
		{
			name: "ExpectDeleteEventMoviesReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectDeleteEventWorks(mockEventRepo, t)
				ExpectDeleteEventMoviesReturnsError(mockEventRepo, t)
			},
			expectedError: true,
		},
		{
			name: "ExpectDeleteEventSeatCategoryByEventIdReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectDeleteEventWorks(mockEventRepo, t)
				ExpectDeleteEventMoviesWorks(mockEventRepo, t)
				ExpectDeleteEventSeatCategoryByEventIdReturnsError(mockEventRepo, t)
			},
			expectedError: true,
		},
		{
			name: "ExpectDeleteEventSeatsByEventIdReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectDeleteEventWorks(mockEventRepo, t)
				ExpectDeleteEventMoviesWorks(mockEventRepo, t)
				ExpectDeleteEventSeatCategoryByEventIdWorks(mockEventRepo, t)
				ExpectDeleteEventSeatsByEventIdReturnsError(mockEventRepo, t)
			},
			expectedError: true,
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
			err := eventController.DeleteEvent(&eventId)

			// then
			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
				if err != kts_errors.KTS_INTERNAL_ERROR {
					t.Errorf("Expected error: %v, but got: %v", kts_errors.KTS_INTERNAL_ERROR, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func ExpectCreateEventWorks(mockEventRepo *mocks.MockEventRepo, eventRequest *models.EventDTO, t *testing.T) {
	mockEventRepo.EXPECT().CreateEvent(gomock.Any()).DoAndReturn(func(event *schemas.Event) (*schemas.Event, error) {
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
}

func ExpectCreateEventReturnsNil(mockEventRepo *mocks.MockEventRepo, eventRequest *models.EventDTO, t *testing.T) {
	mockEventRepo.EXPECT().CreateEvent(gomock.Any()).Return(nil, nil)
}

func ExpectCreateEventReturnsError(mockEventRepo *mocks.MockEventRepo, eventRequest *models.EventDTO, t *testing.T) {
	mockEventRepo.EXPECT().CreateEvent(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExpectCreateMovieWorks(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().CreateMovie(gomock.Any()).DoAndReturn(func(movie *schemas.Movie) (*schemas.Movie, error) {
		if movie == nil {
			t.Errorf("Movie is nil")
		}
		if movie.Id == nil {
			t.Errorf("Movie Id is nil")
		}
		return movie, nil
	})
}

func ExpectCreateMovieReturnsError(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().CreateMovie(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExpectCreateEventSeatCategory(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().CreateEventSeatCategory(gomock.Any()).DoAndReturn(func(eventSeatCategory *schemas.EventSeatCategory) (*schemas.EventSeatCategory, error) {
		if eventSeatCategory == nil {
			t.Errorf("EventSeatCategory is nil")
		}
		return eventSeatCategory, nil
	}).AnyTimes()
}

func ExpectCreateEventSeatCategoryReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().CreateEventSeatCategory(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExcpectGetGenreByNameWorks(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().GetGenreByName(gomock.Any()).DoAndReturn(func(genreName string) (*schemas.Genre, error) {
		id := uuid.New()
		return &schemas.Genre{
			Id:        &id,
			GenreName: genreName,
		}, nil
	}).AnyTimes()
}

func ExcpectGetGenreByNameReturnsError(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().GetGenreByName(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExpectGetGenreByNameReturnsNil(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().GetGenreByName(gomock.Any()).Return(nil, nil).AnyTimes()
}

func ExpectCreateGenreWorks(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().CreateGenre(gomock.Any()).DoAndReturn(func(genre *schemas.Genre) (*schemas.Genre, error) {
		if genre == nil {
			t.Errorf("Genre is nil")
		}
		if genre.Id == nil {
			t.Errorf("Genre Id is nil")
		}
		return genre, nil
	}).AnyTimes()
}

func ExpectCreateGenreReturnsError(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().CreateGenre(gomock.Any()).Return(nil, errors.New("Error"))
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

func ExpectAddMovieGenreReturnsError(mockMovieRepo *mocks.MockMovieRepoI, t *testing.T) {
	mockMovieRepo.EXPECT().AddMovieGenre(gomock.Any(), gomock.Any()).Return(errors.New("Error"))
}

func ExpectAddEventMovieWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().AddEventMovie(gomock.Any(), gomock.Any()).DoAndReturn(func(eventId *uuid.UUID, movieId *uuid.UUID) error {
		if eventId == nil {
			t.Errorf("Event Id is nil")
		}
		if movieId == nil {
			t.Errorf("Movie Id is nil")
		}
		return nil
	}).AnyTimes()
}

func ExpectAddEventMovieReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().AddEventMovie(gomock.Any(), gomock.Any()).Return(errors.New("Error"))
}

func ExpectGetSeatsForCinemaHallWorks(mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
	mockTheatreRepo.EXPECT().GetSeatsForCinemaHall(gomock.Any()).DoAndReturn(func(cinemaHallId *uuid.UUID) ([]schemas.Seat, error) {
		if cinemaHallId == nil {
			t.Errorf("CinemaHall Id is nil")
		}

		ids := uuid.New()

		return []schemas.Seat{
			{
				Id:             &ids,
				Row:            1,
				Column:         1,
				CinemaHallId:   cinemaHallId,
				SeatCategoryId: &ids,
			},
		}, nil
	}).AnyTimes()
}

func ExpectGetSeatsForCinemaHallReturnsError(mockTheatreRepo *mocks.MockTheaterRepoI, t *testing.T) {
	mockTheatreRepo.EXPECT().GetSeatsForCinemaHall(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExpectCreateEventSeatWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().CreateEventSeat(gomock.Any()).DoAndReturn(func(eventSeat *schemas.EventSeat) (*schemas.EventSeat, error) {
		if eventSeat == nil {
			t.Errorf("EventSeat is nil")
		}
		return eventSeat, nil
	}).AnyTimes()
}

func ExpectCreateEventSeatReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().CreateEventSeat(gomock.Any()).Return(nil, errors.New("Error"))
}

func ExpectCreateEventSeatCategoryWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().CreateEventSeatCategory(gomock.Any()).DoAndReturn(func(eventSeatCategory *schemas.EventSeatCategory) (*schemas.EventSeatCategory, error) {
		if eventSeatCategory == nil {
			t.Errorf("EventSeatCategory is nil")
		}
		return eventSeatCategory, nil
	}).AnyTimes()
}

func ExpectDeleteEventWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEvent(gomock.Any()).DoAndReturn(func(eventId *uuid.UUID) error {
		if eventId == nil {
			t.Errorf("Event Id is nil")
		}
		return nil
	}).AnyTimes()
}

func ExpectDeleteEventReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEvent(gomock.Any()).Return(errors.New("Error")).AnyTimes()
}

func ExpectDeleteEventMoviesWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventMovies(gomock.Any()).DoAndReturn(func(eventId *uuid.UUID) error {
		if eventId == nil {
			t.Errorf("Event Id is nil")
		}
		return nil
	}).AnyTimes()
}

func ExpectDeleteEventSeatCategoryByEventIdReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventSeatCategoryByEventId(gomock.Any()).Return(errors.New("Error")).AnyTimes()
}

func ExpectDeleteEventSeatCategoryByEventIdWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventSeatCategoryByEventId(gomock.Any()).DoAndReturn(func(eventId *uuid.UUID) error {
		if eventId == nil {
			t.Errorf("Event Id is nil")
		}
		return nil
	}).AnyTimes()
}

func ExpectDeleteEventSeatsByEventIdReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventSeatsByEventId(gomock.Any()).Return(errors.New("Error")).AnyTimes()
}

func ExpectDeleteEventSeatsByEventIdWorks(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventSeatsByEventId(gomock.Any()).DoAndReturn(func(eventId *uuid.UUID) error {
		if eventId == nil {
			t.Errorf("Event Id is nil")
		}
		return nil
	}).AnyTimes()
}

func ExpectDeleteEventMoviesReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().DeleteEventMovies(gomock.Any()).Return(errors.New("Error")).AnyTimes()
}
func TestEventController_GetEventsForMovie(t *testing.T) {
	movieId := uuid.New()

	tests := []struct {
		name          string
		expectFuncs   func(mockEventRepo *mocks.MockEventRepo, t *testing.T)
		expectedError *models.KTSError
	}{
		{
			name: "ExpectGetEventsForMovieIdReturnsEvents",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectGetEventsForMovieIdReturnsEvents(mockEventRepo, t)
			},
			expectedError: nil,
		},
		{
			name: "ExpectGetEventsForMovieIdReturnsError",
			expectFuncs: func(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
				ExpectGetEventsForMovieIdReturnsError(mockEventRepo, t)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
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
			events, err := eventController.GetEventsForMovie(&movieId)

			// then
			if err != tt.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tt.expectedError, err)
			}
			if tt.expectedError == nil {
				if events == nil {
					t.Errorf("Expected events to be not nil, but got nil")
				}
				if len(events) != 2 {
					t.Errorf("Expected events to have length 2, but got: %v", len(events))
				}
			} else {
				if events != nil {
					t.Errorf("Expected events to be nil, but got: %v", events)
				}
			}
		})
	}
}

func ExpectGetEventsForMovieIdReturnsEvents(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().GetEventsForMovieId(gomock.Any()).DoAndReturn(func(movieId *uuid.UUID) ([]*schemas.Event, error) {
		if movieId == nil {
			t.Errorf("Movie Id is nil")
		}
		id := uuid.New()
		return []*schemas.Event{
			{
				Id:           &id,
				Title:        "Test Event 1",
				Start:        time.Now(),
				End:          time.Now(),
				EventTypeId:  &id,
				CinemaHallId: &id,
			},
			{
				Id:           &id,
				Title:        "Test Event 2",
				Start:        time.Now(),
				End:          time.Now(),
				EventTypeId:  &id,
				CinemaHallId: &id,
			},
		}, nil
	}).AnyTimes()
}

func ExpectGetEventsForMovieIdReturnsError(mockEventRepo *mocks.MockEventRepo, t *testing.T) {
	mockEventRepo.EXPECT().GetEventsForMovieId(gomock.Any()).Return(nil, errors.New("Error"))
}
