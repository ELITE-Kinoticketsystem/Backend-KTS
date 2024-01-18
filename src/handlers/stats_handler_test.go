package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTotalVisitsHandler(t *testing.T) {
	samplesTotalVisitsDay := samples.GetSampleDayVisitsStatsTwoArrays()
	samplesTotalVisitsMonth := samples.GetSampleMonthVisitsStatsTwoArrays()
	samplesTotalVisitsYear := samples.GetSampleYearVisitsStatsTwoArrays()

	testCases := []struct {
		name            string
		startTimeString string
		startTime       time.Time
		endTimeString   string
		endTime         time.Time
		filterBy        string
		setExpectations func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:            "Failed - startTime",
			startTimeString: "2024-01-01wdawd",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - endTime",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03dawdadw",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - filterBy",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "dawdawd",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - firstDate higher than endDate",
			startTimeString: "2024-01-05",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - Function call",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {
				mockController.EXPECT().GetTotalVisits(startTime, endTime, filterBy).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:            "Success - Day",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {
				mockController.EXPECT().GetTotalVisits(startTime, endTime, filterBy).Return(samplesTotalVisitsDay, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsDay,
		},
		{
			name:            "Success - month",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-03-03",
			endTime:         time.Date(2024, 3, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "month",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {
				mockController.EXPECT().GetTotalVisits(startTime, endTime, filterBy).Return(samplesTotalVisitsMonth, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsMonth,
		},
		{
			name:            "Success - year",
			startTimeString: "2022-01-01",
			startTime:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "year",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string) {
				mockController.EXPECT().GetTotalVisits(startTime, endTime, filterBy).Return(samplesTotalVisitsYear, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsYear,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			statsController := mocks.NewMockStatsControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/stats/visits/"+tc.filterBy+"/"+tc.startTimeString+"/"+tc.endTimeString, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "filterBy", Value: tc.filterBy}, {Key: "from", Value: tc.startTimeString}, {Key: "til", Value: tc.endTimeString}}

			tc.setExpectations(statsController, tc.startTime, tc.endTime, tc.filterBy)

			// WHEN
			handlers.GetTotalVisitsHandler(statsController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetTotalVisitsForTheatreHandler(t *testing.T) {
	samplesTotalVisitsDay := samples.GetSampleDayVisitsStatsTwoArrays()
	samplesTotalVisitsMonth := samples.GetSampleMonthVisitsStatsTwoArrays()
	samplesTotalVisitsYear := samples.GetSampleYearVisitsStatsTwoArrays()

	testCases := []struct {
		name            string
		startTimeString string
		startTime       time.Time
		endTimeString   string
		endTime         time.Time
		filterBy        string
		theatreName     string
		setExpectations func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:            "Failed - startTime",
			startTimeString: "2024-01-01wdawd",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - endTime",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03dawdadw",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - filterBy",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "dawdawd",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - firstDate higher than endDate",
			startTimeString: "2024-01-05",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - firstDate higher than endDate",
			startTimeString: "2024-01-05",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:            "Failed - Function call",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {
				mockController.EXPECT().GetTotalVisitsForTheatre(startTime, endTime, filterBy, theatreName).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:            "Success - Day",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "day",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {
				mockController.EXPECT().GetTotalVisitsForTheatre(startTime, endTime, filterBy, theatreName).Return(samplesTotalVisitsDay, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsDay,
		},
		{
			name:            "Success - month",
			startTimeString: "2024-01-01",
			startTime:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-03-03",
			endTime:         time.Date(2024, 3, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "month",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {
				mockController.EXPECT().GetTotalVisitsForTheatre(startTime, endTime, filterBy, theatreName).Return(samplesTotalVisitsMonth, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsMonth,
		},
		{
			name:            "Success - year",
			startTimeString: "2022-01-01",
			startTime:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			endTimeString:   "2024-01-03",
			endTime:         time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			filterBy:        "year",
			theatreName:     "theatre1",
			setExpectations: func(mockController *mocks.MockStatsControllerI, startTime time.Time, endTime time.Time, filterBy string, theatreName string) {
				mockController.EXPECT().GetTotalVisitsForTheatre(startTime, endTime, filterBy, theatreName).Return(samplesTotalVisitsYear, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplesTotalVisitsYear,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			statsController := mocks.NewMockStatsControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/stats/visits/"+tc.filterBy+"/"+tc.startTimeString+"/"+tc.endTimeString+"/"+tc.theatreName, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "filterBy", Value: tc.filterBy}, {Key: "from", Value: tc.startTimeString}, {Key: "til", Value: tc.endTimeString}, {Key: "theatreName", Value: tc.theatreName}}

			tc.setExpectations(statsController, tc.startTime, tc.endTime, tc.filterBy, tc.theatreName)

			// WHEN
			handlers.GetTotalVisitsForTheatreHandler(statsController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetOrdersForStatsHandler(t *testing.T) {
	orders := samples.GetGetOrderDto()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockStatsControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Failed",
			setExpectations: func(mockController *mocks.MockStatsControllerI) {
				mockController.EXPECT().GetOrdersForStats().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockStatsControllerI) {
				mockController.EXPECT().GetOrdersForStats().Return(orders, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   orders,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			statsController := mocks.NewMockStatsControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/stats/orders/", nil)
			c.Request = req

			tc.setExpectations(statsController)

			// WHEN
			handlers.GetOrdersForStatsHandler(statsController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
