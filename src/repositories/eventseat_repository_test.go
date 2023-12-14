package repositories

import (
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func GetEventSeats() *[]models.GetEventSeatsDTO {
	// Create event seats for this DTO

	seatId := utils.NewUUID()
	seatCategoryId := utils.NewUUID()

	eventSeats := []models.GetEventSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      utils.NewUUID(),
				SeatID:       seatId,
			},
			Seat: model.Seats{
				ID:       seatId,
				RowNr:    1,
				ColumnNr: 1,

				SeatCategoryID: seatCategoryId,
			},
			SeatCategory: model.SeatCategories{
				ID:           seatCategoryId,
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        utils.NewUUID(),
				SeatCategoryID: seatCategoryId,
				Price:          100,
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      utils.NewUUID(),
				SeatID:       seatId,
			},
			Seat: model.Seats{
				ID:             seatId,
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: seatCategoryId,
			},
			SeatCategory: model.SeatCategories{
				ID:           seatCategoryId,
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        utils.NewUUID(),
				SeatCategoryID: seatCategoryId,
				Price:          100,
			},
		},
	}

	return &eventSeats
}

func TestGetEventSeats(t *testing.T) {
	eventSeats := *GetEventSeats()

	query := "SELECT .* FROM `KinoTicketSystem`.event_seats .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedSeats   *[]models.GetEventSeatsDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Select event seats",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"event_seats.id", "event_seats.booked", "event_seats.blocked_until", "event_seats.user_id", "event_seats.seat_id", "event_seats.event_id", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "seat_categories.id", "seat_categories.category_name", "event_seat_categories.event_id", "event_seat_categories.seat_category_id", "event_seat_categories.price"}).
						AddRow(eventSeats[0].EventSeat.ID, eventSeats[0].EventSeat.Booked, eventSeats[0].EventSeat.BlockedUntil, eventSeats[0].EventSeat.UserID, eventSeats[0].EventSeat.SeatID, eventSeats[0].EventSeat.EventID, eventSeats[0].Seat.ID, eventSeats[0].Seat.RowNr, eventSeats[0].Seat.ColumnNr, eventSeats[0].Seat.SeatCategoryID, eventSeats[0].Seat.CinemaHallID, eventSeats[0].Seat.Type, eventSeats[0].SeatCategory.ID, eventSeats[0].SeatCategory.CategoryName, eventSeats[0].EventSeatCategory.EventID, eventSeats[0].EventSeatCategory.SeatCategoryID, eventSeats[0].EventSeatCategory.Price).
						AddRow(eventSeats[1].EventSeat.ID, eventSeats[1].EventSeat.Booked, eventSeats[1].EventSeat.BlockedUntil, eventSeats[1].EventSeat.UserID, eventSeats[1].EventSeat.SeatID, eventSeats[1].EventSeat.EventID, eventSeats[1].Seat.ID, eventSeats[1].Seat.RowNr, eventSeats[1].Seat.ColumnNr, eventSeats[1].Seat.SeatCategoryID, eventSeats[1].Seat.CinemaHallID, eventSeats[1].Seat.Type, eventSeats[1].SeatCategory.ID, eventSeats[1].SeatCategory.CategoryName, eventSeats[1].EventSeatCategory.EventID, eventSeats[1].EventSeatCategory.SeatCategoryID, eventSeats[1].EventSeatCategory.Price),
					)

			},
			expectedSeats: &eventSeats,
			expectedError: nil,
		},
		{
			name: "Select event seats - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Select event seats - no rows",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"event_seats.id", "event_seats.booked", "event_seats.blocked_until", "event_seats.user_id", "event_seats.seat_id", "event_seats.event_id", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "seat_categories.id", "seat_categories.category_name", "event_seat_categories.event_id", "event_seat_categories.seat_category_id", "event_seat_categories.price"}))
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the EventSeatRepository with the mock database connection
			eventSeatRepo := &EventSeatRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method on the repository instance
			seats, kts_err := eventSeatRepo.GetEventSeats(eventSeats[0].EventSeat.EventID)

			// Verify that all expectations were met

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedSeats, seats)

		})
	}
}

func TestBlockEventSeatIfAvailable(t *testing.T) {
	eventId := utils.NewUUID()
	seatId := utils.NewUUID()
	userId := utils.NewUUID()
	blockedUntil := time.Now()

	query := "UPDATE `KinoTicketSystem`.event_seats SET .* WHERE .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Block event seat if available",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(&blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Block event seat if available - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(&blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0))

			},
			expectedError: kts_errors.KTS_CONFLICT,
		},
		{
			name: "Block event seat if available - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(&blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			db, mock, err := sqlmock.New()
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

			// When
			ktsErr := eventRepo.BlockEventSeatIfAvailable(eventId, seatId, userId, &blockedUntil)

			// Then
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedError, ktsErr)
		})
	}
}
func TestUpdateBlockedUntilTimeForUserEventSeats(t *testing.T) {
	eventId := utils.NewUUID()
	userId := utils.NewUUID()
	blockedUntil := time.Now()

	query := "UPDATE `KinoTicketSystem`.event_seats SET .* WHERE .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Update blocked until time for user event seats",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Update blocked until time for user event seats - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Update blocked until time for user event seats - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			esr := &EventSeatRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// When
			ktsErr := esr.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

			// Then
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedError, ktsErr)
		})
	}
}