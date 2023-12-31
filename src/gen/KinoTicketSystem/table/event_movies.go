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

var EventMovies = newEventMoviesTable("KinoTicketSystem", "event_movies", "")

type eventMoviesTable struct {
	mysql.Table

	// Columns
	EventID mysql.ColumnString
	MovieID mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EventMoviesTable struct {
	eventMoviesTable

	NEW eventMoviesTable
}

// AS creates new EventMoviesTable with assigned alias
func (a EventMoviesTable) AS(alias string) *EventMoviesTable {
	return newEventMoviesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventMoviesTable with assigned schema name
func (a EventMoviesTable) FromSchema(schemaName string) *EventMoviesTable {
	return newEventMoviesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EventMoviesTable with assigned table prefix
func (a EventMoviesTable) WithPrefix(prefix string) *EventMoviesTable {
	return newEventMoviesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EventMoviesTable with assigned table suffix
func (a EventMoviesTable) WithSuffix(suffix string) *EventMoviesTable {
	return newEventMoviesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEventMoviesTable(schemaName, tableName, alias string) *EventMoviesTable {
	return &EventMoviesTable{
		eventMoviesTable: newEventMoviesTableImpl(schemaName, tableName, alias),
		NEW:              newEventMoviesTableImpl("", "new", ""),
	}
}

func newEventMoviesTableImpl(schemaName, tableName, alias string) eventMoviesTable {
	var (
		EventIDColumn  = mysql.StringColumn("event_id")
		MovieIDColumn  = mysql.StringColumn("movie_id")
		allColumns     = mysql.ColumnList{EventIDColumn, MovieIDColumn}
		mutableColumns = mysql.ColumnList{}
	)

	return eventMoviesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		EventID: EventIDColumn,
		MovieID: MovieIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
