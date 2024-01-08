package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type TheatreControllerI interface {
	CreateTheatre(*models.CreateTheatreRequest) *models.KTSError
	GetTheatres() (*[]model.Theatres, *models.KTSError)
	CreateCinemaHall(*models.CreateCinemaHallRequest) *models.KTSError
	CreateCinemaHallFast(*models.CreateCinemaHallRequest) *models.KTSError
	GetCinemaHallsForTheatre(*myid.UUID) (*[]model.CinemaHalls, *models.KTSError)
}

type TheatreController struct {
	TheatreRepo repositories.TheaterRepoI
}

func (tc *TheatreController) CreateTheatre(theatreData *models.CreateTheatreRequest) *models.KTSError {
	addressId := myid.New()
	address := model.Addresses{
		ID:       addressId,
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

	theatreId := myid.New()
	theatre := model.Theatres{
		ID:        theatreId,
		Name:      theatreData.Name,
		LogoURL:   &theatreData.LogoUrl,
		AddressID: addressId,
	}
	err = tc.TheatreRepo.CreateTheatre(theatre)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (tc *TheatreController) GetTheatres() (*[]model.Theatres, *models.KTSError) {
	return tc.TheatreRepo.GetTheatres()
}

func (tc *TheatreController) CreateCinemaHall(cinemaHallData *models.CreateCinemaHallRequest) *models.KTSError {
	cinemaHallId := myid.New()

	if !isHallValid(cinemaHallData) {
		return kts_errors.KTS_BAD_REQUEST
	}

	cinemaHall := model.CinemaHalls{
		ID:        cinemaHallId,
		Name:      cinemaHallData.HallName,
		Capacity:  computeCapacity(cinemaHallData),
		TheatreID: *cinemaHallData.TheatreId,
	}

	kts_err := tc.TheatreRepo.CreateCinemaHall(cinemaHall)
	if kts_err != nil {
		return kts_err
	}

	seatCategories, kts_err := tc.TheatreRepo.GetSeatCategories()
	if kts_err != nil {
		return kts_err
	}

	seatCategoriesMap := seatCategoriesToMap(seatCategories)

	visible_row := 1
	for _, row := range cinemaHallData.Seats {
		visible_column := 1
		emtpy_seats := 0
		for _, seat := range row {
			if seat.Type == "empty" {
				emtpy_seats++
				visible_column--
			}
			seatId := myid.New()
			seatCategoryId, ok := seatCategoriesMap[seat.Category]
			if !ok {
				return kts_errors.KTS_BAD_REQUEST
			}
			seat := model.Seats{
				ID:              seatId,
				RowNr:           int32(seat.RowNr),
				ColumnNr:        int32(seat.ColumnNr),
				VisibleRowNr:    int32(visible_row),
				VisibleColumnNr: int32(visible_column),
				SeatCategoryID:  seatCategoryId,
				CinemaHallID:    cinemaHallId,
				Type:            seat.Type,
			}
			kts_err = tc.TheatreRepo.CreateSeat(seat)
			if kts_err != nil {
				return kts_err
			}
			visible_column++
		}
		if emtpy_seats == len(row) {
			continue
		}
		visible_row++
	}

	return nil
}

func (tc *TheatreController) CreateCinemaHallFast(cinemaHallData *models.CreateCinemaHallRequest) *models.KTSError {
	cinemaHallId := myid.New()

	if !isHallValid(cinemaHallData) {
		return kts_errors.KTS_BAD_REQUEST
	}

	cinemaHall := model.CinemaHalls{
		ID:        cinemaHallId,
		Name:      cinemaHallData.HallName,
		Capacity:  computeCapacity(cinemaHallData),
		TheatreID: *cinemaHallData.TheatreId,
	}

	kts_err := tc.TheatreRepo.CreateCinemaHall(cinemaHall)
	if kts_err != nil {
		return kts_err
	}

	seatCategories, kts_err := tc.TheatreRepo.GetSeatCategories()
	if kts_err != nil {
		return kts_err
	}

	seatCategoriesMap := seatCategoriesToMap(seatCategories)

	var seats []model.Seats
	visible_row := 1
	for _, row := range cinemaHallData.Seats {
		visible_column := 1
		emtpy_seats := 0
		for _, seat := range row {
			if seat.Type == "empty" {
				emtpy_seats++
				visible_column--
			}
			seatId := myid.New()
			seatCategoryId, ok := seatCategoriesMap[seat.Category]
			if !ok {
				return kts_errors.KTS_BAD_REQUEST
			}
			seat := model.Seats{
				ID:              seatId,
				RowNr:           int32(seat.RowNr),
				ColumnNr:        int32(seat.ColumnNr),
				VisibleRowNr:    int32(visible_row),
				VisibleColumnNr: int32(visible_column),
				SeatCategoryID:  seatCategoryId,
				CinemaHallID:    cinemaHallId,
				Type:            seat.Type,
			}
			seats = append(seats, seat)
			visible_column++
		}
		if emtpy_seats == len(row) {
			continue
		}
		visible_row++
	}
			kts_err = tc.TheatreRepo.CreateSeats(seats)
			if kts_err != nil {
				return kts_err
			}

	return nil
}

func isHallValid(hall *models.CreateCinemaHallRequest) bool {
	for i, row := range hall.Seats {

		// check if hall is rectangular
		if i > 0 && len(row) != len(hall.Seats[i-1]) {
			return false
		}

		// check for valid double seats
		for j, seat := range row {
			if j < len(row)-1 && seat.Type == "double" && row[j+1].Type != "emptyDouble" {
				return false
			}
		}
	}
	return true
}

func computeCapacity(hall *models.CreateCinemaHallRequest) int32 {
	capacity := int32(0)
	for _, row := range hall.Seats {
		for _, seat := range row {
			if seat.Type != "empty" {
				capacity++
			}
		}
	}
	return capacity
}

func seatCategoriesToMap(seatCategories []model.SeatCategories) map[string]myid.UUID {
	seatCategoriesMap := make(map[string]myid.UUID)
	for _, seatCategory := range seatCategories {
		seatCategoriesMap[seatCategory.CategoryName] = seatCategory.ID
	}
	return seatCategoriesMap
}

func (tc *TheatreController) GetCinemaHallsForTheatre(theatreId *myid.UUID) (*[]model.CinemaHalls, *models.KTSError) {
	return tc.TheatreRepo.GetCinemaHallsForTheatre(theatreId)
}
