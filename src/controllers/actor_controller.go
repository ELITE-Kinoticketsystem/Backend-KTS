package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ActorControllerI interface {
	GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError)
	GetActors() (*[]models.GetActorsDTO, *models.KTSError)
}

type ActorController struct {
	ActorRepo repositories.ActorRepoI
}

func (ac *ActorController) GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError) {
	return ac.ActorRepo.GetActorById(actorId)
}

func (ac *ActorController) GetActors() (*[]models.GetActorsDTO, *models.KTSError) {
	return ac.ActorRepo.GetActors()
}
