//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var UserMovies = newUserMoviesTable("KinoTicketSystem", "user_movies", "")

type userMoviesTable struct {
	mysql.Table

	// Columns
	UserID   mysql.ColumnString
	MovieID  mysql.ColumnString
	ListType mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type UserMoviesTable struct {
	userMoviesTable

	NEW userMoviesTable
}

// AS creates new UserMoviesTable with assigned alias
func (a UserMoviesTable) AS(alias string) *UserMoviesTable {
	return newUserMoviesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserMoviesTable with assigned schema name
func (a UserMoviesTable) FromSchema(schemaName string) *UserMoviesTable {
	return newUserMoviesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserMoviesTable with assigned table prefix
func (a UserMoviesTable) WithPrefix(prefix string) *UserMoviesTable {
	return newUserMoviesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserMoviesTable with assigned table suffix
func (a UserMoviesTable) WithSuffix(suffix string) *UserMoviesTable {
	return newUserMoviesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserMoviesTable(schemaName, tableName, alias string) *UserMoviesTable {
	return &UserMoviesTable{
		userMoviesTable: newUserMoviesTableImpl(schemaName, tableName, alias),
		NEW:             newUserMoviesTableImpl("", "new", ""),
	}
}

func newUserMoviesTableImpl(schemaName, tableName, alias string) userMoviesTable {
	var (
		UserIDColumn   = mysql.StringColumn("user_id")
		MovieIDColumn  = mysql.StringColumn("movie_id")
		ListTypeColumn = mysql.StringColumn("list_type")
		allColumns     = mysql.ColumnList{UserIDColumn, MovieIDColumn, ListTypeColumn}
		mutableColumns = mysql.ColumnList{}
	)

	return userMoviesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:   UserIDColumn,
		MovieID:  MovieIDColumn,
		ListType: ListTypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
