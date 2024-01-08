package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

// Validate Ticket will return a struct TicketDTOValidate

type TicketControllerI interface {
	GetTicketById(id *myid.UUID) (*models.TicketDTO, *models.KTSError)
	ValidateTicket(id *myid.UUID) *models.KTSError
}

type TicketController struct {
	TicketRepo repositories.TicketRepositoryI
}

func (tc *TicketController) GetTicketById(id *myid.UUID) (*models.TicketDTO, *models.KTSError) {
	return tc.TicketRepo.GetTicketById(id)
}

func (tc *TicketController) ValidateTicket(id *myid.UUID) *models.KTSError {
	return tc.TicketRepo.ValidateTicket(id)
}
