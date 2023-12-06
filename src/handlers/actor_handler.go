package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetActorByIdHandler(actorController controllers.ActorControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		actorId := uuid.MustParse(c.Param("id"))
		actor, err := actorController.GetActorById(&actorId)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
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
