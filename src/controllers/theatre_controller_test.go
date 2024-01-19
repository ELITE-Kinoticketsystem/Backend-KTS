package controllers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTheatre(t *testing.T) {
	sampleTheatreData := models.CreateTheatreRequest{
		Name:    "Theatre",
		LogoUrl: "LogoURL",
		Address: struct {
			Street   string
			StreetNr string
			Zipcode  string
			City     string
			Country  string
		}{
			Street:   "Street",
			StreetNr: "StreetNr",
			Zipcode:  "Zipcode",
			City:     "City",
			Country:  "Country",
		},
	}

	testCases := []struct {
		name            string
		theatreData     models.CreateTheatreRequest
		setExpectations func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock)
		expectedError   *models.KTSError
	}{
		{
			name:        "CreateTransaction internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock) {
				mockRepo.EXPECT().NewTransaction().Return(nil, sql.ErrTxDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "CreateAddress internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockRepo.EXPECT().CreateAddress(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs(samples.GetSampleAddress())).Return(kts_errors.KTS_INTERNAL_ERROR)
				dbMock.ExpectRollback()
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "CreateTheatre internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockRepo.EXPECT().CreateAddress(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs(samples.GetSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs((samples.GetSampleTheatre()))).Return(kts_errors.KTS_INTERNAL_ERROR)
				dbMock.ExpectRollback()
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "Commit internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockRepo.EXPECT().CreateAddress(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs(samples.GetSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs((samples.GetSampleTheatre()))).Return(nil)
				dbMock.ExpectCommit().WillReturnError(sql.ErrTxDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "Success",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockRepo.EXPECT().CreateAddress(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs(samples.GetSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(gomock.AssignableToTypeOf(&sql.Tx{}), utils.EqExceptUUIDs((samples.GetSampleTheatre()))).Return(nil)
				dbMock.ExpectCommit()
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock theatre repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			theatreRepoMock := mocks.NewMockTheaterRepoI(mockCtrl)
			theatreController := TheatreController{
				TheatreRepo: theatreRepoMock,
			}

			// create mock db manager
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}

			// create mock data
			theatreData := tc.theatreData

			// define expectations
			tc.setExpectations(*theatreRepoMock, theatreData, db, mock)

			// WHEN
			// call CreateTheatre with theatreData
			kts_err := theatreController.CreateTheatre(&theatreData)

			// THEN
			// check expected error
			assert.Equal(t, kts_err, tc.expectedError, "wrong error")
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestCreateCinemaHall(t *testing.T) {
	sampleRequest := samples.GetSampleCreateCinemaHallRequest()
	sampleCinemaHall := samples.GetSampleCinemaHall()
	sampleSeatCategories := samples.GetSampleSeatCategories()
	testCases := []struct {
		name            string
		cinemaHallData  models.CreateCinemaHallRequest
		setExpectations func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest)
		expectedError   *models.KTSError
	}{
		{
			name:            "Hall not rectangular",
			cinemaHallData:  samples.GetSampleCreateCinemaHallRequestNotRectangular(),
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:            "Invalid double seats",
			cinemaHallData:  samples.GetSampleCreateCinemaHallRequestInvalidDoubleSeats(),
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {},
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:           "Invalid seat category",
			cinemaHallData: samples.GetSampleCreateCinemaHallRequestInvalidSeatCategory(),
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {
				mockRepo.EXPECT().CreateCinemaHall(utils.EqExceptUUIDs(sampleCinemaHall)).Return(nil)
				mockRepo.EXPECT().GetSeatCategories().Return(sampleSeatCategories, nil)
			},
			expectedError: kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:           "Create hall internal error",
			cinemaHallData: sampleRequest,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {
				mockRepo.EXPECT().CreateCinemaHall(utils.EqExceptUUIDs(sampleCinemaHall)).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:           "Get seat categories internal error",
			cinemaHallData: sampleRequest,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {
				mockRepo.EXPECT().CreateCinemaHall(utils.EqExceptUUIDs(sampleCinemaHall)).Return(nil)
				mockRepo.EXPECT().GetSeatCategories().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:           "Create seat internal error",
			cinemaHallData: sampleRequest,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {
				mockRepo.EXPECT().CreateCinemaHall(utils.EqExceptUUIDs(sampleCinemaHall)).Return(nil)
				mockRepo.EXPECT().GetSeatCategories().Return(sampleSeatCategories, nil)
				mockRepo.EXPECT().CreateSeat(gomock.AssignableToTypeOf(model.Seats{})).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:           "Success",
			cinemaHallData: sampleRequest,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, cinemaHallData models.CreateCinemaHallRequest) {
				mockRepo.EXPECT().CreateCinemaHall(utils.EqExceptUUIDs(sampleCinemaHall)).Return(nil)
				mockRepo.EXPECT().GetSeatCategories().Return(sampleSeatCategories, nil)
				mockRepo.EXPECT().CreateSeat(gomock.AssignableToTypeOf(model.Seats{})).
					Times(len(sampleRequest.Seats) * len(sampleRequest.Seats[0])).Return(nil)
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock theatre repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			theatreRepoMock := mocks.NewMockTheaterRepoI(mockCtrl)
			theatreController := TheatreController{
				TheatreRepo: theatreRepoMock,
			}

			// create mock data
			cinemaHallData := tc.cinemaHallData

			// define expectations
			tc.setExpectations(*theatreRepoMock, cinemaHallData)

			// WHEN
			// call CreateCinemaHall with cinemaHallData
			err := theatreController.CreateCinemaHall(&cinemaHallData)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}

func TestGetCinemaHallsForTheatre(t *testing.T) {
	sampleCinemaHalls := samples.GetSampleCinemaHalls()
	theatreId := sampleCinemaHalls[0].TheatreID
	testCases := []struct {
		name            string
		theatreId       *uuid.UUID
		setExpectations func(mockRepo mocks.MockTheaterRepoI)
		expectedError   *models.KTSError
	}{
		{
			name:      "Internal error",
			theatreId: theatreId,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI) {
				mockRepo.EXPECT().GetCinemaHallsForTheatre(theatreId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:      "Success",
			theatreId: theatreId,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI) {
				mockRepo.EXPECT().GetCinemaHallsForTheatre(theatreId).Return(&sampleCinemaHalls, nil)
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock theatre repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			theatreRepoMock := mocks.NewMockTheaterRepoI(mockCtrl)
			theatreController := TheatreController{
				TheatreRepo: theatreRepoMock,
			}

			// create mock data
			theatreId := tc.theatreId

			// define expectations
			tc.setExpectations(*theatreRepoMock)

			// WHEN
			// call GetCinemaHallsForTheatre with theatreId
			_, err := theatreController.GetCinemaHallsForTheatre(theatreId)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}

func TestGetTheatres(t *testing.T) {
	testCases := []struct {
		name             string
		userId           *uuid.UUID
		setExpectations  func(mockRepo mocks.MockTheaterRepoI)
		expectedTheaters *[]models.GetTheatreWithAddress
		expectedError    *models.KTSError
	}{
		{
			name:   "Failed",
			userId: utils.NewUUID(),
			setExpectations: func(mockRepo mocks.MockTheaterRepoI) {
				mockRepo.EXPECT().GetTheatres().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedTheaters: nil,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			theatreRepo := mocks.NewMockTheaterRepoI(mockCtrl)
			theatreController := TheatreController{
				TheatreRepo: theatreRepo,
			}

			// define expectations
			tc.setExpectations(*theatreRepo)

			// WHEN
			theatres, kts_err := theatreController.GetTheatres()

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedTheaters, theatres)
		})
	}
}
