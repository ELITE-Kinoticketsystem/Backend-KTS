package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		if err != nil ||
			utils.ContainsEmptyString(
				registrationData.Username, registrationData.Email, registrationData.Password,
				registrationData.FirstName, registrationData.LastName,
			) {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}
		// user is logged in after registration
		loginResponse, kts_err := userCtrl.RegisterUser(registrationData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, *kts_err)
			return
		}

		c.JSON(http.StatusCreated, loginResponse)
	}
}

func LoginUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData models.LoginRequest
		err := c.ShouldBind(&loginData)
		if err != nil ||
			utils.ContainsEmptyString(
				loginData.Username, loginData.Password,
			) {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}
		loginResponse, kts_err := userCtrl.LoginUser(loginData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, *kts_err)
			return
		}

		c.JSON(http.StatusOK, loginResponse)
	}
}

func CheckEmailHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckEmailRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}
		if utils.ContainsEmptyString(requestData.Email) {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}

		err := userCtrl.CheckEmail(requestData.Email)
		if err != nil {
			utils.HandleErrorAndAbort(c, *err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func CheckUsernameHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckUsernameRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}
		if utils.ContainsEmptyString(requestData.Username) {
			utils.HandleErrorAndAbort(c, *kts_errors.KTS_BAD_REQUEST)
			return
		}

		err := userCtrl.CheckUsername(requestData.Username)
		if err != nil {
			utils.HandleErrorAndAbort(c, *err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func TestJwtToken(c *gin.Context) {
	c.Status(http.StatusOK)
}
