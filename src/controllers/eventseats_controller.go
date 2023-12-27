package controllers

import (
	"slices"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type EventSeatControllerI interface {
	GetEventSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[][]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, *models.KTSError)
	BlockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError)
	AreUserSeatsNextToEachOther(eventId *uuid.UUID, userId *uuid.UUID, eventSeatId *uuid.UUID) (bool, *models.KTSError)
	UnblockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError)
	GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError)
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
			Type:          seat.Seat.Type,
			Price:         seat.EventSeatCategory.Price,
		}

		if seat.EventSeat.UserID == userId && !seat.EventSeat.Booked {
			currentUserSeats = append(currentUserSeats, currentSeat)
		}

		if len(currentUserSeats) > 0 {
			blockedUntil = seat.EventSeat.BlockedUntil
		}

		seatRow := seatRows[currentSeat.RowNr]
		seatRows[currentSeat.RowNr] = append(seatRow, currentSeat)
	}

	return seatMapToSlice(seatRows), &currentUserSeats, blockedUntil, nil
}

func (esc *EventSeatController) BlockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError) {
	currentTime := time.Now()
	blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

	if areUserSeatsNextToEachOther, err := esc.AreUserSeatsNextToEachOther(eventId, userId, eventSeatId); err != nil {
		return nil, err
	} else if !areUserSeatsNextToEachOther {
		return nil, kts_errors.KTS_CONFLICT
	}

	err := esc.EventSeatRepo.BlockEventSeatIfAvailable(eventId, eventSeatId, userId, &blockedUntil)

	if err != nil {
		return nil, err
	}

	err = esc.EventSeatRepo.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

	if err != nil {
		return nil, err
	}

	return &blockedUntil, nil
}

func (esc *EventSeatController) UnblockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError) {
	currentTime := time.Now()
	blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

	if areUserSeatsNextToEachOther, err := esc.AreUserSeatsNextToEachOtherWithoutSeat(eventId, userId, eventSeatId); err != nil {
		return nil, err
	} else if !areUserSeatsNextToEachOther {
		return nil, kts_errors.KTS_CONFLICT
	}

	err := esc.EventSeatRepo.UnblockEventSeat(eventId, eventSeatId, userId)

	if err != nil {
		return nil, err
	}

	err = esc.EventSeatRepo.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

	if err != nil {
		return nil, err
	}

	return &blockedUntil, nil
}

func (esc *EventSeatController) AreUserSeatsNextToEachOtherWithoutSeat(eventId *uuid.UUID, userId *uuid.UUID, eventSeatId *uuid.UUID) (bool, *models.KTSError) {
	seats, err := esc.EventSeatRepo.GetEventSeats(eventId)

	if err != nil {
		return false, err
	}

	var rowNr int32 = -1
	var columnNrs []int32
	var emtpySeatArray []models.GetEventSeatsDTO

	for _, seat := range *seats {
		if (seat.EventSeat.UserID == userId && !seat.EventSeat.Booked) && seat.EventSeat.ID != eventSeatId {
			if rowNr == -1 {
				rowNr = seat.Seat.RowNr
			} else if rowNr != seat.Seat.RowNr {
				return false, nil
			}

			columnNrs = append(columnNrs, seat.Seat.ColumnNr)
		}
		if seat.Seat.Type == string(utils.EMPTY) || seat.Seat.Type == string(utils.EMPTY_DOUBLE) {
			emtpySeatArray = append(emtpySeatArray, seat)
		}
	}

	if len(columnNrs) == 0 {
		return true, nil
	}

	if rowNr == -1 {
		return true, nil
	}

	slices.Sort[[]int32](columnNrs)

	for i := columnNrs[0] + 1; i < columnNrs[len(columnNrs)-1]; i++ {
		if slices.Contains(columnNrs, i) {
			continue
		}
		found := false
		for _, seat := range emtpySeatArray {
			if seat.Seat.ColumnNr == i && seat.Seat.RowNr == rowNr {
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}

	return true, nil
}

func (esc *EventSeatController) AreUserSeatsNextToEachOther(eventId *uuid.UUID, userId *uuid.UUID, eventSeatId *uuid.UUID) (bool, *models.KTSError) {
	seats, err := esc.EventSeatRepo.GetEventSeats(eventId)

	if err != nil {
		return false, err
	}

	var rowNr int32 = -1
	var columnNrs []int32
	var emtpySeatArray []models.GetEventSeatsDTO

	for _, seat := range *seats {
		if (seat.EventSeat.UserID == userId && !seat.EventSeat.Booked) || seat.EventSeat.ID == eventSeatId {
			if rowNr == -1 {
				rowNr = seat.Seat.RowNr
			} else if rowNr != seat.Seat.RowNr {
				return false, nil
			}

			columnNrs = append(columnNrs, seat.Seat.ColumnNr)
		}
		if seat.Seat.Type == string(utils.EMPTY) || seat.Seat.Type == string(utils.EMPTY_DOUBLE) {
			emtpySeatArray = append(emtpySeatArray, seat)
		}
	}

	if len(columnNrs) == 0 {
		return false, kts_errors.KTS_NOT_FOUND
	}

	if rowNr == -1 {
		return false, kts_errors.KTS_NOT_FOUND
	}

	slices.Sort[[]int32](columnNrs)

	for i := columnNrs[0] + 1; i < columnNrs[len(columnNrs)-1]; i++ {
		if slices.Contains(columnNrs, i) {
			continue
		}
		found := false
		for _, seat := range emtpySeatArray {
			if seat.Seat.ColumnNr == i && seat.Seat.RowNr == rowNr {
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}

	return true, nil
}

func (esc *EventSeatController) GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError) {
	return esc.EventSeatRepo.GetSelectedSeats(eventId, userId)
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
