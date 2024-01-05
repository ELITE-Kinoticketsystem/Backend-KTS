package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type TheatreControllerI interface {
	CreateTheatre(*models.CreateTheatreRequest) *models.KTSError
	GetTheatres() (*[]model.Theatres, *models.KTSError)
}

type TheatreController struct {
	TheatreRepo repositories.TheaterRepoI
}

func (tc *TheatreController) CreateTheatre(theatreData *models.CreateTheatreRequest) *models.KTSError {
	addressId := uuid.New()
	address := model.Addresses{
		ID:       &addressId,
		Street:   theatreData.Address.Street,
		StreetNr: theatreData.Address.StreetNr,
		Zipcode:  theatreData.Address.Zipcode,
		City:     theatreData.Address.City,
		Country:  theatreData.Address.Country,
	}

	err := tc.TheatreRepo.CreateAddress(address)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	theatreId := uuid.New()
	theatre := model.Theatres{
		ID:        &theatreId,
		Name:      theatreData.Name,
		LogoURL:   &theatreData.LogoUrl,
		AddressID: &addressId,
	}
	err = tc.TheatreRepo.CreateTheatre(theatre)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (tc *TheatreController) GetTheatres() (*[]model.Theatres, *models.KTSError){
	return tc.TheatreRepo.GetTheatres()
}
