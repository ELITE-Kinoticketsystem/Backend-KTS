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

var MovieProducers = newMovieProducersTable("KinoTicketSystem", "movie_producers", "")

type movieProducersTable struct {
	mysql.Table

	// Columns
	MovieID    mysql.ColumnString
	ProducerID mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MovieProducersTable struct {
	movieProducersTable

	NEW movieProducersTable
}

// AS creates new MovieProducersTable with assigned alias
func (a MovieProducersTable) AS(alias string) *MovieProducersTable {
	return newMovieProducersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MovieProducersTable with assigned schema name
func (a MovieProducersTable) FromSchema(schemaName string) *MovieProducersTable {
	return newMovieProducersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MovieProducersTable with assigned table prefix
func (a MovieProducersTable) WithPrefix(prefix string) *MovieProducersTable {
	return newMovieProducersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MovieProducersTable with assigned table suffix
func (a MovieProducersTable) WithSuffix(suffix string) *MovieProducersTable {
	return newMovieProducersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMovieProducersTable(schemaName, tableName, alias string) *MovieProducersTable {
	return &MovieProducersTable{
		movieProducersTable: newMovieProducersTableImpl(schemaName, tableName, alias),
		NEW:                 newMovieProducersTableImpl("", "new", ""),
	}
}

func newMovieProducersTableImpl(schemaName, tableName, alias string) movieProducersTable {
	var (
		MovieIDColumn    = mysql.StringColumn("movie_id")
		ProducerIDColumn = mysql.StringColumn("producer_id")
		allColumns       = mysql.ColumnList{MovieIDColumn, ProducerIDColumn}
		mutableColumns   = mysql.ColumnList{}
	)

	return movieProducersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		MovieID:    MovieIDColumn,
		ProducerID: ProducerIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
