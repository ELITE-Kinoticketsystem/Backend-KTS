package utils

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/gin-gonic/gin"
)

func HandleErrorAndAbort(c *gin.Context, err *models.KTSError) {
	log.Printf("Error while handling request: %d %v %v", err.Status, err.ErrorMessage, err.Details)
	c.AbortWithStatusJSON(err.Status, err)
}
