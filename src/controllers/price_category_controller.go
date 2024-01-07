package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type PriceCategoryControllerI interface {
	GetPriceCategories() (*[]model.PriceCategories, *models.KTSError)
	GetPriceCategoryById(id *myid.UUID) (*model.PriceCategories, *models.KTSError)
	CreatePriceCategory(priceCategory *model.PriceCategories) (*myid.UUID, *models.KTSError)
	UpdatePriceCategory(priceCategory *model.PriceCategories) (*myid.UUID, *models.KTSError)
	DeletePriceCategory(id *myid.UUID) *models.KTSError
}

type PriceCategoryController struct {
	PriceCategoryRepository repositories.PriceCategoryRepositoryI
}

func (pcc *PriceCategoryController) GetPriceCategoryById(priceCategoryID *myid.UUID) (*model.PriceCategories, *models.KTSError) {
	return pcc.PriceCategoryRepository.GetPriceCategoryById(priceCategoryID)
}

func (pcc *PriceCategoryController) GetPriceCategories() (*[]model.PriceCategories, *models.KTSError) {
	return pcc.PriceCategoryRepository.GetPriceCategories()
}

func (pcc *PriceCategoryController) CreatePriceCategory(priceCategory *model.PriceCategories) (*myid.UUID, *models.KTSError) {
	if priceCategory == nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	priceCategoryID, kts_err := pcc.PriceCategoryRepository.CreatePriceCategory(priceCategory)
	if kts_err != nil {
		return nil, kts_err
	}

	return priceCategoryID, nil
}

func (pcc *PriceCategoryController) UpdatePriceCategory(priceCategory *model.PriceCategories) (*myid.UUID, *models.KTSError) {
	if priceCategory == nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	priceCategoryID, kts_err := pcc.PriceCategoryRepository.UpdatePriceCategory(priceCategory)
	if kts_err != nil {
		return nil, kts_err
	}

	return priceCategoryID, nil
}

func (pcc *PriceCategoryController) DeletePriceCategory(id *myid.UUID) *models.KTSError {
	if id == nil {
		return kts_errors.KTS_BAD_REQUEST
	}

	kts_err := pcc.PriceCategoryRepository.DeletePriceCategory(id)
	if kts_err != nil {
		return kts_err
	}

	return nil
}
