package managers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DatabaseManagerI interface {
	ExecuteStatement(query string) (sql.Result, error)
	ExecuteQuery(query string) (*sql.Rows, error)
	ExecuteQueryRow(query string) *sql.Row
}

type DatabaseManager struct {
	Connection *sql.DB
}

func (dm *DatabaseManager) ExecuteStatement(query string) (sql.Result, error) {
	result, err := dm.Connection.Exec(query)
	return result, err
}

func (dm *DatabaseManager) ExecuteQuery(query string) (*sql.Rows, error) {
	result, err := dm.Connection.Query(query)
	return result, err
}

func (dm *DatabaseManager) ExecuteQueryRow(query string) *sql.Row {
	result := dm.Connection.QueryRow(query)
	return result
}

func InitializeDB() (*sql.DB, error) {
	var (
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_NAME")
	)

	config := mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName: dbName,
	}
	db, _ := sql.Open("mysql", config.FormatDSN())

	return db, nil
}