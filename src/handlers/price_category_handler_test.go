package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdatePriceCategoryHandler(t *testing.T) {
	sampleUpdatePriceCategory := samples.GetSamplePriceCategory()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleUpdatePriceCategory,
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {
				mockController.EXPECT().UpdatePriceCategory(gomock.Any()).Return(&sampleUpdatePriceCategory.ID, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleUpdatePriceCategory.ID,
		},
		{
			name: "Internal error",
			body: sampleUpdatePriceCategory,
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {
				mockController.EXPECT().UpdatePriceCategory(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Internal error",
			body: model.PriceCategories{
				ID:           sampleUpdatePriceCategory.ID,
				CategoryName: "",
				Price:        sampleUpdatePriceCategory.Price,
			},
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
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
			priceCategoryController := mocks.NewMockPriceCategoryControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("PUT", "/price-categories/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(priceCategoryController, tc.body)

			// WHEN
			handlers.UpdatePriceCategoryHandler(priceCategoryController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestDeletePriceCategoryHandler(t *testing.T) {
	sampleUpdatePriceCategoryId := myid.NewUUID()

	testCases := []struct {
		name            string
		priceCategoryId string
		setExpectations func(mockController *mocks.MockPriceCategoryControllerI, PriceCategoryId string)
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:            "Success",
			priceCategoryId: sampleUpdatePriceCategoryId.String(),
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, PriceCategoryId string) {
				mockController.EXPECT().DeletePriceCategory(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   ``,
		},
		{
			name:            "Internal error",
			priceCategoryId: sampleUpdatePriceCategoryId.String(),
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, PriceCategoryId string) {
				mockController.EXPECT().DeletePriceCategory(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"INTERNAL_ERROR"}`,
		},
		{
			name:            "Movie not found",
			priceCategoryId: "",
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, PriceCategoryId string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"errorMessage":"BAD_REQUEST"}`,
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
			priceCategoryController := mocks.NewMockPriceCategoryControllerI(mockCtrl)

			req, _ := http.NewRequest("Delete", "/price-categories/"+tc.priceCategoryId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.priceCategoryId}}

			tc.setExpectations(priceCategoryController, tc.priceCategoryId)

			// WHEN
			handlers.DeletePriceCategoryHandler(priceCategoryController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Regexp(t, tc.expectedBody, w.Body.String(), "wrong HTTP response body")
		})
	}
}

func TestGetPriceCategoriesHandler(t *testing.T) {
	samplePriceCategories := samples.GetSamplePriceCategories()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockPriceCategoryControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI) {
				mockController.EXPECT().GetPriceCategories().Return(samplePriceCategories, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplePriceCategories,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI) {
				mockController.EXPECT().GetPriceCategories().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
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
			priceCategoryController := mocks.NewMockPriceCategoryControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/price-categories/", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(priceCategoryController)

			// WHEN
			handlers.GetPriceCategoriesHandler(priceCategoryController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetPriceCategoryByIdHandler(t *testing.T) {
	samplePriceCategory := samples.GetSamplePriceCategory()

	testCases := []struct {
		name            string
		priceCategoryId string
		setExpectations func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:            "Success",
			priceCategoryId: samplePriceCategory.ID.String(),
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryId string) {
				mockController.EXPECT().GetPriceCategoryById(gomock.Any()).Return(samplePriceCategory, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   samplePriceCategory,
		},
		{
			name:            "Internal error",
			priceCategoryId: samplePriceCategory.ID.String(),
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryId string) {
				mockController.EXPECT().GetPriceCategoryById(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:            "Movie title contains empty string",
			priceCategoryId: "",
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryId string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
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
			priceCategoryController := mocks.NewMockPriceCategoryControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/movies/"+tc.priceCategoryId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.priceCategoryId}}

			tc.setExpectations(priceCategoryController, tc.priceCategoryId)

			// WHEN
			handlers.GetPriceCategoryByIdHandler(priceCategoryController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreatePriceCategoryHandler(t *testing.T) {
	sampleCreatePriceCategory := samples.GetSamplePriceCategory()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleCreatePriceCategory,
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {
				mockController.EXPECT().CreatePriceCategory(gomock.Any()).Return(&sampleCreatePriceCategory.ID, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   sampleCreatePriceCategory.ID,
		},
		{
			name: "Internal error",
			body: sampleCreatePriceCategory,
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {
				mockController.EXPECT().CreatePriceCategory(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Title string empty",
			body: model.PriceCategories{
				ID:           sampleCreatePriceCategory.ID,
				CategoryName: "",
				Price:        sampleCreatePriceCategory.Price,
			},
			setExpectations: func(mockController *mocks.MockPriceCategoryControllerI, priceCategoryData interface{}) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
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
			priceCategoryController := mocks.NewMockPriceCategoryControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("POST", "/price-categories/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(priceCategoryController, tc.body)

			// WHEN
			handlers.CreatePriceCategoryHandler(priceCategoryController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
