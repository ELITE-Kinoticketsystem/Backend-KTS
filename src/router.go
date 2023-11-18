package main

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	router := gin.Default()

	// Attach Middleware

	// Create api groups, with special middleware

	// Create managers and repositories

	// Create controllers

	// Set routes
	router.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())

	return router
}
