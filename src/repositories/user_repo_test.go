package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name            string
		data            *schemas.User
		setExpectations func(mock sqlmock.Sqlmock, user *schemas.User)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			data: getSampleUser(),
			setExpectations: func(mock sqlmock.Sqlmock, user *schemas.User) {
				mock.ExpectExec("INSERT INTO users").WithArgs(
					user.Id, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.AddressId,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: getSampleUser(),
			setExpectations: func(mock sqlmock.Sqlmock, user *schemas.User) {
				mock.ExpectExec("INSERT INTO users").WithArgs(
					user.Id, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.AddressId,
				).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			userRepo := UserRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.data)

			// WHEN
			// call CreateUser with user data
			kts_err := userRepo.CreateUser(*tc.data)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCheckIfUsernameExists(t *testing.T) {
	testCases := []struct {
		name            string
		username        string
		setExpectations func(mock sqlmock.Sqlmock, username string)
		expectedError   *models.KTSError
	}{
		{
			name:     "Username exists",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE username = \\?").WithArgs(
					username,
				).WillReturnRows(
					sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1),
				)
			},
			expectedError: kts_errors.KTS_USERNAME_EXISTS,
		},
		{
			name:     "Username doesn't exist",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE username = \\?").WithArgs(
					username,
				).WillReturnRows(
					sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(0),
				)
			},
			expectedError: nil,
		},
		{
			name:     "Internal error",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE username = \\?").WithArgs(
					username,
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			userRepo := UserRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.username)

			// WHEN
			// call CreateUser with user data
			kts_err := userRepo.CheckIfUsernameExists(tc.username)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCheckIfEmailExists(t *testing.T) {
	testCases := []struct {
		name            string
		email           string
		setExpectations func(mock sqlmock.Sqlmock, email string)
		expectedError   *models.KTSError
	}{
		{
			name:  "Email exists",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mock sqlmock.Sqlmock, email string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE email = \\?").WithArgs(
					email,
				).WillReturnRows(
					sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1),
				)
			},
			expectedError: kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:  "Email doesn't exist",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mock sqlmock.Sqlmock, email string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE email = \\?").WithArgs(
					email,
				).WillReturnRows(
					sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(0),
				)
			},
			expectedError: nil,
		},
		{
			name:  "Internal error",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mock sqlmock.Sqlmock, email string) {
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users WHERE email = \\?").WithArgs(
					email,
				).WillReturnError(
					sqlmock.ErrCancelled,
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			userRepo := UserRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, tc.email)

			// WHEN
			// call CreateUser with user data
			kts_err := userRepo.CheckIfEmailExists(tc.email)

			// THEN
			// check expected error and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func getSampleUser() *schemas.User {
	id, _ := uuid.Parse("550e8400-e29b-11d4-a716-446655440000")
	return &schemas.User{
		Id:        &id,
		Username:  "Collinho el niño",
		Email:     "collin.forslund@gmail.com",
		Password:  "Passwort",
		FirstName: "Collin",
		LastName:  "Forslund",
	}
}
