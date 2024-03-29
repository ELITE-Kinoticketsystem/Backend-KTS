package repositories

import (
	"errors"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func GetEventSeats() *[]models.GetEventSeatsDTO {
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
				DatabaseManagerI: &managers.DatabaseManager{
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

func TestGetHallDimensions(t *testing.T) {
	eventId := utils.NewUUID()
	width := int32(10)
	height := int32(10)

	query := "SELECT cinema_halls.width, cinema_halls.height FROM events LEFT JOIN cinema_halls ON cinema_halls.id = events.cinema_hall_id WHERE events.id = ?"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedWidth   int32
		expectedHeight  int32
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(eventId[:]).
					WillReturnRows(sqlmock.NewRows([]string{"width", "height"}).
						AddRow(width, height),
					)
			},
			expectedWidth:  width,
			expectedHeight: height,
			expectedError:  nil,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(eventId[:]).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedWidth:  0,
			expectedHeight: 0,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "No rows",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(eventId[:]).
					WillReturnRows(sqlmock.NewRows([]string{"width", "height"}))
			},
			expectedWidth:  0,
			expectedHeight: 0,
			expectedError:  kts_errors.KTS_NOT_FOUND,
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
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// When
			width, height, kts_err := eventSeatRepo.GetHallDimensions(eventId)

			// Then
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedWidth, width)
			assert.Equal(t, tc.expectedHeight, height)
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
		{
			name: "Block event rowsaffected failed",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(&blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
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

			eventRepo := &EventSeatRepository{
				DatabaseManagerI: &managers.DatabaseManager{
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
		name                 string
		setExpectations      func(mock sqlmock.Sqlmock)
		expectedError        *models.KTSError
		expectedAffectedRows int64
	}{
		{
			name: "Update blocked until time for user event seats",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError:        nil,
			expectedAffectedRows: 1,
		},
		{
			name: "Update blocked until time for user event seats - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError:        nil,
			expectedAffectedRows: 0,
		},
		{
			name: "Update blocked until time for user event seats - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
			expectedAffectedRows: 0,
		},
		{
			name: "Update blocked rowsaffected failed",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(blockedUntil, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
			expectedAffectedRows: 0,
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
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// When
			rows_affected, ktsErr := esr.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

			// Then
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedAffectedRows, rows_affected)

			assert.Equal(t, tc.expectedError, ktsErr)
		})
	}
}

func TestUnblockEventSeat(t *testing.T) {
	eventId := utils.NewUUID()
	seatId := utils.NewUUID()
	userId := utils.NewUUID()

	query := "UPDATE `KinoTicketSystem`.event_seats SET .* WHERE .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Unblock event seat",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(nil, nil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Unblock event seat - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(nil, nil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Unblock event seat - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(nil, nil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Update blocked rowsaffected failed",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(nil, nil, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
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
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// When
			ktsErr := esr.UnblockEventSeat(eventId, seatId, userId)

			// Then
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedError, ktsErr)
		})
	}
}

func TestUnblockAllEventSeats(t *testing.T) {
	eventId := utils.NewUUID()
	userId := utils.NewUUID()

	query := "UPDATE" // `KinoTicketSystem`.event_seats SET .* WHERE .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(
					query,
				).
					WithArgs(
						nil,
						nil,
						utils.EqUUID(eventId),
						utils.EqUUID(userId),
					).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Not found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(
						nil,
						nil,
						utils.EqUUID(eventId),
						utils.EqUUID(userId),
					).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(
						nil,
						nil,
						utils.EqUUID(eventId),
						utils.EqUUID(userId),
					).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Rows affected internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).
					WithArgs(
						nil,
						nil,
						utils.EqUUID(eventId),
						utils.EqUUID(userId),
					).WillReturnResult(sqlmock.NewErrorResult(sqlmock.ErrCancelled))
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			esr := &EventSeatRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// WHEN
			ktsErr := esr.UnblockAllEventSeats(eventId, userId)

			// THEN
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			assert.Equal(t, tc.expectedError, ktsErr)
		})
	}
}

func TestGetSelectedSeats(t *testing.T) {
	eventId := utils.NewUUID()
	userId := utils.NewUUID()

	eventSeats := []models.GetSlectedSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       1,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
	}

	query := "SELECT .* FROM `KinoTicketSystem`.event_seats .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedSeats   *[]models.GetSlectedSeatsDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Select selected seats",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"event_seats.id", "event_seats.booked", "event_seats.blocked_until", "event_seats.user_id", "event_seats.seat_id", "event_seats.event_id", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "seat_categories.id", "seat_categories.category_name", "event_seat_categories.event_id", "event_seat_categories.seat_category_id", "event_seat_categories.price"}).
						AddRow(eventSeats[0].EventSeat.ID, eventSeats[0].EventSeat.Booked, eventSeats[0].EventSeat.BlockedUntil, eventSeats[0].EventSeat.UserID, eventSeats[0].EventSeat.SeatID, eventSeats[0].EventSeat.EventID, eventSeats[0].Seat.ID, eventSeats[0].Seat.RowNr, eventSeats[0].Seat.ColumnNr, eventSeats[0].Seat.SeatCategoryID, eventSeats[0].Seat.CinemaHallID, eventSeats[0].Seat.Type, eventSeats[0].SeatCategory.ID, eventSeats[0].SeatCategory.CategoryName, eventSeats[0].EventSeatCategory.EventID, eventSeats[0].EventSeatCategory.SeatCategoryID, eventSeats[0].EventSeatCategory.Price).
						AddRow(eventSeats[1].EventSeat.ID, eventSeats[1].EventSeat.Booked, eventSeats[1].EventSeat.BlockedUntil, eventSeats[1].EventSeat.UserID, eventSeats[1].EventSeat.SeatID, eventSeats[1].EventSeat.EventID, eventSeats[1].Seat.ID, eventSeats[1].Seat.RowNr, eventSeats[1].Seat.ColumnNr, eventSeats[1].Seat.SeatCategoryID, eventSeats[1].Seat.CinemaHallID, eventSeats[1].Seat.Type, eventSeats[1].SeatCategory.ID, eventSeats[1].SeatCategory.CategoryName, eventSeats[1].EventSeatCategory.EventID, eventSeats[1].EventSeatCategory.SeatCategoryID, eventSeats[1].EventSeatCategory.Price),
					)
			},
			expectedSeats: &eventSeats,
			expectedError: nil,
		},
		{
			name: "Select selected seats - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sqlmock.ErrCancelled)
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Select selected seats - no rows",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
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

			eventSeatRepo := &EventSeatRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			seats, ktsErr := eventSeatRepo.GetSelectedSeats(eventId, userId)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			assert.Equal(t, tc.expectedSeats, seats)
		})
	}
}

func TestUpdateEventSeat(t *testing.T) {
	blockUntil, _ := time.Parse("2006-01-01", "2019-01-01")

	eventSeat := &model.EventSeats{
		ID:           utils.NewUUID(),
		Booked:       false,
		BlockedUntil: &blockUntil,
		UserID:       utils.NewUUID(),
		EventID:      utils.NewUUID(),
		SeatID:       utils.NewUUID(),
	}

	query := "UPDATE `KinoTicketSystem`.event_seats SET booked = ?, blocked_until = CAST(? AS DATETIME), user_id = ?, seat_id = ?, event_id = ? WHERE event_seats.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, eventSeat *model.EventSeats)
		expectedError   *models.KTSError
	}{
		{
			name: "Update event seat",
			setExpectations: func(mock sqlmock.Sqlmock, eventSeat *model.EventSeats) {
				mock.ExpectExec(query).WithArgs(eventSeat.Booked, eventSeat.BlockedUntil, utils.EqUUID(eventSeat.UserID), utils.EqUUID(eventSeat.SeatID), utils.EqUUID(eventSeat.EventID), utils.EqUUID(eventSeat.ID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Update event seat - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, eventSeat *model.EventSeats) {
				mock.ExpectExec(query).WithArgs(eventSeat.Booked, eventSeat.BlockedUntil, utils.EqUUID(eventSeat.UserID), utils.EqUUID(eventSeat.SeatID), utils.EqUUID(eventSeat.EventID), utils.EqUUID(eventSeat.ID)).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Update event seat - error",
			setExpectations: func(mock sqlmock.Sqlmock, eventSeat *model.EventSeats) {
				mock.ExpectExec(query).WithArgs(eventSeat.Booked, eventSeat.BlockedUntil, utils.EqUUID(eventSeat.UserID), utils.EqUUID(eventSeat.SeatID), utils.EqUUID(eventSeat.EventID), utils.EqUUID(eventSeat.ID)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Update event rowsaffected failed",
			setExpectations: func(mock sqlmock.Sqlmock, eventSeat *model.EventSeats) {
				mock.ExpectExec(query).WithArgs(eventSeat.Booked, eventSeat.BlockedUntil, utils.EqUUID(eventSeat.UserID), utils.EqUUID(eventSeat.SeatID), utils.EqUUID(eventSeat.EventID), utils.EqUUID(eventSeat.ID)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			eventSeatRepo := &EventSeatRepository{
				DatabaseManagerI: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, eventSeat)

			kts_err := eventSeatRepo.UpdateEventSeat(eventSeat)

			assert.Equal(t, tc.expectedError, kts_err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
