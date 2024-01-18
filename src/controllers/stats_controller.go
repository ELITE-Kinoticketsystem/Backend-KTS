package controllers

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type StatsControllerI interface {
	GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError)
	GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*models.StatsVisitsTwoArrays, *models.KTSError)
	GetTotalVisitsForTheatre(startTime time.Time, endTime time.Time, in string, theatreName string) (*models.StatsVisitsTwoArrays, *models.KTSError)
}

type StatsController struct {
	StatsRepo repositories.StatsRepositoryI
}

func (sc *StatsController) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
	return sc.StatsRepo.GetOrdersForStats()
}

func (sc *StatsController) GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*models.StatsVisitsTwoArrays, *models.KTSError) {

	visitsTwoArrays := GenerateStatsArray(startTime, endTime, in)
	visitsTwoArrays := GenerateStatsArray(startTime, endTime, in)

	vists, err := sc.StatsRepo.GetTotalVisits(startTime, endTime, in)
	if err != nil {
		return nil, err
	}

	// Loop through all dates
	for i := 0; i < len(visitsTwoArrays.Count); i++ {

		// Loop through all visits
		for _, visit := range *vists {

			// Depending on what the user wants to filter by, check if the dates are equal
			switch in {
			case "day":
				if visitsTwoArrays.Date[i].Equal(visit.Date.Truncate(24 * time.Hour)) {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			case "month":
				if visitsTwoArrays.Date[i].Month() == visit.Date.Month() &&
					visitsTwoArrays.Date[i].Year() == visit.Date.Year() {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			case "year":
				if visitsTwoArrays.Date[i].Year() == visit.Date.Year() {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			}
		}

	}

	return visitsTwoArrays, nil
}

func GenerateStatsArray(startDate, endDate time.Time, filterBy string) *models.StatsVisitsTwoArrays {
	var statsArray models.StatsVisitsTwoArrays

	// Initialize the start and end date based on the filter
	switch filterBy {
	case "day":
		startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
		endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, time.UTC)
	case "month":
		startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		endDate = time.Date(endDate.Year(), endDate.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	case "year":
		startDate = time.Date(startDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		endDate = time.Date(endDate.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	// Initialize the array with zero values
	for currentDate := startDate; currentDate.Before(endDate); {
		statsArray.Count = append(statsArray.Count, 0)
		statsArray.Date = append(statsArray.Date, currentDate)
		statsArray.Revenue = append(statsArray.Revenue, 0)

		// Increment the current date based on the filter
		switch filterBy {
		case "day":
			currentDate = currentDate.AddDate(0, 0, 1)
		case "month":
			currentDate = currentDate.AddDate(0, 1, 0)
		case "year":
			currentDate = currentDate.AddDate(1, 0, 0)
		}
	}

	return &statsArray
}

func (sc *StatsController) GetTotalVisitsForTheatre(startTime time.Time, endTime time.Time, in string, theatreName string) (*models.StatsVisitsTwoArrays, *models.KTSError) {

	visitsTwoArrays := GenerateStatsArray(startTime, endTime, in)

	vists, err := sc.StatsRepo.GetTotalVisitsForTheatre(startTime, endTime, in, theatreName)
	if err != nil {
		return nil, err
	}

	// Loop through all dates
	for i := 0; i < len(visitsTwoArrays.Count); i++ {

		// Loop through all visits
		for _, visit := range *vists {

			// Depending on what the user wants to filter by, check if the dates are equal
			switch in {
			case "day":
				if visitsTwoArrays.Date[i].Equal(visit.Date.Truncate(24 * time.Hour)) {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			case "month":
				if visitsTwoArrays.Date[i].Month() == visit.Date.Month() &&
					visitsTwoArrays.Date[i].Year() == visit.Date.Year() {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			case "year":
				if visitsTwoArrays.Date[i].Year() == visit.Date.Year() {
					visitsTwoArrays.Count[i] = visit.Count
					visitsTwoArrays.Revenue[i] = visit.Revenue
				}
			}
		}

	}

	return visitsTwoArrays, nil
}
