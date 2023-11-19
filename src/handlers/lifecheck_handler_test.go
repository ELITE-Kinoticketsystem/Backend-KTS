package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
)

func TestLifeCheckHandler(t *testing.T) {
	// Create a mock context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the LifeCheckHandler function with the mock context
	LifeCheckHandler()(c)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code is not OK")

	// Parse the response body
	var response models.LifeCheckResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	// Check the Alive field
	assert.True(t, response.Alive, "Alive field is not true")

	// Check the Timestamp field
	assert.NotZero(t, response.Timestamp, "Timestamp field is zero")
}
