package controllers

import (
	"log"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type EventSeatControllerI interface {
	GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError)
	BlockEventSeat(eventSeatId *uuid.UUID, userId *uuid.UUID) *models.KTSError
}

type EventSeatController struct {
	EventSeatRepo repositories.EventSeatRepoI
}

func (esc *EventSeatController) GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError) {
	return esc.EventSeatRepo.GetEventSeats(eventId)
}

func (esc *EventSeatController) BlockEventSeat(eventSeatId *uuid.UUID, userId *uuid.UUID) *models.KTSError {

	seat, err := esc.EventSeatRepo.GetEventSeat(eventSeatId)

	if err != nil {
		return err
	}

	if (seat.BlockedUntil != nil && seat.BlockedUntil.After(time.Now()) && seat.UserID != nil) || seat.Booked {
		return kts_errors.KTS_CONFLICT
	}

	currentTime := time.Now()
	blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)
	log.Println(blockedUntil)
	err = esc.EventSeatRepo.BlockEventSeat(eventSeatId, userId, &blockedUntil)

	if err != nil {
		return err
	}

	err = esc.EventSeatRepo.ResetTimerOnUserSeats(userId, &currentTime, &blockedUntil)

	if err != nil {
		return err
	}

	return nil
}
