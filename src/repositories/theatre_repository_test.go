package repositories

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
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
					utils.EqUUID(theatre.ID),
					theatre.Name,
					*theatre.LogoURL,
					utils.EqUUID(theatre.AddressID),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: samples.GetSampleTheatre(),
			setExpectations: func(mock sqlmock.Sqlmock, theatre *model.Theatres) {
				mock.ExpectExec("INSERT INTO `KinoTicketSystem`.theatres").WithArgs(
					utils.EqUUID(theatre.ID),
					theatre.Name,
					*theatre.LogoURL,
					utils.EqUUID(theatre.AddressID),
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
			// call CreateTheatre with theatre data
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

func TestCreateCinemaHall(t *testing.T) {
	testCases := []struct {
		name            string
		data            model.CinemaHalls
		setExpectations func(mock sqlmock.Sqlmock, cinemaHall *model.CinemaHalls)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			data: samples.GetSampleCinemaHall(),
			setExpectations: func(mock sqlmock.Sqlmock, cinemaHall *model.CinemaHalls) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.cinema_halls",
				).WithArgs(
					utils.EqUUID(cinemaHall.ID),
					cinemaHall.Name,
					cinemaHall.Capacity,
					utils.EqUUID(cinemaHall.TheatreID),
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: samples.GetSampleCinemaHall(),
			setExpectations: func(mock sqlmock.Sqlmock, cinemaHall *model.CinemaHalls) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.cinema_halls",
				).WithArgs(
					utils.EqUUID(cinemaHall.ID),
					cinemaHall.Name,
					cinemaHall.Capacity,
					utils.EqUUID(cinemaHall.TheatreID),
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
			// call CreateCinemaHall with cinema hall data
			kts_err := theatreRepo.CreateCinemaHall(tc.data)

			// THEN
			// check expected error
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreateSeat(t *testing.T) {
	testCases := []struct {
		name            string
		data            model.Seats
		setExpectations func(mock sqlmock.Sqlmock, seat *model.Seats)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			data: samples.GetSampleSeat(),
			setExpectations: func(mock sqlmock.Sqlmock, seat *model.Seats) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.seats",
				).WithArgs(
					utils.EqUUID(seat.ID),
					seat.RowNr,
					seat.ColumnNr,
					seat.VisibleRowNr,
					seat.VisibleColumnNr,
					utils.EqUUID(seat.SeatCategoryID),
					utils.EqUUID(seat.CinemaHallID),
					seat.Type,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: samples.GetSampleSeat(),
			setExpectations: func(mock sqlmock.Sqlmock, seat *model.Seats) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.seats",
				).WithArgs(
					utils.EqUUID(seat.ID),
					seat.RowNr,
					seat.ColumnNr,
					seat.VisibleRowNr,
					seat.VisibleColumnNr,
					utils.EqUUID(seat.SeatCategoryID),
					utils.EqUUID(seat.CinemaHallID),
					seat.Type,
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
			// call CreateSeat with seat data
			kts_err := theatreRepo.CreateSeat(tc.data)

			// THEN
			// check expected error
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetSeatCategories(t *testing.T) {
	sampleSeatCategories := samples.GetSampleSeatCategories()
	testCases := []struct {
		name                   string
		setExpectations        func(mock sqlmock.Sqlmock)
		expectedSeatCategories []model.SeatCategories
		expectedError          *models.KTSError
	}{
		{
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT .* FROM `KinoTicketSystem`.seat_categories;",
				).WillReturnRows(
					sqlmock.NewRows([]string{"seat_categories.id", "seat_categories.category_name"}).
						AddRow(sampleSeatCategories[0].ID, sampleSeatCategories[0].CategoryName).
						AddRow(sampleSeatCategories[1].ID, sampleSeatCategories[1].CategoryName).
						AddRow(sampleSeatCategories[2].ID, sampleSeatCategories[2].CategoryName),
				)
			},
			expectedSeatCategories: sampleSeatCategories,
			expectedError:          nil,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT .* FROM `KinoTicketSystem`.seat_categories;",
				).WillReturnError(sql.ErrConnDone)
			},
			expectedSeatCategories: nil,
			expectedError:          kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Not found",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT .* FROM `KinoTicketSystem`.seat_categories;",
				).WillReturnError(sql.ErrNoRows)
			},
			expectedSeatCategories: nil,
			expectedError:          kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
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

			seatCategories, kts_err := theatreRepo.GetSeatCategories()

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedSeatCategories, seatCategories)
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

func TestGetTheatres(t *testing.T) {
	theatres := samples.GetSampleTheatres()
	query := "SELECT .* FROM `KinoTicketSystem`.theatres"
	testCases := []struct {
		name             string
		data             []model.Theatres
		setExpectations  func(mock sqlmock.Sqlmock, theatres []model.Theatres)
		expectedTheatres *[]model.Theatres
		expectedError    *models.KTSError
	}{
		{
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock, theatres []model.Theatres) {
				mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"theatres.id", "theatres.name", "theatres.logo_url", "theatres.address_id"}).AddRow(theatres[0].ID, theatres[0].Name, theatres[0].LogoURL, theatres[0].AddressID).AddRow(theatres[1].ID, theatres[1].Name, theatres[1].LogoURL, theatres[1].AddressID))
			},
			expectedTheatres: &theatres,
			expectedError:    nil,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock, theatres []model.Theatres) {
				mock.ExpectQuery(query).WillReturnError(sql.ErrConnDone)
			},
			expectedTheatres: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
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
			tc.setExpectations(mock, theatres)

			// WHEN
			theatres, kts_err := theatreRepo.GetTheatres()

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedTheatres, theatres)
		})
	}

}
