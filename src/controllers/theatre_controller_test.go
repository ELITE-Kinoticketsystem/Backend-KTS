package controllers

import (
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
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
		setExpectations func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest)
		expectedError   *models.KTSError
	}{
		{
			name:        "CreateAddress internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(samples.GetSampleAddress())).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "CreateTheatre internal error",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(samples.GetSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(utils.EqExceptId((samples.GetSampleTheatre()))).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:        "Success",
			theatreData: sampleTheatreData,
			setExpectations: func(mockRepo mocks.MockTheaterRepoI, theatreData models.CreateTheatreRequest) {
				mockRepo.EXPECT().CreateAddress(utils.EqExceptId(samples.GetSampleAddress())).Return(nil)
				mockRepo.EXPECT().CreateTheatre(utils.EqExceptId((samples.GetSampleTheatre()))).Return(nil)
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
