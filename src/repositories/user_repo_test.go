package repositories

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	sampleUser := samples.GetSampleUser()
	id := sampleUser.ID
	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedUser    *model.Users
		expectedError   *models.KTSError
	}{
		{
			name: "No rows",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.id = ?;",
				).WithArgs(
					utils.EqUUID(id),
				).WillReturnError(sql.ErrNoRows)
			},
			expectedUser:  nil,
			expectedError: kts_errors.KTS_USER_NOT_FOUND,
		},
		{
			name: "Internal error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.id = ?;",
				).WithArgs(
					utils.EqUUID(id),
				).WillReturnError(sql.ErrConnDone)
			},
			expectedUser:  nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Success",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.id = ?;",
				).WithArgs(
					utils.EqUUID(id),
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"users.id", "users.username", "users.email", "users.password", "users.firstname", "users.lastname",
					}).AddRow(
						sampleUser.ID, sampleUser.Username, sampleUser.Email, sampleUser.Password, sampleUser.Firstname, sampleUser.Lastname,
					),
				)
			},
			expectedUser:  &sampleUser,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			userRepo := UserRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock)

			// WHEN
			// call GetUserByUsername with username
			user, kts_err := userRepo.GetUserById(&id)

			// THEN
			// check expected error, user and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedUser, user)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	sampleUser := samples.GetSampleUser()
	testCases := []struct {
		name            string
		username        string
		setExpectations func(mock sqlmock.Sqlmock, username string)
		expectedUser    *model.Users
		expectedError   *models.KTSError
	}{
		{
			name:     "No rows",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
					username,
				).WillReturnError(sql.ErrNoRows)
			},
			expectedUser:  nil,
			expectedError: kts_errors.KTS_USER_NOT_FOUND,
		},
		{
			name:     "Internal error",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
					username,
				).WillReturnError(sql.ErrConnDone)
			},
			expectedUser:  nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:     "Success",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery(
					"SELECT users.id AS \"users.id\",\n" +
						"users.username AS \"users.username\",\n" +
						"users.email AS \"users.email\",\n" +
						"users.password AS \"users.password\",\n" +
						"users.firstname AS \"users.firstname\",\n" +
						"users.lastname AS \"users.lastname\"\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
					username,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"users.id", "users.username", "users.email", "users.password", "users.firstname", "users.lastname",
					}).AddRow(
						sampleUser.ID, sampleUser.Username, sampleUser.Email, sampleUser.Password, sampleUser.Firstname, sampleUser.Lastname,
					),
				)
			},
			expectedUser:  &sampleUser,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
			// call GetUserByUsername with username
			user, kts_err := userRepo.GetUserByUsername(tc.username)

			// THEN
			// check expected error, user and expectations
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedUser, user)
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name            string
		data            model.Users
		setExpectations func(mock sqlmock.Sqlmock, user *model.Users)
		expectedError   *models.KTSError
	}{
		{
			name: "Success",
			data: samples.GetSampleUser(),
			setExpectations: func(mock sqlmock.Sqlmock, user *model.Users) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.users (id, username, email, password, firstname, lastname)\n"+
						"VALUES (?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(user.ID), user.Username, user.Email, user.Password, user.Firstname, user.Lastname,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Internal error",
			data: samples.GetSampleUser(),
			setExpectations: func(mock sqlmock.Sqlmock, user *model.Users) {
				mock.ExpectExec(
					"INSERT INTO `KinoTicketSystem`.users (id, username, email, password, firstname, lastname)\n"+
						"VALUES (?, ?, ?, ?, ?, ?);",
				).WithArgs(
					utils.EqUUID(user.ID), user.Username, user.Email, user.Password, user.Firstname, user.Lastname,
				).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock db manager
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}
			userRepo := UserRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			// define expectations
			tc.setExpectations(mock, &tc.data)

			// WHEN
			// call CreateUser with user data
			kts_err := userRepo.CreateUser(tc.data)

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
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
					username,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"COUNT(users.id)",
					}).AddRow(
						1,
					),
				)
			},
			expectedError: kts_errors.KTS_USERNAME_EXISTS,
		},
		{
			name:     "Username doesn't exist",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
					username,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"COUNT(users.id)",
					}).AddRow(
						0,
					),
				)
			},
			expectedError: nil,
		},
		{
			name:     "Internal error",
			username: "Collinho el niño",
			setExpectations: func(mock sqlmock.Sqlmock, username string) {
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.username = ?;",
				).WithArgs(
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
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.email = ?;",
				).WithArgs(
					email,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"COUNT(users.email)",
					}).AddRow(
						1,
					),
				)
			},
			expectedError: kts_errors.KTS_EMAIL_EXISTS,
		},
		{
			name:  "Email doesn't exist",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mock sqlmock.Sqlmock, email string) {
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.email = ?;",
				).WithArgs(
					email,
				).WillReturnRows(
					sqlmock.NewRows([]string{
						"COUNT(users.email)",
					}).AddRow(
						0,
					),
				)
			},
			expectedError: nil,
		},
		{
			name:  "Internal error",
			email: "collin.forslund@gmail.com",
			setExpectations: func(mock sqlmock.Sqlmock, email string) {
				mock.ExpectQuery(
					"SELECT COUNT(users.id)\n" +
						"FROM `KinoTicketSystem`.users\n" +
						"WHERE users.email = ?;",
				).WithArgs(
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
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
