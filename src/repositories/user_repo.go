package repositories

import (
	"fmt"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type UserRepositoryI interface {
	GetUserById(id *uuid.UUID) (*model.Users, *models.KTSError)
	GetUserByUsername(username string) (*model.Users, *models.KTSError)
	CreateUser(user model.Users) *models.KTSError
	CheckIfUsernameExists(username string) *models.KTSError
	CheckIfEmailExists(email string) *models.KTSError
}

type UserRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (ur *UserRepository) GetUserById(id *uuid.UUID) (*model.Users, *models.KTSError) {
	var user model.Users
	stmt := mysql.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.Firstname,
		table.Users.Lastname,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.ID.EQ(utils.MysqlUuid(id)),
	)
	err := stmt.Query(ur.DatabaseManager.GetDatabaseConnection(), &user)
	if err != nil {
		if err.Error() == "jet: sql: no rows in result set" {
			return nil, kts_errors.KTS_USER_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*model.Users, *models.KTSError) {
	var user model.Users
	stmt := mysql.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.Firstname,
		table.Users.Lastname,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Username.EQ(mysql.String(username)),
	)
	err := stmt.Query(ur.DatabaseManager.GetDatabaseConnection(), &user)
	if err != nil {
		if err.Error() == "jet: sql: no rows in result set" {
			return nil, kts_errors.KTS_USER_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(user model.Users) *models.KTSError {
	stmt := table.Users.INSERT(
		table.Users.ID,
		table.Users.Username,
		table.Users.Email,
		table.Users.Password,
		table.Users.Firstname,
		table.Users.Lastname,
	).VALUES(
		utils.MysqlUuid(user.ID),
		user.Username,
		user.Email,
		user.Password,
		user.Firstname,
		user.Lastname,
	)
	fmt.Println(stmt.Sql())
	_, err := stmt.Exec(ur.DatabaseManager.GetDatabaseConnection())

	fmt.Println(err)
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}

func (ur *UserRepository) CheckIfUsernameExists(username string) *models.KTSError {
	count, err := utils.CountStatement(table.Users, table.Users.Username.EQ(mysql.String(username)), ur.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if count > 0 {
		return kts_errors.KTS_USERNAME_EXISTS
	}
	return nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) *models.KTSError {
	count, err := utils.CountStatement(table.Users, table.Users.Email.EQ(mysql.String(email)), ur.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if count > 0 {
		return kts_errors.KTS_EMAIL_EXISTS
	}
	return nil
}
