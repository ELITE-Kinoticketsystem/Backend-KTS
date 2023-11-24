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
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockUserControllerI, registrationData interface{})
		expectedStatus  int
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
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Internal Error",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Email exists",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Upstream Error",
			body: utils.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(nil, kts_errors.KTS_UPSTREAM_ERROR)
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
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {},
			expectedStatus:  http.StatusBadRequest,
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

		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		name            string
		body            models.LoginRequest
		setExpectations func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest)
		expectedStatus  int
	}{
		{
			name:            "Empty field",
			body:            models.LoginRequest{Username: "", Password: "Passwort"},
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name: "User not found",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_USER_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Internal error",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "Incorrect password",
			body: utils.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				mockController.EXPECT().LoginUser(loginData).Return(nil, kts_errors.KTS_CREDENTIALS_INVALID)
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
		})
	}
}
func TestCheckEmail(t *testing.T) {
	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockUserControllerI)
		expectedStatus  int
	}{
		{
			name: "Exists",
			body: models.CheckEmailRequest{
				Email: "collin.forslund@gmail.com",
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckEmail(gomock.Any()).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Doesn't exist",
			body: models.CheckEmailRequest{
				Email: "collin.forslund@gmail.com",
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckEmail(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedStatus:  http.StatusBadRequest,
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

		})
	}
}

func TestCheckUsername(t *testing.T) {
	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockUserControllerI)
		expectedStatus  int
	}{
		{
			name: "Exists",
			body: models.CheckUsernameRequest{
				Username: "Collinho el niño",
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckUsername(gomock.Any()).Return(kts_errors.KTS_USERNAME_EXISTS)
			},
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Doesn't exist",
			body: models.CheckUsernameRequest{
				Username: "collin.forslund@gmail.com",
			},
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().CheckUsername(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:            "Malformatted data",
			body:            map[string]string{},
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:            "No data",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedStatus:  http.StatusBadRequest,
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

		})
	}
}
