package handlers

import (
	"net/http"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/gin-gonic/gin"
)

func LifeCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		response := &models.LifeCheckResponse{
			Alive: true,
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusOK, response)
	}
}