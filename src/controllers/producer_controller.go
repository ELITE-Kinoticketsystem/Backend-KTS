package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type ProducerControllerI interface {
	GetProducers() (*[]models.GetProducersDTO, *models.KTSError)
	GetProducerById(producerId *uuid.UUID) (*models.ProducerDTO, *models.KTSError)
	CreateProducer(producerDto *models.CreateProducerDTO) (*uuid.UUID, *models.KTSError)
}

type ProducerController struct {
	ProducerRepo repositories.ProducerRepositoryI
}

func (pc *ProducerController) GetProducers() (*[]models.GetProducersDTO, *models.KTSError) {
	return pc.ProducerRepo.GetProducers()
}

func (pc *ProducerController) GetProducerById(producerId *uuid.UUID) (*models.ProducerDTO, *models.KTSError) {
	return pc.ProducerRepo.GetProducerById(producerId)
}

func (pc *ProducerController) CreateProducer(producerDto *models.CreateProducerDTO) (*uuid.UUID, *models.KTSError) {
	if producerDto.Producers == (model.Producers{}) {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	producer := producerDto.Producers

	producerId, kts_err := pc.ProducerRepo.CreateProducer(&producer)

	if kts_err != nil {
		return nil, kts_err
	}

	imageUrls := producerDto.PicturesUrls

	for _, imageUrl := range imageUrls {
		_, kts_err := pc.ProducerRepo.CreateProducerPicture(&model.ProducerPictures{
			ProducerID: producerId,
			PicURL:     &imageUrl,
		})

		if kts_err != nil {
			return nil, kts_err
		}
	}

	return producerId, nil

}
