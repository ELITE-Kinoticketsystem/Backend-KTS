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

var CinemaHalls = newCinemaHallsTable("KinoTicketSystem", "cinema_halls", "")

type cinemaHallsTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString
	Name      mysql.ColumnString
	Capacity  mysql.ColumnInteger
	TheatreID mysql.ColumnString
	Width     mysql.ColumnInteger
	Height    mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type CinemaHallsTable struct {
	cinemaHallsTable

	NEW cinemaHallsTable
}

// AS creates new CinemaHallsTable with assigned alias
func (a CinemaHallsTable) AS(alias string) *CinemaHallsTable {
	return newCinemaHallsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CinemaHallsTable with assigned schema name
func (a CinemaHallsTable) FromSchema(schemaName string) *CinemaHallsTable {
	return newCinemaHallsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CinemaHallsTable with assigned table prefix
func (a CinemaHallsTable) WithPrefix(prefix string) *CinemaHallsTable {
	return newCinemaHallsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CinemaHallsTable with assigned table suffix
func (a CinemaHallsTable) WithSuffix(suffix string) *CinemaHallsTable {
	return newCinemaHallsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCinemaHallsTable(schemaName, tableName, alias string) *CinemaHallsTable {
	return &CinemaHallsTable{
		cinemaHallsTable: newCinemaHallsTableImpl(schemaName, tableName, alias),
		NEW:              newCinemaHallsTableImpl("", "new", ""),
	}
}

func newCinemaHallsTableImpl(schemaName, tableName, alias string) cinemaHallsTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		NameColumn      = mysql.StringColumn("name")
		CapacityColumn  = mysql.IntegerColumn("capacity")
		TheatreIDColumn = mysql.StringColumn("theatre_id")
		WidthColumn     = mysql.IntegerColumn("width")
		HeightColumn    = mysql.IntegerColumn("height")
		allColumns      = mysql.ColumnList{IDColumn, NameColumn, CapacityColumn, TheatreIDColumn, WidthColumn, HeightColumn}
		mutableColumns  = mysql.ColumnList{NameColumn, CapacityColumn, TheatreIDColumn, WidthColumn, HeightColumn}
	)

	return cinemaHallsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		Capacity:  CapacityColumn,
		TheatreID: TheatreIDColumn,
		Width:     WidthColumn,
		Height:    HeightColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
