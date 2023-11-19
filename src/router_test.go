package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifeCheckRoute(t *testing.T) {
	// Initialize the router
	router := createRouter()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/lifecheck", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Test the lifecheck route
	router.ServeHTTP(w, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, w.Code)
}
