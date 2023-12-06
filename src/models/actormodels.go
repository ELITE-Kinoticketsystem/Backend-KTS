package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
)

type ActorDTO struct {
	model.Actors

	Pictures []struct {
		model.ActorPictures
	}

	Movies []struct {
		model.Movies
	}
}

type GetActorsDTO struct {
	model.Actors

	Pictures []struct {
		model.ActorPictures
	}
}
