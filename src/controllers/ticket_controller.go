package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

// Validate Ticket will return a struct TicketDTOValidate

type TicketControllerI interface {
	GetTicketById(id *uuid.UUID) (*models.TicketDTO, *models.KTSError)
	ValidateTicket(id *uuid.UUID) *models.KTSError
}

type TicketController struct {
	TicketRepo repositories.TicketRepositoryI
}

func (tc *TicketController) GetTicketById(id *uuid.UUID) (*models.TicketDTO, *models.KTSError) {
	return tc.TicketRepo.GetTicketById(id)
}

func (tc *TicketController) ValidateTicket(id *uuid.UUID) *models.KTSError {
	return tc.TicketRepo.ValidateTicket(id)
}
