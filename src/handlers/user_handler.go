package handlers

import (
	"log"
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		log.Println(registrationData)
		if err != nil ||
			utils.ContainsEmptyString(
				registrationData.Username, registrationData.Email, registrationData.Password,
			) {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}
		kts_err := userCtrl.RegisterUser(registrationData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, *kts_err)
			return
		}

		c.AbortWithStatus(http.StatusCreated)
	}
}
