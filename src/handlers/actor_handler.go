package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetActorByIdHandler(actorController controllers.ActorControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		actorId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_INTERNAL_ERROR)
			return
		}
		actor, kts_err := actorController.GetActorById(&actorId)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, actor)
	}
}

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
