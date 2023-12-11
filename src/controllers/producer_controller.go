package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ProducerControllerI interface {
	GetProducers() (*[]model.Producers, *models.KTSError)
	GetProducerById(id *uuid.UUID) (*model.Producers, *models.KTSError)
	CreateProducer(producer *model.Producers) *models.KTSError
	UpdateProducer(producer *model.Producers) *models.KTSError
	DeleteProducer(id *uuid.UUID) *models.KTSError
}

type ProducerController struct {
	ProducerRepo repositories.ProducerRepositoryI
}
