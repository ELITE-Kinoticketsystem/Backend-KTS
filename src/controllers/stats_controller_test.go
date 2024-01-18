package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetOrdersForStats(t *testing.T) {
	orders := samples.GetGetOrderDto()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockStatsRepositoryI)
		expectedOrders  *[]models.GetOrderDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			setExpectations: func(mockRepo mocks.MockStatsRepositoryI) {
				mockRepo.EXPECT().GetOrdersForStats().Return(orders, nil)
			},
			expectedOrders: orders,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock review repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepoMock := mocks.NewMockStatsRepositoryI(mockCtrl)
			statsController := StatsController{
				StatsRepo: mockRepoMock,
			}

			// define expectations
			tc.setExpectations(*mockRepoMock)

			// WHEN
			// call DeleteReview with review data
			orders, err := statsController.GetOrdersForStats()

			// THEN
			// check expected error and id
			assert.Equal(t, orders, tc.expectedOrders)
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}
