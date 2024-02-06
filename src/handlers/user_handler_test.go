package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
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
			body: samples.GetSampleRegistrationData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, registrationData interface{}) {
				mockController.EXPECT().RegisterUser(registrationData).Return(
					&models.LoginResponse{
						User: samples.GetSampleUser(),
						/* Token */
						/* RefreshToken */
					}, nil)
			},
			expectedResponseBody: samples.GetSampleUser(),
			expectedStatus:       http.StatusCreated,
		},
		{
			name: "Internal Error",
			body: samples.GetSampleRegistrationData(),
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
			body: samples.GetSampleRegistrationData(),
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
			body: samples.GetSampleRegistrationData(),
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
			body: samples.GetSampleLoginData(),
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
			body: samples.GetSampleLoginData(),
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
			body: samples.GetSampleLoginData(),
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
			body: samples.GetSampleLoginData(),
			setExpectations: func(mockController *mocks.MockUserControllerI, loginData models.LoginRequest) {
				user := samples.GetSampleUser()
				mockController.EXPECT().LoginUser(loginData).Return(
					&models.LoginResponse{
						User: user,
						/* Token */
						/* RefreshToken */
					}, nil)
			},
			expectedResponseBody: samples.GetSampleUser(),
			expectedStatus:       http.StatusOK,
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

func TestLogoutUser(t *testing.T) {
	expectedStatus := http.StatusOK
	expectedResponseBody := ""
	expectedCookieHeader := "token=; Path=/; Max-Age=0; HttpOnly; Secure; SameSite=None"

	// GIVEN
	// create mock context
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	// WHEN
	// call LogoutUserHandler with mock context
	LogoutUserHandler(c)

	// THEN
	// check the HTTP status code, body and cookie
	assert.Equal(t, expectedStatus, w.Code, "wrong HTTP status code")
	assert.Equal(t, expectedResponseBody, w.Body.String(), "wrong response body")
	assert.Equal(t, expectedCookieHeader, w.Header().Get("Set-Cookie"), "wrong cookie")

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
			body:            "",
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: models.KTSError{
				KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "BAD_REQUEST"},
				Details:         "Email is empty",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:            "No data - Binding failed",
			body:            nil,
			setExpectations: func(mockController *mocks.MockUserControllerI) {},
			expectedResponseBody: models.KTSError{
				KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "BAD_REQUEST"},
				Details:         "Email is empty",
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
			body:            "",
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

func TestTestJwtToken(t *testing.T) {

	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "GetJWTToken",
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

			// create mock request
			req := httptest.NewRequest("POST", "/auth/register", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations

			// WHEN
			TestJwtToken(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
		})
	}
}

func TestLoggedInHandler(t *testing.T) {
	id := utils.NewUUID()

	// token, refreshToken, jwtError := utils.GenerateJWT(id)
	token, _, jwtError := utils.GenerateJWT(id)
	if jwtError != nil {
		t.Error("JWT generation failed")
	}

	testCases := []struct {
		name                 string
		setToken             bool
		token                string
		expectedStatus       int
		expectedResponseBody interface{}
	}{
		{
			name:           "GetJWTToken",
			setToken:       false,
			token:          "",
			expectedStatus: http.StatusOK,
			expectedResponseBody: models.LoggedInResponse{
				LoggedIn: false,
				Id:       nil,
			},
		},
		{
			name:           "GetJWTToken",
			setToken:       true,
			token:          token,
			expectedStatus: http.StatusOK,
			expectedResponseBody: models.LoggedInResponse{
				LoggedIn: true,
				Id:       id,
			},
		},
		{
			name:           "GetJWTToken",
			setToken:       true,
			token:          "token",
			expectedStatus: http.StatusOK,
			expectedResponseBody: models.LoggedInResponse{
				LoggedIn: false,
				Id:       nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock request
			req := httptest.NewRequest("POST", "/auth/logged-in", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			if tc.setToken {
				c.Request.Header.Set("Cookie", "token="+tc.token)
			}

			// WHEN
			LoggedInHandler(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")

			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetUserHandler(t *testing.T) {
	user := samples.GetSampleUser()
	testCases := []struct {
		name                 string
		setContextUser       bool
		setExpectations      func(mockController *mocks.MockUserControllerI)
		expectedStatus       int
		expectedResponseBody interface{}
	}{
		{
			name:           "Success",
			setContextUser: true,
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().GetUserById(user.ID).Return(&user, nil)
			},
			expectedStatus:       http.StatusOK,
			expectedResponseBody: user,
		},
		{
			name:                 "No user",
			setContextUser:       false,
			setExpectations:      func(mockController *mocks.MockUserControllerI) {},
			expectedStatus:       http.StatusUnauthorized,
			expectedResponseBody: gin.H{"errorMessage": "UNAUTHORIZED"},
		},
		{
			name:           "Internal Error",
			setContextUser: true,
			setExpectations: func(mockController *mocks.MockUserControllerI) {
				mockController.EXPECT().GetUserById(user.ID).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus:       http.StatusInternalServerError,
			expectedResponseBody: gin.H{"errorMessage": "INTERNAL_ERROR"},
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
			req := httptest.NewRequest("POST", "/auth/get-me", nil)
			req.Header.Set("Content-Type", "application/json")

			if tc.setContextUser {
				ctx := context.WithValue(req.Context(), models.ContextKeyUserID, user.ID)
				c.Request = req.WithContext(ctx)
			} else {
				c.Request = req
			}

			// define expectations
			tc.setExpectations(userController)

			// WHEN
			// call GetUserHandler with mock context
			GetUserHandler(userController)(c)

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

func TestIsAdmin(t *testing.T) {
	userId := samples.GetSampleUser().ID
	adminId := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	testCases := []struct {
		name                 string
		userId               *uuid.UUID
		setContextUser       bool
		expectedStatus       int
		expectedResponseBody interface{}
	}{
		{
			name:                 "Admin",
			userId: 			 &adminId,
			setContextUser:       true,
			expectedStatus:       http.StatusOK,
			expectedResponseBody: true,
		},
		{
			name:                 "No admin",
			userId:               userId,
			setContextUser:       true,
			expectedStatus:       http.StatusOK,
			expectedResponseBody: false,
		},
		{
			name:                 "No user",
			setContextUser:       false,
			expectedStatus:       http.StatusUnauthorized,
			expectedResponseBody: gin.H{"errorMessage": "UNAUTHORIZED"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock request
			req := httptest.NewRequest("POST", "/auth/get-me", nil)
			req.Header.Set("Content-Type", "application/json")

			if tc.setContextUser {
				ctx := context.WithValue(req.Context(), models.ContextKeyUserID, tc.userId)
				c.Request = req.WithContext(ctx)
			} else {
				c.Request = req
			}

			// WHEN
			// call GetUserHandler with mock context
			IsAdminHandler(c)

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
