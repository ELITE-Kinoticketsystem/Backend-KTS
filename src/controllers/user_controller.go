package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type UserControllerI interface {
	RegisterUser(registrationData models.RegistrationRequest) *models.KTSError
	CheckEmail(email string) *models.KTSError
	CheckUsername(username string) *models.KTSError
}

type UserController struct {
	UserRepo repositories.UserRepositoryI
}

func (uc *UserController) RegisterUser(registrationData models.RegistrationRequest) *models.KTSError {
	userId := uuid.New()

	hash, err := utils.HashPassword(registrationData.Password)
	if err != nil {
		return kts_errors.KTS_UPSTREAM_ERROR
	}

	kts_err := uc.UserRepo.CheckIfEmailExists(registrationData.Email)
	if kts_err != nil {
		return kts_err
	}

	user := schemas.User{
		Id:        &userId,
		Username:  registrationData.Username,
		Email:     registrationData.Email,
		Password:  string(hash),
		FirstName: registrationData.FirstName,
		LastName:  registrationData.LastName,
	}

	kts_err = uc.UserRepo.CreateUser(user)
	if kts_err != nil {
		return kts_err
	}
	return nil
}

func (uc *UserController) CheckEmail(email string) *models.KTSError {
	return uc.UserRepo.CheckIfEmailExists(email)
}

func (uc *UserController) CheckUsername(username string) *models.KTSError {
	return uc.UserRepo.CheckIfUsernameExists(username)
}
