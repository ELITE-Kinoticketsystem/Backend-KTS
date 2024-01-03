package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
)

type ActorDTO struct {
	model.Actors

	Pictures []model.ActorPictures

	Movies []model.Movies
}

type GetActorsDTO struct {
	model.Actors

	Pictures []model.ActorPictures
}

type CreateActorDTO struct {
	model.Actors

	PicturesUrls []string
}
