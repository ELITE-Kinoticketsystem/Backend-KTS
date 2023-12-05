package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		name             string
		registrationData models.RegistrationRequest
		setExpectations  func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest)
		expectedResponse *models.LoginResponse
		expectedError    *models.KTSError
	}{
		{
			name:             "Email exists",
			registrationData: utils.GetSampleRegistrationData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(kts_errors.KTS_EMAIL_EXISTS)
			},
			expectedResponse: nil,
			expectedError:    kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:             "Email internal error",
			registrationData: utils.GetSampleRegistrationData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponse: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:             "CreateUser internal error",
			registrationData: utils.GetSampleRegistrationData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				user := model.Users{
					/* Id */
					Username:  &registrationData.Username,
					Email:     registrationData.Email,
					Password:  registrationData.Password,
					Firstname: &registrationData.FirstName,
					Lastname:  &registrationData.LastName,
				}

				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(nil)
				mockRepo.EXPECT().CreateUser(utils.EqUserMatcher(user, registrationData.Password)).Return(kts_errors.KTS_INTERNAL_ERROR)

			},
			expectedResponse: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:             "Success",
			registrationData: utils.GetSampleRegistrationData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, registrationData models.RegistrationRequest) {
				user := model.Users{
					/* Id */
					Username:  &registrationData.Username,
					Email:     registrationData.Email,
					Password:  registrationData.Password, // unhashed password
					Firstname: &registrationData.FirstName,
					Lastname:  &registrationData.LastName,
				}

				mockRepo.EXPECT().CheckIfEmailExists(registrationData.Email).Return(nil)
				mockRepo.EXPECT().CreateUser(utils.EqUserMatcher(user, registrationData.Password)).Return(nil)
			},
			expectedResponse: &models.LoginResponse{
				User: utils.GetSampleUser(),
				/* Token */
				/* RefreshToken */
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
			loginResponse, err := userController.RegisterUser(registrationData)

			// THEN
			// check expected error and user
			assert.Equal(t, err, tc.expectedError, "wrong error")
			if tc.expectedResponse != nil {
				assert.False(t, loginResponse == nil, "loginResponse should not be nil")
				assert.True(t, utils.UserEqual(tc.expectedResponse.User, loginResponse.User))
			}
		})
	}

}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		name             string
		loginData        models.LoginRequest
		setExpectations  func(mockRepo mocks.MockUserRepositoryI, loginData models.LoginRequest)
		expectedError    *models.KTSError
		expectedResponse *models.LoginResponse
	}{
		{
			name:      "User not found",
			loginData: utils.GetSampleLoginData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, loginData models.LoginRequest) {
				mockRepo.EXPECT().GetUserByUsername(loginData.Username).Return(nil, kts_errors.KTS_USER_NOT_FOUND)
			},
			expectedError:    kts_errors.KTS_USER_NOT_FOUND,
			expectedResponse: nil,
		},
		{
			name:      "Internal error",
			loginData: utils.GetSampleLoginData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, loginData models.LoginRequest) {
				mockRepo.EXPECT().GetUserByUsername(loginData.Username).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
			expectedResponse: nil,
		},
		{
			name:      "Incorrect password",
			loginData: models.LoginRequest{Username: "Collinho el niño", Password: "incorrect"},
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, loginData models.LoginRequest) {
				user := utils.GetSampleUser()
				mockRepo.EXPECT().GetUserByUsername(loginData.Username).Return(&user, nil)
			},
			expectedError:    kts_errors.KTS_CREDENTIALS_INVALID,
			expectedResponse: nil,
		},
		{
			name:      "success",
			loginData: utils.GetSampleLoginData(),
			setExpectations: func(mockRepo mocks.MockUserRepositoryI, loginData models.LoginRequest) {
				user := utils.GetSampleUser()
				mockRepo.EXPECT().GetUserByUsername(loginData.Username).Return(&user, nil)
			},
			expectedError: nil,
			expectedResponse: &models.LoginResponse{
				User: utils.GetSampleUser(),
				/* Token */
				/* RefreshToken */
			},
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
			loginData := tc.loginData

			// define expectations
			tc.setExpectations(*userRepoMock, loginData)

			// WHEN
			// call RegisterUser with registrationData
			loginResponse, err := userController.LoginUser(loginData)

			// THEN
			// check expected error
			assert.Equal(t, tc.expectedError, err, "wrong error")
			if tc.expectedResponse != nil {
				assert.False(t, loginResponse == nil, "loginResponse should not be nil")
				assert.Equal(t, tc.expectedResponse.User, loginResponse.User, "wrong response")
			}
		})
	}

}

func TestCheckEmail(t *testing.T) {
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

func TestCheckUsername(t *testing.T) {
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
