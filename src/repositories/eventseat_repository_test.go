package repositories

import (
	"database/sql"
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
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
		},
	}

	return &eventSeats
}

func TestGetEventSeats(t *testing.T) {
	eventSeats := *GetEventSeats()

	query := "\nSELECT event_seats.id AS \"event_seats.id\",\n     event_seats.event_id AS \"event_seats.event_id\",\n     event_seats.seat_id AS \"event_seats.seat_id\",\n     seats.id AS \"seats.id\",\n     seats.seat_category_id AS \"seats.seat_category_id\",\n     seat_categories.id AS \"seat_categories.id\",\n     seat_categories.name AS \"seat_categories.name\"\nFROM `KinoTicketSystem`.event_seats\n     LEFT JOIN `KinoTicketSystem`.seats ON (event_seats.seat_id = seats.id)\n     LEFT JOIN `KinoTicketSystem`.seat_categories ON (seats.seat_category_id = seat_categories.id)\nWHERE event_seats.event_id = ?;\n"

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
					WithArgs(utils.EqUUID(&eventSeats[0].EventID)).
					WillReturnRows(sqlmock.NewRows([]string{"event_seats.id", "event_seats.event_id", "event_seats.seat_id", "seats.id", "seats.seat_category_id", "seat_categories.id", "seat_categories.name"}).
						AddRow(eventSeats[0].ID, eventSeats[0].EventID, eventSeats[0].SeatID, eventSeats[0].Seat.ID, eventSeats[0].Seat.SeatCategoryID, eventSeats[0].Seat.SeatCategory.ID, eventSeats[0].Seat.SeatCategory.Name).
						AddRow(eventSeats[1].ID, eventSeats[1].EventID, eventSeats[1].SeatID, eventSeats[1].Seat.ID, eventSeats[1].Seat.SeatCategoryID, eventSeats[1].Seat.SeatCategory.ID, eventSeats[1].Seat.SeatCategory.Name),
					)

			},
			expectedSeats: &eventSeats,
			expectedError: nil,
		},
		{
			name: "Select event seats - no seats found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(utils.EqUUID(&eventSeats[0].EventID)).
					WillReturnRows(sqlmock.NewRows([]string{"event_seats.id", "event_seats.event_id", "event_seats.seat_id", "seats.id", "seats.seat_category_id", "seat_categories.id", "seat_categories.name"}))
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Select event seats - internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).
					WithArgs(utils.EqUUID(&eventSeats[0].EventID)).
					WillReturnError(sql.ErrConnDone)
			},
			expectedSeats: nil,
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

			// Create a new instance of the EventSeatRepository with the mock database connection
			eventSeatRepo := &EventSeatRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method on the repository instance
			seats, kts_err := eventSeatRepo.GetEventSeats(&eventSeats[0].EventID)

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
