package controllers

import (
	"reflect"
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"go.uber.org/mock/gomock"
)

func TestEventSeatController_GetEventSeats(t *testing.T) {
	eventId := utils.NewUUID()
	userId := utils.NewUUID()

	blockedUntil := time.Now().Add(time.Minute * 5)

	eventSeats := []models.GetEventSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       1,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				Price:          100,
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				Price:          100,
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: &blockedUntil,
				UserID:       userId,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          2,
				ColumnNr:       1,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				Price:          100,
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
			},
		},
	}

	seatsSlice := []models.GetSeatsForSeatSelectorDTO{
		{
			ID:            eventSeats[0].EventSeat.ID,
			RowNr:         eventSeats[0].Seat.RowNr,
			ColumnNr:      eventSeats[0].Seat.ColumnNr,
			Available:     true,
			BookedByOther: false,
			Category:      eventSeats[0].SeatCategory.CategoryName,
			Price:         eventSeats[0].EventSeatCategory.Price,
		},
		{
			ID:            eventSeats[1].EventSeat.ID,
			RowNr:         eventSeats[1].Seat.RowNr,
			ColumnNr:      eventSeats[1].Seat.ColumnNr,
			Available:     true,
			BookedByOther: false,
			Category:      eventSeats[1].SeatCategory.CategoryName,
			Price:         eventSeats[1].EventSeatCategory.Price,
		},
		{
			ID:            eventSeats[2].EventSeat.ID,
			RowNr:         eventSeats[2].Seat.RowNr,
			ColumnNr:      eventSeats[2].Seat.ColumnNr,
			Available:     false,
			BookedByOther: false,
			Category:      eventSeats[2].SeatCategory.CategoryName,
			Price:         eventSeats[2].EventSeatCategory.Price,
		},
	}

	expectedSeatRows := [][]models.GetSeatsForSeatSelectorDTO{
		{
			seatsSlice[0],
			seatsSlice[1],
		},
		{
			seatsSlice[2],
		},
	}

	expectedCurrentUserSeats := []models.GetSeatsForSeatSelectorDTO{
		seatsSlice[2],
	}

	tests := []struct {
		name                     string
		expectFuncs              func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T)
		expectedError            *models.KTSError
		expectedSeatRows         *[][]models.GetSeatsForSeatSelectorDTO
		expectedCurrentUserSeats *[]models.GetSeatsForSeatSelectorDTO
		BlockedUntil             *time.Time
	}{
		{
			name: "Get event seats",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError:            nil,
			expectedSeatRows:         &expectedSeatRows,
			expectedCurrentUserSeats: &expectedCurrentUserSeats,
			BlockedUntil:             &blockedUntil,
		},
		{
			name: "Get event seats - error",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:            kts_errors.KTS_INTERNAL_ERROR,
			expectedSeatRows:         nil,
			expectedCurrentUserSeats: nil,
			BlockedUntil:             nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)

			tt.expectFuncs(mockEventSeatRepo, t)

			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			// When

			seatMap, currentUserSeats, timeBlock, err := eventSeatController.GetEventSeats(eventId, userId)

			// Then
			if !reflect.DeepEqual(seatMap, tt.expectedSeatRows) {
				t.Errorf("Expected seat map: %v, but got: %v", expectedSeatRows, seatMap)
			}

			if !reflect.DeepEqual(currentUserSeats, tt.expectedCurrentUserSeats) {
				t.Errorf("Expected current user seats: %v, but got: %v", expectedCurrentUserSeats, currentUserSeats)
			}

			if !reflect.DeepEqual(timeBlock, tt.BlockedUntil) {
				t.Errorf("Expected blocked until: %v, but got: %v", tt.BlockedUntil, timeBlock)
			}

			if err != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}

}

func TestEventSeatController_BlockEventSeat(t *testing.T) {
	eventId := utils.NewUUID()
	eventSeatId := utils.NewUUID()
	userId := utils.NewUUID()

	tests := []struct {
		name          string
		expectFuncs   func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T)
		expectedError *models.KTSError
		expectTime    bool
	}{
		{
			name: "Block event seat",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().BlockEventSeatIfAvailable(eventId, eventSeatId, userId, gomock.Any()).Return(nil)
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(nil)
			},
			expectedError: nil,
			expectTime:    true,
		},
		{
			name: "Block event seat - error",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().BlockEventSeatIfAvailable(eventId, eventSeatId, userId, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)

			tt.expectFuncs(mockEventSeatRepo, t)

			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			// When
			blockedUntil, err := eventSeatController.BlockEventSeat(eventId, eventSeatId, userId)

			// Then
			if err != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
			}

			if tt.expectTime && blockedUntil == nil {
				t.Errorf("Expected blocked until time but got nil")
			} else if !tt.expectTime && blockedUntil != nil {
				t.Errorf("Expected nil but got blocked until time")
			}
		})
	}
}
