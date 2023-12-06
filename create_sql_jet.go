package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/template"
	"github.com/google/uuid"

	mysql2 "github.com/go-jet/jet/v2/generator/mysql"
	mysql3 "github.com/go-jet/jet/v2/mysql"
)

func main() {
	log.Println("Starting Jet Generator for MySQL")

	var (
		dbHost    = os.Getenv("DB_HOST")
		dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
		dbUser    = os.Getenv("DB_USER")
		dbPass    = os.Getenv("DB_PASSWORD")
		dbName    = os.Getenv("DB_NAME")
	)

	var dbMySQLConnection = mysql2.DBConnection{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPass,
		DBName:   dbName,
	}

	GenerateJetMySQL(dbMySQLConnection)
}

func GenerateJetMySQL(dbMySQLConnection mysql2.DBConnection) {
	const gen_path = "./src/.gen"

	err := mysql2.Generate(
		gen_path,
		dbMySQLConnection,
		template.Default(mysql3.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)

									switch defaultTableModelField.Type.Name {
									case "[]byte":
										defaultTableModelField.Type = template.NewType(&uuid.UUID{})
									}
									return defaultTableModelField
								})
						}),
					)
			}),
	)

	if err != nil {
		panic(err)
	}
}
