package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTheatre(t *testing.T) {
	sampleTheatreData := models.CreateTheatreRequest{
		Name: "Theatre",
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
		setExpectations func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest)
		expectedError   *models.KTSError
	}{
		{
			name:        "CreateAddress internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(getSampleAddress())).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "CreateTheatre internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(getSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(utils.EqExceptId((getSampleTheatre()))).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "Success",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(getSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(utils.EqExceptId((getSampleTheatre()))).Return(nil)
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
				theatreRepo: theatreRepoMock,
			}

			// create mock data
			theatreData := tc.theatreData

			// define expectations
			tc.setExpectations(*theatreRepoMock, theatreData)

			// WHEN
			// call CreateTheatre with theatreData
			err := theatreController.CreateTheatre(&theatreData)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}

}

func getSampleAddress() model.Addresses {
	return model.Addresses{
		/* Id */
		Street:   "Street",
		StreetNr: "StreetNr",
		Zipcode:  "Zipcode",
		City:     "City",
		Country:  "Country",
	}
}

func getSampleTheatre() model.Theatres {
	return model.Theatres{
		/* ID */
		Name: "Theatre",
		/* AddressID */
	}
}
func getSampleTheatreData() models.CreateTheatreRequest {
	return models.CreateTheatreRequest{
		Name: "Theatre",
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
}
