package controllers

import (
	"slices"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type EventSeatControllerI interface {
	GetEventSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[][]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, *models.KTSError)
}

type EventSeatController struct {
	EventSeatRepo repositories.EventSeatRepoI
}

func (esc *EventSeatController) GetEventSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[][]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, *models.KTSError) {
	seats, kts_err := esc.EventSeatRepo.GetEventSeats(eventId)

	if kts_err != nil {
		return nil, nil, nil, kts_err
	}

	seatRows := make(map[int32][]models.GetSeatsForSeatSelectorDTO)
	currentUserSeats := []models.GetSeatsForSeatSelectorDTO{}
	var blockedUntil *time.Time

	for _, seat := range *seats {
		currentSeat := models.GetSeatsForSeatSelectorDTO{
			ID:            seat.EventSeat.ID,
			RowNr:         seat.Seat.RowNr,
			ColumnNr:      seat.Seat.ColumnNr,
			Available:     (seat.EventSeat.BlockedUntil == nil || seat.EventSeat.BlockedUntil.Before(time.Now()) || seat.EventSeat.UserID == nil) && !seat.EventSeat.Booked,
			BookedByOther: (seat.EventSeat.BlockedUntil != nil && (seat.EventSeat.BlockedUntil.After(time.Now()) && (seat.EventSeat.UserID != nil && seat.EventSeat.UserID != userId))) && !seat.EventSeat.Booked,
			Category:      seat.SeatCategory.CategoryName,
			Price:         seat.EventSeatCategory.Price,
		}

		if seat.EventSeat.UserID == userId && !seat.EventSeat.Booked {
			blockedUntil = seat.EventSeat.BlockedUntil
			currentUserSeats = append(currentUserSeats, currentSeat)
		}

		seatRow := seatRows[currentSeat.RowNr]
		seatRows[currentSeat.RowNr] = append(seatRow, currentSeat)
	}

	return seatMapToSlice(seatRows), &currentUserSeats, blockedUntil, nil
}

func seatMapToSlice(seatMap map[int32][]models.GetSeatsForSeatSelectorDTO) *[][]models.GetSeatsForSeatSelectorDTO {
	seatSlice := [][]models.GetSeatsForSeatSelectorDTO{}

	for _, seatRow := range seatMap {
		// sort seatrow by columnNr
		slices.SortFunc(seatRow, func(a, b models.GetSeatsForSeatSelectorDTO) int {
			return (int)(a.ColumnNr - b.ColumnNr)
		})

		seatSlice = append(seatSlice, seatRow)
	}

	return &seatSlice
}
