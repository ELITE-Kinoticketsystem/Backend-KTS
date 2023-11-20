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

	exists, err := uc.UserRepo.CheckIfEmailExists(registrationData.Email)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if exists {
		return kts_errors.KTS_EMAIL_EXISTS
	}

	user := schemas.User{
		Id:        &userId,
		FirstName: "",
		LastName:  "",
		Email:     registrationData.Email,
		Password:  string(hash),
		/* Address */
	}

	err = uc.UserRepo.CreateUser(user)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}
