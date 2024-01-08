package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

type UserControllerI interface {
	RegisterUser(registrationData models.RegistrationRequest) (*models.LoginResponse, *models.KTSError)
	LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.KTSError)
	CheckEmail(email string) *models.KTSError
	CheckUsername(username string) *models.KTSError
}

type UserController struct {
	UserRepo repositories.UserRepositoryI
}

func (uc *UserController) RegisterUser(registrationData models.RegistrationRequest) (*models.LoginResponse, *models.KTSError) {
	userId := myid.New()

	hash, err := utils.HashPassword(registrationData.Password)
	if err != nil {
		return nil, kts_errors.KTS_UPSTREAM_ERROR
	}

	kts_err := uc.UserRepo.CheckIfEmailExists(registrationData.Email)
	if kts_err != nil {
		return nil, kts_err
	}

	user := model.Users{
		ID:        userId,
		Username:  &registrationData.Username,
		Email:     registrationData.Email,
		Password:  string(hash),
		Firstname: &registrationData.FirstName,
		Lastname:  &registrationData.LastName,
	}

	kts_err = uc.UserRepo.CreateUser(user)
	if kts_err != nil {
		return nil, kts_err
	}

	token, refreshToken, err := utils.GenerateJWT(&user.ID)
	if err != nil {
		return nil, kts_errors.KTS_UPSTREAM_ERROR
	}

	return &models.LoginResponse{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserController) LoginUser(loginData models.LoginRequest) (*models.LoginResponse, *models.KTSError) {

	// get user from database
	user, kts_err := uc.UserRepo.GetUserByUsername(loginData.Username)
	if kts_err != nil {
		return nil, kts_err
	}

	// check if password is correct
	if ok := utils.ComparePasswordHash(loginData.Password, user.Password); !ok {
		return nil, kts_errors.KTS_CREDENTIALS_INVALID
	}

	// generate JWT token
	token, refreshToken, err := utils.GenerateJWT(&user.ID)
	if err != nil {
		return nil, kts_errors.KTS_UPSTREAM_ERROR
	}

	return &models.LoginResponse{
		User:         *user,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UserController) CheckEmail(email string) *models.KTSError {
	return uc.UserRepo.CheckIfEmailExists(email)
}

func (uc *UserController) CheckUsername(username string) *models.KTSError {
	return uc.UserRepo.CheckIfUsernameExists(username)
}
