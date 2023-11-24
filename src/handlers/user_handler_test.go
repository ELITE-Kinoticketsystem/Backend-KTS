package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		name                 string
		body                 interface{}
		setExpectations      func(mockController *mocks.MockUserControllerI, registrationData interface{})
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		{
			name: "Success",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(
					&models.LoginResponse{
						User: utils.GetSampleUser(),
						/* Token */
						/* RefreshToken */
					}, nil)
			},
			expectedResponseBody: models.LoginResponse{
				User: utils.GetSampleUser(),
				/* Token */
				/* RefreshToken */
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Internal Error",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Email exists",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "EMAIL_EXISTS",
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Upstream Error",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_UPSTREAM_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "UPSTREAM_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Empty field",
			body: models.RegistrationRequest{
				Username:  "",
				Email:     "collin.forslund@gmail.com",
				Password:  "Passwort",
				FirstName: "Collin",
				LastName:  "Forslund",
			},
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userController := mocks.NewMockUserControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(userController, tc.body)

			// WHEN
			// call RegisterUserHandler with mock context
			RegisterUserHandler(userController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		name                 string
		body                 models.LoginRequest
		setExpectations      func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest)
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		{
			name:            "Empty field",
			body:            models.LoginRequest{Username: "", Password: "Passwort"},
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "User not found",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_USER_NOT_FOUND)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "USER_NOT_FOUND",
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Internal error",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Incorrect password",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_CREDENTIALS_INVALID)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "CREDENTIALS_INVALID",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Success",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				user := utils.GetSampleUser()
				mockController.EXPECT().LoginUser(loginData).Return(
					&models.LoginResponse{
						User: user,
						/* Token */
						/* RefreshToken */
					}, nil)
			},
			expectedResponseBody: models.LoginResponse{
				User: utils.GetSampleUser(),
				/* Token */
				/* RefreshToken */
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userController := mocks.NewMockUserControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(userController, tc.body)

			// WHEN
			// call LoginUserHandler with mock context
			LoginUserHandler(userController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
func TestCheckEmail(t *testing.T) {
	email := "collin.forslund@gmail.com"
	testCases := []struct {
		name                 string
		body                 interface{}
		setExpectations      func(mockController *mocks.MockUserControllerI)
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		{
			name: "Exists",
			body: models.CheckEmailRequest{
				Email: email,
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckEmail(email).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "EMAIL_EXISTS",
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Doesn't exist",
			body: models.CheckEmailRequest{
				Email: email,
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckEmail(email).Return(nil)
			},
			expectedResponseBody: nil,
			expectedStatus:       http.StatusOK,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userController := mocks.NewMockUserControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(userController)

			// WHEN
			// call CheckEmailHandler with mock context
			CheckEmailHandler(userController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			if w.Body.Len() == 0 {
				assert.True(t, tc.expectedResponseBody == nil, "expected empty response body")
			} else {
				expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
				assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
			}
		})
	}
}

func TestCheckUsername(t *testing.T) {
	username := "Collinho el niño"
	testCases := []struct {
		name                 string
		body                 interface{}
		setExpectations      func(mockController *mocks.MockUserControllerI)
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		{
			name: "Exists",
			body: models.CheckUsernameRequest{
				Username: username,
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckUsername(username).Return(kts_errors.KTS_USERNAME_EXISTS)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "USERNAME_EXISTS",
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Doesn't exist",
			body: models.CheckUsernameRequest{
				Username: username,
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckUsername(username).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userController := mocks.NewMockUserControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(userController)

			// WHEN
			// call CheckUsernameHandler with mock context
			CheckUsernameHandler(userController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			if w.Body.Len() == 0 {
				assert.True(t, tc.expectedResponseBody == nil, "expected empty response body")
			} else {
				expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
				assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
			}
		})
	}
}
