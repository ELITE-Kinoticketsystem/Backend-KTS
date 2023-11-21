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

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func TestHandlerCreateUserCreated(t *testing.T) {
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
	registrationData := models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}
	jsonData, _ := json.Marshal(registrationData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations
	userController.EXPECT().RegisterUser(registrationData).Return(nil)

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusCreated, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserInternalError(t *testing.T) {
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
	registrationData := models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}
	jsonData, _ := json.Marshal(registrationData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations
	userController.EXPECT().RegisterUser(registrationData).Return(kts_errors.KTS_INTERNAL_ERROR)

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusInternalServerError, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserEmailExists(t *testing.T) {
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
	registrationData := models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}
	jsonData, _ := json.Marshal(registrationData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations
	userController.EXPECT().RegisterUser(registrationData).Return(kts_errors.KTS_EMAIL_EXISTS)

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusConflict, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserUpstreamError(t *testing.T) {
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
	registrationData := models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}
	jsonData, _ := json.Marshal(registrationData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations
	userController.EXPECT().RegisterUser(registrationData).Return(kts_errors.KTS_UPSTREAM_ERROR)

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusInternalServerError, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserEmptyField(t *testing.T) {
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
	registrationData := models.RegistrationRequest{
		Username: "",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}
	jsonData, _ := json.Marshal(registrationData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations: none

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusBadRequest, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserMalformattedData(t *testing.T) {
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
	malformattedData := map[string]string{
		"User": "Colllinho el niño", // User instead of Username
		"Email":    "collin.forslund@gmail.com",
		"Password": "Passwort",
		"FirstName": "Collin",
		"LastName": "Forslund",
	}
	jsonData, _ := json.Marshal(malformattedData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations: none

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusBadRequest, w.Code, "HTTP status code is not OK")
}

func TestHandlerCreateUserNoData(t *testing.T) {
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
	var malformattedData map[string]string
	jsonData, _ := json.Marshal(malformattedData)
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	// define expectations: none

	// WHEN
	// call RegisterUserHandler function with the mock context
	RegisterUserHandler(userController)(c)

	// THEN
	// check the HTTP status code
	assert.Equal(t, http.StatusBadRequest, w.Code, "HTTP status code is not OK")
}