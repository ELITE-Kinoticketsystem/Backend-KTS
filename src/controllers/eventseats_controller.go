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
	GetEventSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, int32, int32, *models.KTSError)
	BlockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError)
	AreUserSeatsNextToEachOther(eventId *uuid.UUID, userId *uuid.UUID, eventSeatId *uuid.UUID) (bool, *models.KTSError)
	UnblockEventSeat(eventId *uuid.UUID, eventSeatId *uuid.UUID, userId *uuid.UUID) (*time.Time, *models.KTSError)
	UnblockAllEventSeats(eventId *uuid.UUID, userId *uuid.UUID) *models.KTSError
	GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSlectedSeatsDTO, *models.KTSError)
}

type EventSeatController struct {
	EventSeatRepo repositories.EventSeatRepoI
}

func (esc *EventSeatController) GetEventSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, int32, int32, *models.KTSError) {
	seats, kts_err := esc.EventSeatRepo.GetEventSeats(eventId)
	if kts_err != nil {
		return nil, nil, nil, 0, 0, kts_err
	}

	width, height, kts_err := esc.EventSeatRepo.GetHallDimensions(eventId)
	if kts_err != nil {
		return nil, nil, nil, 0, 0, kts_err
	}

	currentUserSeats := []models.GetSeatsForSeatSelectorDTO{}
	var blockedUntil *time.Time

	event_seats := []models.GetSeatsForSeatSelectorDTO{}

	for _, seat := range *seats {
		currentSeat := models.GetSeatsForSeatSelectorDTO{
			ID:             seat.EventSeat.ID,
			RowNr:          seat.Seat.Y,
			ColumnNr:       seat.Seat.X,
			Available:      (seat.EventSeat.BlockedUntil == nil || seat.EventSeat.BlockedUntil.Before(time.Now()) || seat.EventSeat.UserID == nil) && !seat.EventSeat.Booked,
			BlockedByOther: (seat.EventSeat.BlockedUntil != nil && (seat.EventSeat.BlockedUntil.After(time.Now()) && (seat.EventSeat.UserID != nil && *seat.EventSeat.UserID != *userId))) || seat.EventSeat.Booked,
			Category:       seat.SeatCategory.CategoryName,
			Type:           seat.Seat.Type,
			Price:          seat.EventSeatCategory.Price,
		}

		if seat.EventSeat.UserID != nil && *seat.EventSeat.UserID == *userId && !seat.EventSeat.Booked && seat.EventSeat.BlockedUntil != nil && seat.EventSeat.BlockedUntil.After(time.Now()) {
			currentUserSeats = append(currentUserSeats, currentSeat)
			if len(currentUserSeats) == 1 {
				blockedUntil = seat.EventSeat.BlockedUntil
			}
		}
		event_seats = append(event_seats, currentSeat)
	}

	return &event_seats, &currentUserSeats, blockedUntil, width, height, nil
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

	_, err = esc.EventSeatRepo.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

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

	affectedRows, err := esc.EventSeatRepo.UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, &blockedUntil)

	if err != nil {
		return nil, err
	}
	if affectedRows == 0 {
		return nil, nil
	}

	return &blockedUntil, nil
}

func (esc *EventSeatController) UnblockAllEventSeats(eventId *uuid.UUID, userId *uuid.UUID) *models.KTSError {
	return esc.EventSeatRepo.UnblockAllEventSeats(eventId, userId)
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
		if ((seat.EventSeat.UserID != nil && *seat.EventSeat.UserID == *userId) && (seat.EventSeat.BlockedUntil != nil && seat.EventSeat.BlockedUntil.After(time.Now())) && !seat.EventSeat.Booked) && *seat.EventSeat.ID != *eventSeatId {
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

	// @Collin - I think this is not needed
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
		if (((seat.EventSeat.UserID != nil && *seat.EventSeat.UserID == *userId) && (seat.EventSeat.BlockedUntil != nil && seat.EventSeat.BlockedUntil.After(time.Now()))) && !seat.EventSeat.Booked) || *seat.EventSeat.ID == *eventSeatId {
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

	// @Collin - I think this is not needed
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

func (esc *EventSeatController) GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSlectedSeatsDTO, *models.KTSError) {
	return esc.EventSeatRepo.GetSelectedSeats(eventId, userId)
}
