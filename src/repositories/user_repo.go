package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
)

type UserRepositoryI interface {
	CreateUser(user schemas.User) error
	CheckIfUsernameExists(username string) (bool, error)
	CheckIfEmailExists(email string) (bool, error)
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

func (ur *UserRepository) CheckIfUsernameExists(username string) (bool, error) {
	// not implemented, no username in database schema
	return false, nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) (bool, error) {
	exists, err := ur.DatabaseManager.CheckIfExists(
		"SELECT COUNT(*) FROM users WHERE email = ?", email,
	)
	if err != nil {
		return false, err
	}
	return exists, nil
}


type UserRepositoryMock struct {
	UserRepository
}

func (ur UserRepositoryMock) CheckIfEmailExists(email string) (bool, error) {
	return email == "exists@email.com", nil
}

func (ur *UserRepositoryMock) CreateUser(user schemas.User) error {
	return nil
}