package repositories

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetSeatsForCinemaHall(t *testing.T) {
	cinemaHallID := utils.NewUUID()

	expectedSeats := []model.Seats{
		{
			ID:             utils.NewUUID(),
			CinemaHallID:   cinemaHallID,
			ColumnNr:       1,
			RowNr:          1,
			SeatCategoryID: utils.NewUUID(),
		},
		{
			ID:             utils.NewUUID(),
			CinemaHallID:   cinemaHallID,
			ColumnNr:       2,
			RowNr:          1,
			SeatCategoryID: utils.NewUUID(),
		},
	}

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedSeats   []model.Seats
		expectedError   *models.KTSError
	}{
		{
			name: "Get seats for cinema hall",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.seats .*").WillReturnRows(
					sqlmock.NewRows([]string{"seats.id", "seats.cinema_hall_id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id"}).
						AddRow(expectedSeats[0].ID, expectedSeats[0].CinemaHallID, expectedSeats[0].RowNr, expectedSeats[0].ColumnNr, expectedSeats[0].SeatCategoryID).
						AddRow(expectedSeats[1].ID, expectedSeats[1].CinemaHallID, expectedSeats[1].RowNr, expectedSeats[1].ColumnNr, expectedSeats[1].SeatCategoryID),
				)
			},
			expectedSeats: expectedSeats,
			expectedError: nil,
		},
		{
			name: "Get seats for cinema hall sql error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.seats .*").WillReturnError(sql.ErrConnDone)
			},
			expectedSeats: nil,
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

			theatreRepo := &TheatreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			seats, ktsErr := theatreRepo.GetSeatsForCinemaHall(cinemaHallID)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if !reflect.DeepEqual(seats, tc.expectedSeats) {
				t.Errorf("Expected seats: %v, got: %v", tc.expectedSeats, seats)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
