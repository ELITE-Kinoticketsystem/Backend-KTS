package repositories

import (
	"database/sql"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/go-jet/jet/v2/mysql"
)

type UserRepositoryI interface {
	GetUserByUsername(username string) (*model.Users, *models.KTSError)
	CreateUser(user model.Users) *models.KTSError
	CheckIfUsernameExists(username string) *models.KTSError
	CheckIfEmailExists(email string) *models.KTSError
}

type UserRepository struct {
	DatabaseManager managers.DatabaseManagerI
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

	query, args := stmt.Sql()
	err := ur.DatabaseManager.ExecuteQueryRow(query, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Firstname,
		&user.Lastname,
	)

	if err != nil {
		if err == sql.ErrNoRows {
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
	).MODEL(
		user,
	)

	_, err := stmt.Exec(ur.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}

func (ur *UserRepository) CheckIfUsernameExists(username string) *models.KTSError {
	stmt := mysql.SELECT(
		mysql.COUNT(table.Users.ID),
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Username.EQ(mysql.String(username)),
	)

	query, args := stmt.Sql()
	exists, err := ur.DatabaseManager.CheckIfExists(query, args...)

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if exists {
		return kts_errors.KTS_USERNAME_EXISTS
	}
	return nil
}

func (ur *UserRepository) CheckIfEmailExists(email string) *models.KTSError {
	stmt := mysql.SELECT(
		mysql.COUNT(table.Users.ID),
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Email.EQ(mysql.String(email)),
	)

	query, args := stmt.Sql()
	exists, err := ur.DatabaseManager.CheckIfExists(query, args...)

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	if exists {
		return kts_errors.KTS_EMAIL_EXISTS
	}
	return nil
}
