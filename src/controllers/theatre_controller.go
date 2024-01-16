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
	GetTheatres() (*[]models.GetTheatreWithAddress, *models.KTSError)
	CreateCinemaHall(*models.CreateCinemaHallRequest) *models.KTSError
	GetCinemaHallsForTheatre(*uuid.UUID) (*[]model.CinemaHalls, *models.KTSError)
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

	tx, err := tc.TheatreRepo.NewTransaction()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	defer tx.Rollback()

	kts_err := tc.TheatreRepo.CreateAddress(tx, address)
	if kts_err != nil {
		return kts_err
	}

	theatreId := uuid.New()
	theatre := model.Theatres{
		ID:        &theatreId,
		Name:      theatreData.Name,
		LogoURL:   &theatreData.LogoUrl,
		AddressID: &addressId,
	}
	kts_err = tc.TheatreRepo.CreateTheatre(tx, theatre)
	if kts_err != nil {
		return kts_err
	}

	if err = tx.Commit(); err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (tc *TheatreController) GetTheatres() (*[]models.GetTheatreWithAddress, *models.KTSError) {
	return tc.TheatreRepo.GetTheatres()
}

func (tc *TheatreController) CreateCinemaHall(cinemaHallData *models.CreateCinemaHallRequest) *models.KTSError {
	cinemaHallId := uuid.New()

	if !isHallValid(cinemaHallData) {
		return kts_errors.KTS_BAD_REQUEST
	}

	cinemaHall := model.CinemaHalls{
		ID:        &cinemaHallId,
		Name:      cinemaHallData.HallName,
		Width:     int32(cinemaHallData.Width),
		Height:    int32(cinemaHallData.Height),
		Capacity:  computeCapacity(cinemaHallData),
		TheatreID: cinemaHallData.TheatreId,
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

	row_nr := 1
	column_nr := 1
	for i, seat := range cinemaHallData.Seats {
		seatId := uuid.New()
		seatCategoryId, ok := seatCategoriesMap[seat.Category]
		if !ok {
			return kts_errors.KTS_BAD_REQUEST
		}
		seat := model.Seats{
			ID:             &seatId,
			RowNr:          int32(row_nr),
			ColumnNr:       int32(column_nr),
			X:              int32(seat.X),
			Y:              int32(seat.Y),
			SeatCategoryID: &seatCategoryId,
			CinemaHallID:   &cinemaHallId,
			Type:           seat.Type,
		}
		kts_err = tc.TheatreRepo.CreateSeat(seat)
		if kts_err != nil {
			return kts_err
		}
		if i > 0 && seat.Y != int32(cinemaHallData.Seats[i-1].Y) {
			row_nr++
			column_nr = 1
		} else {
			column_nr++
		}
	}
	return nil
}

func isHallValid(hall *models.CreateCinemaHallRequest) bool {
	// valid width and height
	if hall.Width <= 0 || hall.Height <= 0 {
		return false
	}

	// hall not empty
	if len(hall.Seats) == 0 {
		return false
	}

	currentRow := 0
	for i, seat := range hall.Seats {
		// valid coordinates
		if !(0 <= seat.X && seat.X < hall.Width && 0 <= seat.Y && seat.Y < hall.Height) {
			return false
		}

		// seats ascending inside row
		if i > 0 && seat.X < hall.Seats[i-1].X && seat.Y == hall.Seats[i-1].Y {
			return false
		}

		// rows in order
		if seat.Y < currentRow {
			return false
		} else if seat.Y > currentRow {
			currentRow = seat.Y
		}

		// double seats need to span two spaces
		if seat.Type == "double" {
			if i < len(hall.Seats)-1 && seat.Y == hall.Seats[i+1].Y && seat.X+1 == hall.Seats[i+1].X {
				return false
			}
			if seat.X == hall.Width-1 {
				return false
			}
		}
	}

	return true
}

func computeCapacity(hall *models.CreateCinemaHallRequest) int32 {
	capacity := int32(0)
	for _, seat := range hall.Seats {
		if seat.Type == "double" {
			capacity += 2
		} else {
			capacity += 1
		}
	}
	return capacity
}

func seatCategoriesToMap(seatCategories []model.SeatCategories) map[string]uuid.UUID {
	seatCategoriesMap := make(map[string]uuid.UUID)
	for _, seatCategory := range seatCategories {
		seatCategoriesMap[seatCategory.CategoryName] = *seatCategory.ID
	}
	return seatCategoriesMap
}

func (tc *TheatreController) GetCinemaHallsForTheatre(theatreId *uuid.UUID) (*[]model.CinemaHalls, *models.KTSError) {
	return tc.TheatreRepo.GetCinemaHallsForTheatre(theatreId)
}
