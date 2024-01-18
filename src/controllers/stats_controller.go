package controllers

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type StatsControllerI interface {
	GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError)
	GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*[]models.StatsStruct, *models.KTSError)
}

type StatsController struct {
	StatsRepo repositories.StatsRepositoryI
}

func (sc *StatsController) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
	return sc.StatsRepo.GetOrdersForStats()
}

func (sc *StatsController) GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*[]models.StatsStruct, *models.KTSError) {
	return sc.StatsRepo.GetTotalVisits(startTime, endTime, in)
}
