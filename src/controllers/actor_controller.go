package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ActorControllerI interface {
	GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError)
}

type ActorController struct {
	ActorRepo repositories.ActorRepoI
}

func (ac *ActorController) GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError) {
	return ac.ActorRepo.GetActorById(actorId)
}
