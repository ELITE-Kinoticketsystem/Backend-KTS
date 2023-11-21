package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"go.uber.org/mock/gomock"
)

func TestRegisterEmailExists(t *testing.T) {
	// GIVEN
	// create mock user repo
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
	userController := UserController{
		UserRepo: userRepoMock,
	}

	// create mock data
	var registrationData = models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
	}

	// define expectations
	userRepoMock.EXPECT().CheckIfEmailExists("collin.forslund@gmail.com").Return(true, nil)

	// WHEN
	// call RegisterUser with registrationData
	err := userController.RegisterUser(registrationData)

	// THEN
	// check expected error
	if err != kts_errors.KTS_EMAIL_EXISTS {
		t.Fail()
	}
}

func TestRegisterEmailDoesntExist(t *testing.T) {
	// GIVEN
	// create mock user repo
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
	userController := &UserController{
		UserRepo: userRepoMock,
	}

	// create mock data
	var registrationData = models.RegistrationRequest{
		Username: "Colllinho el niño",
		Email:    "collin.forslund@gmail.com",
		Password: "Passwort",
		FirstName: "Collin",
		LastName: "Forslund",
	}

	// define expectations
	user := schemas.User{
		/* Id */
		FirstName: registrationData.FirstName,
		LastName: registrationData.LastName,
		Email: registrationData.Email,
		Age: 0,
		Password: registrationData.Password,
		/* Address*/
	}

	userRepoMock.EXPECT().CheckIfEmailExists("collin.forslund@gmail.com").Return(false, nil)
	userRepoMock.EXPECT().CreateUser(utils.EqUserMatcher(user, registrationData.Password)).Return(nil)

	// WHEN
	// call RegisterUser with registrationData
	err := userController.RegisterUser(registrationData)

	// THEN
	// check expected error
	if err == kts_errors.KTS_EMAIL_EXISTS {
		t.Fail()
	}
}
