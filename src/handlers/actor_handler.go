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

// @Summary Get actor by id
// @Description Get actor by id
// @Tags Actors
// @Accept  json
// @Produce  json
// @Param id path string true "Actor ID"
// @Success 200 {object} models.ActorDTO
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /actors/{id} [get]
func GetActorByIdHandler(actorController controllers.ActorControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		actorId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		actor, kts_err := actorController.GetActorById(&actorId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, actor)
	}
}

// @Summary Get actors
// @Description Get actors
// @Tags Actors
// @Accept  json
// @Produce  json
// @Success 200 {array} models.GetActorsDTO
// @Failure 500 {object} models.KTSErrorMessage
// @Router /actors [get]
func GetActorsHandler(actorController controllers.ActorControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		actors, err := actorController.GetActors()
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, actors)
	}
}

// @Summary Create actor
// @Description Create actor
// @Tags Actors
// @Accept  json
// @Produce  json
// @Param actor body models.CreateActorDTO true "Actor"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /actors [post]
func CreateActorHandler(actorController controllers.ActorControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var actorDto models.CreateActorDTO
		if err := c.ShouldBindJSON(&actorDto); utils.ContainsEmptyString(actorDto.Actors.Name) || err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		actor, kts_err := actorController.CreateActor(&actorDto)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, actor)
	}
}
