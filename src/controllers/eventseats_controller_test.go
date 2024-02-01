package controllers

import (
	"reflect"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
				RowNr:          0,
				ColumnNr:       0,
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
				RowNr:          0,
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
				BlockedUntil: &blockedUntil,
				UserID:       userId,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       0,
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
			ID:             eventSeats[0].EventSeat.ID,
			RowNr:          eventSeats[0].Seat.RowNr,
			ColumnNr:       eventSeats[0].Seat.ColumnNr,
			Available:      true,
			BlockedByOther: false,
			Category:       eventSeats[0].SeatCategory.CategoryName,
			Price:          eventSeats[0].EventSeatCategory.Price,
		},
		{
			ID:             eventSeats[1].EventSeat.ID,
			RowNr:          eventSeats[1].Seat.RowNr,
			ColumnNr:       eventSeats[1].Seat.ColumnNr,
			Available:      true,
			BlockedByOther: false,
			Category:       eventSeats[1].SeatCategory.CategoryName,
			Price:          eventSeats[1].EventSeatCategory.Price,
		},
		{
			ID:             eventSeats[2].EventSeat.ID,
			RowNr:          eventSeats[2].Seat.RowNr,
			ColumnNr:       eventSeats[2].Seat.ColumnNr,
			Available:      false,
			BlockedByOther: false,
			Category:       eventSeats[2].SeatCategory.CategoryName,
			Price:          eventSeats[2].EventSeatCategory.Price,
		},
	}

	expectedSeatRows := []models.GetSeatsForSeatSelectorDTO{
		seatsSlice[0],
		seatsSlice[1],
		seatsSlice[2],
	}

	expectedCurrentUserSeats := []models.GetSeatsForSeatSelectorDTO{
		seatsSlice[2],
	}

	tests := []struct {
		name                     string
		expectFuncs              func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T)
		expectedError            *models.KTSError
		expectedSeatRows         *[]models.GetSeatsForSeatSelectorDTO
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
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(int64(1), nil)
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError: nil,
			expectTime:    true,
		},
		{
			name: "Block event seat - error",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().BlockEventSeatIfAvailable(eventId, eventSeatId, userId, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
		{
			name: "Block event seat seats are not next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				eventSeats[2].Seat.RowNr = 99

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError: kts_errors.KTS_CONFLICT,
			expectTime:    false,
		},
		{
			name: "Block event seat seats are not next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				eventSeats[2].Seat.RowNr = 99

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)

			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
		{
			name: "Block event seat",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().BlockEventSeatIfAvailable(eventId, eventSeatId, userId, gomock.Any()).Return(nil)
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(int64(1), kts_errors.KTS_INTERNAL_ERROR)
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
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
func TestEventSeatController_AreUserSeatsNextToEachOther(t *testing.T) {

	eventSeatId := utils.NewUUID()
	eventId := utils.NewUUID()
	userId := utils.NewUUID()

	test := []struct {
		name           string
		expectFuncs    func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T)
		expectedError  *models.KTSError
		expectedResult bool
	}{
		{
			name: "Are user seats next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError:  nil,
			expectedResult: true,
		},
		{
			name: "Are user seats next to each other - error",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
			expectedResult: false,
		},
		{
			name: "Are user seats next to each other wrong row nr - false",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)
				eventSeats[2].Seat.RowNr = 99
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError:  nil,
			expectedResult: false,
		},
		{
			name: "Are user seats next to each other wrong column nr - false",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)
				eventSeats[2].Seat.ColumnNr = 99
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError:  nil,
			expectedResult: false,
		},
		{
			name: "Are user seats next to each other empty seat in between",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := GetEventSeatsDTO(eventId, userId, eventSeatId)
				eventSeats[3].Seat.Type = string(utils.EMPTY)
				eventSeats[3].EventSeat.UserID = nil
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedError:  nil,
			expectedResult: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			// Given

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)

			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			tt.expectFuncs(mockEventSeatRepo, t)

			// Test when seats are next to each other
			areNextToEachOther, err := eventSeatController.AreUserSeatsNextToEachOther(eventId, userId, eventSeatId)

			if err != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
			}

			if areNextToEachOther != tt.expectedResult {
				t.Errorf("Expected result: %v, but got: %v", tt.expectedResult, areNextToEachOther)
			}
		})
	}
}

func GetEventSeatsDTO(eventId *uuid.UUID, userId *uuid.UUID, eventSeatId *uuid.UUID) []models.GetEventSeatsDTO {
	blockedUntil := time.Now().Add(time.Minute * 5)
	return []models.GetEventSeatsDTO{
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
				ColumnNr:       3,
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
}
func TestEventSeatController_AreUserSeatsNextToEachOtherWithoutSeat(t *testing.T) {
	eventId := utils.NewUUID()
	userId := utils.NewUUID()
	eventSeatId := utils.NewUUID()

	blockedUntil := time.Now().Add(time.Minute * 5)

	tests := []struct {
		name           string
		expectFuncs    func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T)
		expectedResult bool
		expectedError  *models.KTSError
	}{
		{
			name: "Are user seats next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {

				eventSeats := []models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 2,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 3,
						},
					},
				}

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name: "Are user seats not next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := []models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 2,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 3,
						},
					},
				}

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: false,
			expectedError:  nil,
		},
		{
			name: "Error getting event seats",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResult: false,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "No seats found",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := []models.GetEventSeatsDTO{}
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name: "Are user seats are in other rows",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := []models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 2,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    2,
							ColumnNr: 3,
						},
					},
				}

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: false,
			expectedError:  nil,
		},
		{
			name: "Empty seats",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := []models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 2,
							Type:     string(utils.EMPTY),
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 3,
						},
					},
				}

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name: "Only the seat to be removed",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				eventSeats := []models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
				}

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&eventSeats, nil)
			},
			expectedResult: true,
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)

			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			tt.expectFuncs(mockEventSeatRepo, t)

			// When
			result, err := eventSeatController.AreUserSeatsNextToEachOtherWithoutSeat(eventId, userId, eventSeatId)

			// Then
			if result != tt.expectedResult {
				t.Errorf("Expected result: %v, but got: %v", tt.expectedResult, result)
			}

			if err != tt.expectedError {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
func TestEventSeatController_UnblockEventSeat(t *testing.T) {
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
			name: "Unblock event seat",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				currentTime := time.Now()
				blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&[]models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
				}, nil)
				mockEventSeatRepo.EXPECT().UnblockEventSeat(eventId, eventSeatId, userId).Return(nil)
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(int64(1), nil)
			},
			expectedError: nil,
			expectTime:    true,
		},
		{
			name: "Unblock event seat - error",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				currentTime := time.Now()
				blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&[]models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
				}, nil)
				mockEventSeatRepo.EXPECT().UnblockEventSeat(eventId, eventSeatId, userId).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
		{
			name: "Unblock event seat - seats are not next to each other",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				currentTime := time.Now()
				blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&[]models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 2,
						},
					},
					{
						EventSeat: model.EventSeats{
							ID:           utils.NewUUID(),
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 3,
						},
					},
				}, nil)
			},
			expectedError: kts_errors.KTS_CONFLICT,
			expectTime:    false,
		},
		{
			name: "Unblock event seat - error getting event seats",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
		{
			name: "Unblock event seat ",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				currentTime := time.Now()
				blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&[]models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
				}, nil)
				mockEventSeatRepo.EXPECT().UnblockEventSeat(eventId, eventSeatId, userId).Return(nil)
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(int64(0), kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
			expectTime:    false,
		},
		{
			name: "Unblock event seat ",
			expectFuncs: func(mockEventSeatRepo *mocks.MockEventSeatRepoI, t *testing.T) {
				currentTime := time.Now()
				blockedUntil := currentTime.Add(utils.BLOCKED_TICKET_TIME)

				mockEventSeatRepo.EXPECT().GetEventSeats(eventId).Return(&[]models.GetEventSeatsDTO{
					{
						EventSeat: model.EventSeats{
							ID:           eventSeatId,
							UserID:       userId,
							Booked:       false,
							BlockedUntil: &blockedUntil,
						},
						Seat: model.Seats{
							RowNr:    1,
							ColumnNr: 1,
						},
					},
				}, nil)
				mockEventSeatRepo.EXPECT().UnblockEventSeat(eventId, eventSeatId, userId).Return(nil)
				mockEventSeatRepo.EXPECT().UpdateBlockedUntilTimeForUserEventSeats(eventId, userId, gomock.Any()).Return(int64(0), nil)
			},
			expectedError: nil,
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
			blockedUntil, err := eventSeatController.UnblockEventSeat(eventId, eventSeatId, userId)

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

func TestUnblockAllEventSeats(t *testing.T) {

	testCases := []struct {
		name            string
		eventId         *uuid.UUID
		userId          *uuid.UUID
		setExpectations func(mockRepo mocks.MockEventSeatRepoI, eventId *uuid.UUID, userId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name:    "Failed",
			eventId: utils.NewUUID(),
			userId:  utils.NewUUID(),
			setExpectations: func(mockRepo mocks.MockEventSeatRepoI, eventId *uuid.UUID, userId *uuid.UUID) {
				mockRepo.EXPECT().UnblockAllEventSeats(eventId, userId).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)
			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			// define expectations
			tc.setExpectations(*mockEventSeatRepo, tc.eventId, tc.userId)

			// WHEN
			kts_err := eventSeatController.UnblockAllEventSeats(tc.eventId, tc.userId)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetSelectedSeats(t *testing.T) {

	testCases := []struct {
		name                 string
		eventId              *uuid.UUID
		userId               *uuid.UUID
		setExpectations      func(mockRepo mocks.MockEventSeatRepoI, eventId *uuid.UUID, userId *uuid.UUID)
		expectedSlectedSeats *[]models.GetSlectedSeatsDTO
		expectedError        *models.KTSError
	}{
		{
			name:    "Failed",
			eventId: utils.NewUUID(),
			userId:  utils.NewUUID(),
			setExpectations: func(mockRepo mocks.MockEventSeatRepoI, eventId *uuid.UUID, userId *uuid.UUID) {
				mockRepo.EXPECT().GetSelectedSeats(eventId, userId).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedSlectedSeats: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)
			eventSeatController := EventSeatController{
				EventSeatRepo: mockEventSeatRepo,
			}

			// define expectations
			tc.setExpectations(*mockEventSeatRepo, tc.eventId, tc.userId)

			// WHEN
			slectedSeats, kts_err := eventSeatController.GetSelectedSeats(tc.eventId, tc.userId)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedSlectedSeats, slectedSeats)
		})
	}
}
