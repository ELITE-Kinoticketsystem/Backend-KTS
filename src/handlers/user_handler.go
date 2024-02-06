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

// @Summary Register user
// @Description Register user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.RegistrationRequest true "User data"
// @Success 201 {object} model.Users
// @Failure 400 {object} models.KTSErrorMessage
// @Router /auth/register [post]
func RegisterUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registrationData models.RegistrationRequest
		err := c.ShouldBind(&registrationData)
		if err != nil ||
			utils.ContainsEmptyString(
				registrationData.Username, registrationData.Email, registrationData.Password,
				registrationData.FirstName, registrationData.LastName,
			) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		// user is logged in after registration
		loginResponse, kts_err := userCtrl.RegisterUser(registrationData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		utils.SetJWTCookies(c, loginResponse.Token, loginResponse.RefreshToken, false)
		c.JSON(http.StatusCreated, loginResponse.User)
	}
}

// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.LoginRequest true "User data"
// @Success 200 {object} model.Users
// @Failure 400 {object} models.KTSErrorMessage
// @Router /auth/login [post]
func LoginUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData models.LoginRequest
		err := c.ShouldBind(&loginData)
		if err != nil ||
			utils.ContainsEmptyString(
				loginData.Username, loginData.Password,
			) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		loginResponse, kts_err := userCtrl.LoginUser(loginData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		utils.SetJWTCookies(c, loginResponse.Token, loginResponse.RefreshToken, false)
		c.JSON(http.StatusOK, loginResponse.User)
	}
}

func LogoutUserHandler(c *gin.Context) {
	utils.SetJWTCookies(c, "", "", true)
	c.Status(http.StatusOK)
}

// @Summary Check email
// @Description Check email
// @Tags Users
// @Accept  json
// @Produce  json
// @Param checkEmailRequest body models.CheckEmailRequest true "Email data"
// @Success 200
// @Failure 400 {object} models.KTSErrorMessage
// @Router /auth/check-email [post]
func CheckEmailHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckEmailRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST.WithDetails("Email is empty"))
			return
		}

		err := userCtrl.CheckEmail(requestData.Email)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

// @Summary Check username
// @Description Check username
// @Tags Users
// @Accept  json
// @Produce  json
// @Param checkUsernameRequest body models.CheckUsernameRequest true "Username data"
// @Success 200
// @Failure 400 {object} models.KTSErrorMessage
// @Router /auth/check-username [post]
func CheckUsernameHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData models.CheckUsernameRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		err := userCtrl.CheckUsername(requestData.Username)
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.Status(http.StatusOK)
	}
}

func TestJwtToken(c *gin.Context) {
	c.Status(http.StatusOK)
}

// @Summary Logged in
// @Description Check if user is logged in
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.LoggedInResponse
// @Failure 400 {object} models.KTSErrorMessage
// @Router /auth/logged-in [get]
func LoggedInHandler(c *gin.Context) {
	var token string

	// check if token is set
	token, err := c.Cookie("token")
	if err != nil {
		// token is not set, check if refresh token is set
		token, err = c.Cookie("refreshToken")
		if err != nil {
			c.JSON(http.StatusOK, models.LoggedInResponse{
				LoggedIn: false,
			})
			return
		}
	}

	id, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusOK, models.LoggedInResponse{
			LoggedIn: false,
		})
		return
	}

	c.JSON(http.StatusOK, models.LoggedInResponse{
		LoggedIn: true,
		Id:       id,
	})
}

func GetUserHandler(userCtrl controllers.UserControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		user, kts_err := userCtrl.GetUserById(userId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func IsAdminHandler(c *gin.Context) {
	userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
	if !ok {
		utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
		return
	}
	adminId := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
	c.JSON(http.StatusOK, *userId == adminId)
}
