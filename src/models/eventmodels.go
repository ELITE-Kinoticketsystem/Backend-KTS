package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

type CreateEvtDTO struct {
	model.Events

	Movies []*myid.UUID

	EventSeatCategories []model.EventSeatCategories
}

type GetSpecialEventsDTO struct {
	model.Events

	Movies []*model.Movies
}
