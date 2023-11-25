package managers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManagerI interface {
	ExecuteStatement(query string, args... any) (sql.Result, error)
	ExecuteQuery(query string, args... any) (*sql.Rows, error)
	ExecuteQueryRow(query string, args... any) *sql.Row
	CheckIfExists(query string, args ...any) (bool, error)
}

type DatabaseManager struct {
	Connection *sql.DB
}

func (dm *DatabaseManager) ExecuteStatement(query string, args... any) (sql.Result, error) {
	result, err := dm.Connection.Exec(query, args...)
	return result, err
}

func (dm *DatabaseManager) ExecuteQuery(query string, args... any) (*sql.Rows, error) {
	result, err := dm.Connection.Query(query, args...)
	return result, err
}

func (dm *DatabaseManager) ExecuteQueryRow(query string, args... any) *sql.Row {
	result := dm.Connection.QueryRow(query, args...)
	return result
}

func (dm *DatabaseManager) CheckIfExists(query string, args ...any) (bool, error) {
	var count int
	err := dm.Connection.QueryRow(query, args...).Scan(&count)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InitializeDB() (*sql.DB, error) {
	var (
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_NAME")
	)
	if utils.ContainsEmptyString(dbHost, dbPort, dbUser, dbPass, dbName) {
		return nil, fmt.Errorf("error initializing datbase connection: environment variables not set")
	}

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	
	return db, nil
}