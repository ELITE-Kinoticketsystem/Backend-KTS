package controllers

import (
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
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

func TestGetTotalVisits(t *testing.T) {

	testCases := []struct {
		name            string
		start           time.Time
		end             time.Time
		in              string
		setExpectations func(mockRepo mocks.MockStatsRepositoryI, startTime time.Time, endTime time.Time, in string)
		expectedStats   *models.StatsVisitsTwoArrays
		expectedError   *models.KTSError
	}{
		{
			name:  "Failed",
			start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			in:    "day",
			setExpectations: func(mockRepo mocks.MockStatsRepositoryI, startTime time.Time, endTime time.Time, in string) {
				mockRepo.EXPECT().GetTotalVisits(startTime, endTime, in).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedStats: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name:  "Success",
			start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			in:    "day",
			setExpectations: func(mockRepo mocks.MockStatsRepositoryI, startTime time.Time, endTime time.Time, in string) {
				mockRepo.EXPECT().GetTotalVisits(startTime, endTime, in).Return(
					samples.GetSampleDayVisitsStats(), nil,
				)
			},
			expectedStats: samples.GetSampleDayVisitsStatsTwoArrays(),
			expectedError: nil,
		},
		{
			name:  "Success",
			start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2024, 3, 3, 0, 0, 0, 0, time.UTC),
			in:    "month",
			setExpectations: func(mockRepo mocks.MockStatsRepositoryI, startTime time.Time, endTime time.Time, in string) {
				mockRepo.EXPECT().GetTotalVisits(startTime, endTime, in).Return(
					samples.GetSampleMonthVisitsStats(), nil,
				)
			},
			expectedStats: samples.GetSampleMonthVisitsStatsTwoArrays(),
			expectedError: nil,
		},
		{
			name:  "Success",
			start: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			in:    "year",
			setExpectations: func(mockRepo mocks.MockStatsRepositoryI, startTime time.Time, endTime time.Time, in string) {
				mockRepo.EXPECT().GetTotalVisits(startTime, endTime, in).Return(
					samples.GetSampleYearVisitsStats(), nil,
				)
			},
			expectedStats: samples.GetSampleYearVisitsStatsTwoArrays(),
			expectedError: nil,
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
			tc.setExpectations(*mockRepoMock, tc.start, tc.end, tc.in)

			// WHEN
			// call DeleteReview with review data
			visitStats, err := statsController.GetTotalVisits(tc.start, tc.end, tc.in)

			// THEN
			// check expected error and id
			assert.Equal(t, visitStats, tc.expectedStats)
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}
}
