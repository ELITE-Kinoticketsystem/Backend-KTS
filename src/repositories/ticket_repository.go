package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

type TicketRepoI interface {
	CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError)
}

type TicketRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TicketRepository) CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError) {
	return nil, nil
}
