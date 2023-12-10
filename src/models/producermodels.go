package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
)

type ProducerDTO struct {
	model.Producers

	Pictures []model.ProducerPictures

	Movies []model.Movies
}

type GetProducersDTO struct {
	model.Producers

	Pictures []model.ProducerPictures
}

type CreateProducerDTO struct {
	model.Producers

	PicturesUrls []string
}
