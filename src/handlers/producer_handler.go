package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetProducerByIdHandler(producerController controllers.ProducerControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		producerId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		producer, kts_err := producerController.GetProducerById(&producerId)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, producer)
	}
}

func GetProducersHandler(producerController controllers.ProducerControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		producers, err := producerController.GetProducers()
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, producers)
	}
}

func CreateProducerHandler(producerController controllers.ProducerControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var producerDto models.CreateProducerDTO
		if err := c.ShouldBindJSON(&producerDto); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		producer, kts_err := producerController.CreateProducer(&producerDto)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, producer)
	}
}
