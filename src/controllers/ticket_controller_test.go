package controllers

import (
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTicketById(t *testing.T) {
	sampleTicket := samples.GetSampleTicket()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID)
		expectedTicket  *models.TicketDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().GetTicketById(id).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedTicket: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "One Ticket",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().GetTicketById(sampleTicket.ID).Return(sampleTicket, nil)
			},
			expectedTicket: sampleTicket,
			expectedError:  nil,
		},
		{
			name: "Error while querying for ticket",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().GetTicketById(id).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedTicket: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketRepoMock := mocks.NewMockTicketRepositoryI(mockCtrl)
			ticketController := TicketController{
				TicketRepo: ticketRepoMock,
			}

			// define expectations
			tc.setExpectations(*ticketRepoMock, sampleTicket.ID)

			// WHEN
			ticket, kts_err := ticketController.GetTicketById(sampleTicket.ID)

			// THEN
			assert.Equal(t, tc.expectedTicket, ticket)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestValidateTicket(t *testing.T) {
	id := utils.NewUUID()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "ticket conflict",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().ValidateTicket(id).Return(kts_errors.KTS_CONFLICT)
			},
			expectedError: kts_errors.KTS_CONFLICT,
		},
		{
			name: "Ticket validated",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().ValidateTicket(id).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while validating ticket",
			setExpectations: func(mockRepo mocks.MockTicketRepositoryI, id *uuid.UUID) {
				mockRepo.EXPECT().ValidateTicket(id).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketRepoMock := mocks.NewMockTicketRepositoryI(mockCtrl)
			ticketController := TicketController{
				TicketRepo: ticketRepoMock,
			}

			// define expectations
			tc.setExpectations(*ticketRepoMock, id)

			// WHEN
			kts_err := ticketController.ValidateTicket(id)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
