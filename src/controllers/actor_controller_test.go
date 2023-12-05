package controllers

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func TestGetActorById(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// GIVEN
	actorId := uuid.New()
	expectedActorDTO := &models.ActorDTO{
		Actors: model.Actors{
			ID:          actorId,
			Name:        "Test Actor",
			Description: "Test Description",
		},
		Pictures: nil,
		Movies:   nil,
	}

	mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

	mockActorRepo.EXPECT().GetActorById(&actorId).Return(expectedActorDTO, nil)

	ac := &ActorController{
		ActorRepo: mockActorRepo,
	}

	// WHEN
	actorDTO, err := ac.GetActorById(&actorId)
	fmt.Println(err)

	// THEN
	if err != nil {
		t.Errorf("Expected error to be nil")
	}
	assert.Equal(t, expectedActorDTO, actorDTO, "Expected actor DTO to match")
}

func TestGetActorByIdFails(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// GIVEN
	actorId := uuid.New()

	mockActorRepo := mocks.NewMockActorRepoI(mockCtrl)

	mockActorRepo.EXPECT().GetActorById(&actorId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)

	ac := &ActorController{
		ActorRepo: mockActorRepo,
	}

	// WHEN
	actorDTO, err := ac.GetActorById(&actorId)

	// THEN
	if err == nil {
		t.Errorf("Expected error to be not nil")
	}
	assert.Equal(t, (*models.ActorDTO)(nil), actorDTO, "Expected actor DTO to be nil")
}
