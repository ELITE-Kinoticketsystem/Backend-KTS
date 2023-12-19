package repositories

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestCreateEvent(t *testing.T) {

	description := "Test event description"

	event := &model.Events{
		Title:        "Test Event",
		Start:        time.Now(),
		End:          time.Now().Add(time.Hour),
		Description:  &description,
		EventType:    "Test event type",
		CinemaHallID: utils.NewUUID(),
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectEventID   bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create event",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.events .*").WithArgs(sqlmock.AnyArg(), event.Title, sqlmock.AnyArg(), sqlmock.AnyArg(), event.Description, event.EventType, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectEventID: true,
			expectedError: nil,
		},
		{
			name: "Create event sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.events .*").WithArgs(sqlmock.AnyArg(), event.Title, sqlmock.AnyArg(), sqlmock.AnyArg(), event.Description, event.EventType, sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)
			},
			expectEventID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create event no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.events .*").WithArgs(sqlmock.AnyArg(), event.Title, sqlmock.AnyArg(), sqlmock.AnyArg(), event.Description, event.EventType, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectEventID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			id, ktsErr := eventRepo.CreateEvent(event)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if tc.expectEventID && id == nil {
				t.Error("Expected event ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
func TestCreateEventSeatCategory(t *testing.T) {
	eventSeatCategory := &model.EventSeatCategories{
		EventID:        utils.NewUUID(),
		SeatCategoryID: utils.NewUUID(),
		Price:          999,
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectError     bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create event seat category",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seat_categories .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeatCategory.Price).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectError:   false,
			expectedError: nil,
		},
		{
			name: "Create event seat category sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seat_categories .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeatCategory.Price).WillReturnError(sql.ErrConnDone)
			},
			expectError:   true,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create event seat category no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seat_categories .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeatCategory.Price).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectError:   true,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			kts_err := eventRepo.CreateEventSeatCategory(eventSeatCategory)

			if (kts_err != nil) != tc.expectError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			if kts_err != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, kts_err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestAddEventMovie(t *testing.T) {
	eventID := utils.NewUUID()
	movieID := utils.NewUUID()

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Add event movie",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_movies .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Add event movie sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_movies .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Add event movie no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_movies .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			ktsErr := eventRepo.AddEventMovie(eventID, movieID)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreateEventSeat(t *testing.T) {
	eventSeat := &model.EventSeats{
		ID:      utils.NewUUID(),
		EventID: utils.NewUUID(),
		SeatID:  utils.NewUUID(),
		Booked:  false,
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Create event seat",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seats .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeat.Booked).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Create event seat sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seats .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeat.Booked).WillReturnError(sql.ErrConnDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create event seat no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.event_seats .*").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), eventSeat.Booked).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			ktsErr := eventRepo.CreateEventSeat(eventSeat)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
func TestGetEventsForMovie(t *testing.T) {
	movieID := utils.NewUUID()

	expectedEvents := []*model.Events{
		{
			ID:           utils.NewUUID(),
			Title:        "Test Event 1",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  utils.GetStringPointer("Test event description 1"),
			EventType:    "showing",
			CinemaHallID: utils.NewUUID(),
		},
		{
			ID:           utils.NewUUID(),
			Title:        "Test Event 2",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  utils.GetStringPointer("Test event description 2"),
			EventType:    "showing",
			CinemaHallID: utils.NewUUID(),
		},
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedEvents  []*model.Events
		expectedError   *models.KTSError
	}{
		{
			name: "Get events for movie",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.events .*").WithArgs(sqlmock.AnyArg()).WillReturnRows(
					sqlmock.NewRows([]string{"events.id", "events.title", "events.start", "events.end", "events.description", "events.event_type", "events.cinema_hall_id"}).
						AddRow(expectedEvents[0].ID, expectedEvents[0].Title, expectedEvents[0].Start, expectedEvents[0].End, expectedEvents[0].Description, expectedEvents[0].EventType, expectedEvents[0].CinemaHallID).
						AddRow(expectedEvents[1].ID, expectedEvents[1].Title, expectedEvents[1].Start, expectedEvents[1].End, expectedEvents[1].Description, expectedEvents[1].EventType, expectedEvents[1].CinemaHallID),
				)
				
			},
			expectedEvents: expectedEvents,
			expectedError:  nil,
		},
		{
			name: "Get events for movie sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.events .*").WithArgs(sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)
			},
			expectedEvents: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			events, ktsErr := eventRepo.GetEventsForMovie(movieID)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if !reflect.DeepEqual(events, tc.expectedEvents) {
				t.Errorf("Expected events: %v, got: %v", tc.expectedEvents, events)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
func TestGetSpecialEvents(t *testing.T) {
	expectedSpecialEvents := []models.GetSpecialEventsDTO{
		{
			Events: model.Events{
				ID:          utils.NewUUID(),
				Title:       "Special Event 1",
				Start:       time.Now(),
				End:         time.Now().Add(time.Hour),
				Description: utils.GetStringPointer("Special event description 1"),
			},
			Movies: []*model.Movies{
				{
					ID:    utils.NewUUID(),
					Title: "Movie 1",
				},
			},
		},
		{
			Events: model.Events{
				ID:          utils.NewUUID(),
				Title:       "Special Event 2",
				Start:       time.Now(),
				End:         time.Now().Add(time.Hour),
				Description: utils.GetStringPointer("Special event description 2"),
			},
			Movies: []*model.Movies{
				{
					ID:    utils.NewUUID(),
					Title: "Movie 2",
				},
			},
		},
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedEvents  *[]models.GetSpecialEventsDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Get special events",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.events .*").WillReturnRows(
					sqlmock.NewRows([]string{"events.id", "events.title", "events.start", "events.end", "events.description", "movies.id", "movies.title"}).
						AddRow(expectedSpecialEvents[0].ID, expectedSpecialEvents[0].Title, expectedSpecialEvents[0].Start, expectedSpecialEvents[0].End, expectedSpecialEvents[0].Description, expectedSpecialEvents[0].Movies[0].ID, expectedSpecialEvents[0].Movies[0].Title).
						AddRow(expectedSpecialEvents[1].ID, expectedSpecialEvents[1].Title, expectedSpecialEvents[1].Start, expectedSpecialEvents[1].End, expectedSpecialEvents[1].Description, expectedSpecialEvents[1].Movies[0].ID, expectedSpecialEvents[1].Movies[0].Title),
				)
			},
			expectedEvents: &expectedSpecialEvents,
			expectedError:  nil,
		},
		{
			name: "Get special events sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.events .*").WillReturnError(sql.ErrConnDone)
			},
			expectedEvents: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Get special events no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.events .*").WillReturnRows(
					sqlmock.NewRows([]string{"events.id", "events.title", "events.start", "events.end", "events.description", "movies.id", "movies.title"}),
				)
			},
			expectedEvents: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventRepo := &EventRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			specialEvents, ktsErr := eventRepo.GetSpecialEvents()

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if !reflect.DeepEqual(specialEvents, tc.expectedEvents) {
				t.Errorf("Expected special events: %v, got: %v", tc.expectedEvents, specialEvents)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
