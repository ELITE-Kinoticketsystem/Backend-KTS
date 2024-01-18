package controllers

import (
	"log"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type StatsControllerI interface {
	GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError)
	GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*models.StatsVisitsTwoArrays, *models.KTSError)
}

type StatsController struct {
	StatsRepo repositories.StatsRepositoryI
}

func (sc *StatsController) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
	return sc.StatsRepo.GetOrdersForStats()
}

func (sc *StatsController) GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*models.StatsVisitsTwoArrays, *models.KTSError) {
	// Calculate the number of days between start and end dates
	numDays := int(endTime.Sub(startTime).Hours() / 24)

	visitsTwoArrays := GenerateStatsBetweenDates(startTime, endTime, numDays)

	vists, err := sc.StatsRepo.GetTotalVisits(startTime, endTime, in)
	if err != nil {
		return nil, err
	}

	log.Print(vists)
	log.Print(visitsTwoArrays)

	for i := 0; i < numDays; i++ {
		log.Println(visitsTwoArrays.Date[i].GoString())

		for _, visit := range *vists {
			log.Printf("\t %v", visit.Date.GoString())
			if visitsTwoArrays.Date[i].Equal(visit.Date.Truncate(24 * time.Hour)) {
				log.Print(visit.Count)
				visitsTwoArrays.Count[i] = visit.Count
				log.Print(visitsTwoArrays.Count[i])
			}
		}

	}

	log.Println(visitsTwoArrays)

	return visitsTwoArrays, nil
}

// GenerateStatsBetweenDates generates an array of StatsVisitsTwoArrays between two given dates.
func GenerateStatsBetweenDates(startDate time.Time, endDate time.Time, numDays int) *models.StatsVisitsTwoArrays {
	var statsArray models.StatsVisitsTwoArrays

	// Initialize arrays with the number of days
	statsArray.Count = make([]int, numDays+1)
	statsArray.Date = make([]time.Time, numDays+1)

	// Populate the arrays with dates and initial count of 0
	for i := 0; i <= numDays; i++ {
		statsArray.Date[i] = startDate.Add(time.Duration(i) * 24 * time.Hour)
		statsArray.Count[i] = 0
	}

	return &statsArray
}
