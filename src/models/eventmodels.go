package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type CreateEvtDTO struct {
	model.Events

	Movies []*uuid.UUID

	EventSeatCategories []model.EventSeatCategories
}
