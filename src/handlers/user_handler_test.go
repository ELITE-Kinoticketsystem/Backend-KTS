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

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func TestCreateUser(t *testing.T) {
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