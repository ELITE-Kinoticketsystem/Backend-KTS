package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestControllerRegister(t *testing.T) {
	testCases := []struct {
		name             string
		registrationData models.RegistrationRequest
		setExpectations  func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest)
		expectedError    *models.KTSError
	}{
		{
			name:             "Email exists",
			registrationData: getSampleRegistratonData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedError: kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:             "Email internal error",
			registrationData: getSampleRegistratonData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:             "create user internal error",
			registrationData: getSampleRegistratonData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				user := schemas.User{
					/* Id */
					Username:  registrationData.Username,
					Email:     registrationData.Email,
					Password:  registrationData.Password,
					FirstName: registrationData.FirstName,
					LastName:  registrationData.LastName,
				}

				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(nil)
				mockRepo.EXPECT().CreateUser(utils.EqUserMatcher(user, registrationData.Password)).Return(kts_errors.KTS_INTERNAL_ERROR)

			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:             "success",
			registrationData: getSampleRegistratonData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				user := schemas.User{
					/* Id */
					Username:  registrationData.Username,
					Email:     registrationData.Email,
					Password:  registrationData.Password,
					FirstName: registrationData.FirstName,
					LastName:  registrationData.LastName,
				}

				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(nil)
				mockRepo.EXPECT().CreateUser(utils.EqUserMatcher(user, registrationData.Password)).Return(nil)
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
			userController := UserController{
				UserRepo: userRepoMock,
			}

			// create mock data
			registrationData := tc.registrationData

			// define expectations
			tc.setExpectations(*userRepoMock, registrationData)

			// WHEN
			// call RegisterUser with registrationData
			err := userController.RegisterUser(registrationData)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}

}

func TestControllerCheckEmail(t *testing.T) {
	testCases := []struct {
		name            string
		email           string
		setExpectations func(mockRepo mocks.MockUserRepositoryI, email string)
		expectedError   *models.KTSError
	}{
		{
			name:  "Email exists",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, email string) {
				mockRepo.EXPECT().CheckIfEmailExists(email).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedError: kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:  "Internal error",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, email string) {
				mockRepo.EXPECT().CheckIfEmailExists(email).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
			userController := UserController{
				UserRepo: userRepoMock,
			}

			// define expectations
			tc.setExpectations(*userRepoMock, tc.email)

			// WHEN
			// call RegisterUser with registrationData
			err := userController.CheckEmail(tc.email)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}

}

func TestControllerCheckUsername(t *testing.T) {
	testCases := []struct {
		name            string
		username        string
		setExpectations func(mockRepo mocks.MockUserRepositoryI, email string)
		expectedError   *models.KTSError
	}{
		{
			name:     "Email exists",
			username: "Collinho el niño",
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, email string) {
				mockRepo.EXPECT().CheckIfUsernameExists(email).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedError: kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:     "Internal error",
			username: "Collinho el niño",
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, email string) {
				mockRepo.EXPECT().CheckIfUsernameExists(email).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			userRepoMock := mocks.NewMockUserRepositoryI(mockCtrl)
			userController := UserController{
				UserRepo: userRepoMock,
			}

			// define expectations
			tc.setExpectations(*userRepoMock, tc.username)

			// WHEN
			// call RegisterUser with registrationData
			err := userController.CheckUsername(tc.username)

			// THEN
			// check expected error
			assert.Equal(t, err, tc.expectedError, "wrong error")
		})
	}

}
func getSampleRegistratonData() models.RegistrationRequest {
	return models.RegistrationRequest{
		Username:  "Collinho el niño",
		Email:     "collin.forslund@gmail.com",
		Password:  "Passwort",
		FirstName: "Collin",
		LastName:  "Forslund",
	}
}
