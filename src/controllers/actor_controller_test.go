package controllers

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetActorByID(t *testing.T) {

	id := utils.NewUUID()

	testCases := []struct {
		name             string
		actorId          *uuid.UUID
		expectedActorDTO *models.ActorDTO
		expectedError    *models.KTSError
	}{
		{
			name:    "Get actor by id",
			actorId: id,
			expectedActorDTO: &models.ActorDTO{
				Actors: model.Actors{
					ID:          id,
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
						ID:          utils.NewUUID(),
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
	newActorId := utils.NewUUID()

	actor := &models.CreateActorDTO{
		Actors: model.Actors{
			ID:          utils.NewUUID(),
			Name:        "Test Actor",
			Description: "Test Description",
		},
		PicturesUrls: []string{
			"test1",
			"test2",
		},
	}

	testCases := []struct {
		name            string
		actorDto        *models.CreateActorDTO
		setExpectations func(mockRepo mocks.MockActorRepoI, actor *models.CreateActorDTO)
		expectedActorId *uuid.UUID
		expectedError   *models.KTSError
	}{
		{
			name:     "Create actor",
			actorDto: actor,
			setExpectations: func(mockRepo mocks.MockActorRepoI, actor *models.CreateActorDTO) {
				mockRepo.EXPECT().CreateActor(&actor.Actors).Return(newActorId, nil)
				mockRepo.EXPECT().CreateActorPicture(gomock.Any()).Return(nil, nil).AnyTimes()
			},
			expectedActorId: newActorId,
			expectedError:   nil,
		},
		{
			name:     "Create actor empty",
			actorDto: nil,
			setExpectations: func(mockRepo mocks.MockActorRepoI, actor *models.CreateActorDTO) {
				
			},
			expectedActorId: nil,
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:     "Create actor - failed",
			actorDto: actor,
			setExpectations: func(mockRepo mocks.MockActorRepoI, actor *models.CreateActorDTO) {
				mockRepo.EXPECT().CreateActor(&actor.Actors).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedActorId: nil,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:     "Create actor",
			actorDto: actor,
			setExpectations: func(mockRepo mocks.MockActorRepoI, actor *models.CreateActorDTO) {
				mockRepo.EXPECT().CreateActor(&actor.Actors).Return(newActorId, nil)
				mockRepo.EXPECT().CreateActorPicture(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedActorId: nil,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

			actorController := &ActorController{
				ActorRepo: mockActorRepo,
			}

			tc.setExpectations(*mockActorRepo, tc.actorDto)

			actorId, kts_err := actorController.CreateActor(tc.actorDto)

			if kts_err != tc.expectedError {
				t.Error("Expected error different than acutal")
			}
			assert.Equal(t, tc.expectedActorId, actorId, "Expected actor id to match")
		})
	}
}
