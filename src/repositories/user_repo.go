package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
)

type UserRepositoryI interface {
	CreateUser(user schemas.User) error
	CheckIfUsernameExists(username string) *models.KTSError
	CheckIfEmailExists(email string) *models.KTSError
}

type UserRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ur *UserRepository) CreateUser(user schemas.User) error {
	_, err := ur.DatabaseManager.ExecuteStatement(
		"INSERT INTO users (id, firstname, lastname, email, age, password, address_id, user_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		user.Id, user.FirstName, user.LastName, user.Email, 0 /* user.Age (missing in schema) */, user.Password, user.AddressId, 0 /* user.Age (missing in schema) */,
	)
	return err
}

func (ur *UserRepository) CheckIfUsernameExists(username string) *models.KTSError {
	exists, err := ur.DatabaseManager.CheckIfExists(
		"SELECT COUNT(*) FROM users WHERE username = ?", username,
	)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if exists {
		return kts_errors.KTS_USERNAME_EXISTS
	}
	return nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) *models.KTSError {
	exists, err := ur.DatabaseManager.CheckIfExists(
		"SELECT COUNT(*) FROM users WHERE email = ?", email,
	)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if exists {
		return kts_errors.KTS_EMAIL_EXISTS
	}
	return nil
}