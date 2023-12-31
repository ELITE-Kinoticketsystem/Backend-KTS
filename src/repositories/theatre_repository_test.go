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
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTheatre(t *testing.T) {
	testCases := []struct {
		name            string
		data            model.Theatres
		setExpectations func(mock sqlmock.Sqlmock, theatre *model.Theatres)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			data: samples.GetSampleTheatre(),
			setExpectations: func(mock sqlmock.Sqlmock, theatre *model.Theatres) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.theatres").WithArgs(
					theatre.ID, theatre.Name, theatre.AddressID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: samples.GetSampleTheatre(),
			setExpectations: func(mock sqlmock.Sqlmock, theatre *model.Theatres) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.theatres").WithArgs(
					theatre.ID, theatre.Name, theatre.AddressID,
				).WillReturnError(sql.ErrConnDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			theatreRepo := TheatreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, &tc.data)

			// WHEN
			// call CreateTheatre with user data
			kts_err := theatreRepo.CreateTheatre(tc.data)

			// THEN
			// check expected error
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

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
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.seats .*").WithArgs(sqlmock.AnyArg()).WillReturnRows(
					sqlmock.NewRows([]string{"seats.id", "seats.cinema_hall_id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id"}).
						AddRow(expectedSeats[0].ID, expectedSeats[0].CinemaHallID, expectedSeats[0].RowNr, expectedSeats[0].ColumnNr, expectedSeats[0].SeatCategoryID).
						AddRow(expectedSeats[1].ID, expectedSeats[1].CinemaHallID, expectedSeats[1].RowNr, expectedSeats[1].ColumnNr, expectedSeats[1].SeatCategoryID),
				)
			},
			expectedSeats: expectedSeats,
			expectedError: nil,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.seats .*").WithArgs(sqlmock.AnyArg()).WillReturnError(sql.ErrConnDone)
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Not found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .* FROM `KinoTicketSystem`.seats .*").WithArgs(sqlmock.AnyArg()).WillReturnError(sql.ErrNoRows)
			},
			expectedSeats: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
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
