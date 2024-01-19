package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ActorControllerI interface {
	GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError)
	GetActors() (*[]models.GetActorsDTO, *models.KTSError)
	CreateActor(actor *models.CreateActorDTO) (*uuid.UUID, *models.KTSError)
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

func (ac *ActorController) CreateActor(actorDto *models.CreateActorDTO) (*uuid.UUID, *models.KTSError) {
	if actorDto == nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	actor := actorDto.Actors

	actorId, kts_err := ac.ActorRepo.CreateActor(&actor)
	if kts_err != nil {
		return nil, kts_err
	}

	imageUrls := actorDto.PicturesUrls

	for _, imageUrl := range imageUrls {
		_, kts_err := ac.ActorRepo.CreateActorPicture(&model.ActorPictures{
			ActorID: actorId,
			PicURL:  &imageUrl,
		})

		if kts_err != nil {
			return nil, kts_err
		}
	}

	return actorId, nil
}
