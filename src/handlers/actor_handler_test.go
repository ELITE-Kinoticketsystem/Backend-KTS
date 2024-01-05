package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetActorsHandler(t *testing.T) {
	sampleActors := samples.GetSampleActors()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockActorControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockActorControllerI) {
				mockController.EXPECT().GetActors().Return(sampleActors, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleActors,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockActorControllerI) {
				mockController.EXPECT().GetActors().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			actorController := mocks.NewMockActorControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/actors/", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(actorController)

			// WHEN
			handlers.GetActorsHandler(actorController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetActorByIdHandler(t *testing.T) {
	sampleActor := samples.GetSampleActor()

	testCases := []struct {
		name            string
		actorId         string
		setExpectations func(mockController *mocks.MockActorControllerI, actorId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:    "Success",
			actorId: sampleActor.ID.String(),
			setExpectations: func(mockController *mocks.MockActorControllerI, actorId string) {
				mockController.EXPECT().GetActorById(gomock.Any()).Return(sampleActor, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleActor,
		},
		{
			name:    "Internal error",
			actorId: sampleActor.ID.String(),
			setExpectations: func(mockController *mocks.MockActorControllerI, actorId string) {
				mockController.EXPECT().GetActorById(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:    "Genre Name contains empty string",
			actorId: "",
			setExpectations: func(mockController *mocks.MockActorControllerI, actorId string) {

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
			actorController := mocks.NewMockActorControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/genre/"+tc.actorId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.actorId}}

			tc.setExpectations(actorController, tc.actorId)

			// WHEN
			handlers.GetActorByIdHandler(actorController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreateActorHandler(t *testing.T) {
	sampleCreateActor := samples.GetSampleActor()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockActorControllerI, actor interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleCreateActor,
			setExpectations: func(mockController *mocks.MockActorControllerI, actor interface{}) {
				mockController.EXPECT().CreateActor(gomock.Any()).Return(sampleCreateActor.ID, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   sampleCreateActor.ID,
		},
		{
			name: "Internal error",
			body: sampleCreateActor,
			setExpectations: func(mockController *mocks.MockActorControllerI, actor interface{}) {
				mockController.EXPECT().CreateActor(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			actorController := mocks.NewMockActorControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("POST", "/actors/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(actorController, tc.body)

			// WHEN
			handlers.CreateActorHandler(actorController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
