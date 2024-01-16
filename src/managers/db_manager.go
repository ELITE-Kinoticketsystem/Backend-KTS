package managers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManagerI interface {
	GetDatabaseConnection() *sql.DB
	NewTransaction() (*sql.Tx, error)
}

type DatabaseManager struct {
	Connection *sql.DB
}

func (dm *DatabaseManager) GetDatabaseConnection() *sql.DB {
	return dm.Connection
}

func (dm *DatabaseManager) NewTransaction() (*sql.Tx, error) {
	return dm.Connection.Begin()
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
