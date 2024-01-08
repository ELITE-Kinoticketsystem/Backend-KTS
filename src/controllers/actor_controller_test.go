package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

func TestGetActorByID(t *testing.T) {

	id := myid.NewUUID()

	testCases := []struct {
		name             string
		actorId          *myid.UUID
		expectedActorDTO *models.ActorDTO
		expectedError    *models.KTSError
	}{
		{
			name:    "Get actor by id",
			actorId: id,
			expectedActorDTO: &models.ActorDTO{
				Actors: model.Actors{
					ID:          *id,
					Name:        "Test Actor",
					Description: "Test Description",
				},
				Pictures: nil,
				Movies:   nil,
			},
			expectedError: nil,
		},
		{
			name:             "Get actor by id fails",
			actorId:          id,
			expectedActorDTO: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

			mockActorRepo.EXPECT().GetActorById(testCase.actorId).Return(testCase.expectedActorDTO, testCase.expectedError)

			ac := &ActorController{
				ActorRepo: mockActorRepo,
			}

			actorDTO, err := ac.GetActorById(testCase.actorId)

			if err != testCase.expectedError {
				t.Errorf("Expected error to be nil")
			}
			assert.Equal(t, testCase.expectedActorDTO, actorDTO, "Expected actor DTO to match")
		})
	}
}

func TestGetActors(t *testing.T) {

	testCases := []struct {
		name              string
		expectedActorsDTO *[]models.GetActorsDTO
		expectedError     *models.KTSError
	}{
		{
			name: "Get actors",
			expectedActorsDTO: &[]models.GetActorsDTO{
				{
					Actors: model.Actors{
						ID:          *myid.NewUUID(),
						Name:        "Test Actor",
						Description: "Test Description",
					},
					Pictures: nil,
				},
			},
			expectedError: nil,
		},
		{
			name:              "Get actors fails",
			expectedActorsDTO: nil,
			expectedError:     kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

			mockActorRepo.EXPECT().GetActors().Return(testCase.expectedActorsDTO, testCase.expectedError)

			ac := &ActorController{
				ActorRepo: mockActorRepo,
			}

			actorsDTO, err := ac.GetActors()

			if err != testCase.expectedError {
				t.Errorf("Expected error to be nil")
			}
			assert.Equal(t, testCase.expectedActorsDTO, actorsDTO, "Expected actors DTO to match")
		})
	}
}

func TestCreateActor(t *testing.T) {

	id := myid.NewUUID()

	testCases := []struct {
		name          string
		actorDto      *models.CreateActorDTO
		expectedActor *myid.UUID
		expectedError *models.KTSError
	}{
		{
			name: "Create actor",
			actorDto: &models.CreateActorDTO{
				Actors: model.Actors{
					ID:          *id,
					Name:        "Test Actor",
					Description: "Test Description",
				},
				PicturesUrls: []string{
					"test1",
					"test2",
				},
			},
			expectedActor: id,
			expectedError: nil,
		},
		{
			name:          "Create actor fails",
			actorDto:      nil,
			expectedActor: nil,
			expectedError: kts_errors.KTS_BAD_REQUEST,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

			if testCase.actorDto != nil {
				mockActorRepo.EXPECT().CreateActor(gomock.Any()).Return(testCase.expectedActor, testCase.expectedError)

				if testCase.actorDto.PicturesUrls != nil {
					mockActorRepo.EXPECT().CreateActorPicture(gomock.Any()).Return(id, testCase.expectedError).AnyTimes()
				}
			}

			ac := &ActorController{
				ActorRepo: mockActorRepo,
			}

			actorId, kts_err := ac.CreateActor(testCase.actorDto)

			if kts_err != testCase.expectedError {
				t.Error("Expected error different than acutal")
			}
			assert.Equal(t, testCase.expectedActor, actorId, "Expected actor id to match")
		})
	}
}
