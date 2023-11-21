package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLifeCheckRoute(t *testing.T) {
	// GIVEN
	// Initialize the router
	router := createRouter(new(sql.DB)) // sufficient for this case, as no db is used

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/lifecheck", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// WHEN
	// Test the lifecheck route
	router.ServeHTTP(w, req)

	// THEN
	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, w.Code, "HTTP status code is not OK")
}
