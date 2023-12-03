package controllers

import (
	"log"

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
	log.Printf("GetActorById" + actorId.String())
	return ac.ActorRepo.GetActorById(actorId)
}
