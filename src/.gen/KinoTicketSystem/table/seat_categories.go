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

var SeatCategories = newSeatCategoriesTable("KinoTicketSystem", "seat_categories", "")

type seatCategoriesTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnString
	CategoryName mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type SeatCategoriesTable struct {
	seatCategoriesTable

	NEW seatCategoriesTable
}

// AS creates new SeatCategoriesTable with assigned alias
func (a SeatCategoriesTable) AS(alias string) *SeatCategoriesTable {
	return newSeatCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SeatCategoriesTable with assigned schema name
func (a SeatCategoriesTable) FromSchema(schemaName string) *SeatCategoriesTable {
	return newSeatCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SeatCategoriesTable with assigned table prefix
func (a SeatCategoriesTable) WithPrefix(prefix string) *SeatCategoriesTable {
	return newSeatCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SeatCategoriesTable with assigned table suffix
func (a SeatCategoriesTable) WithSuffix(suffix string) *SeatCategoriesTable {
	return newSeatCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSeatCategoriesTable(schemaName, tableName, alias string) *SeatCategoriesTable {
	return &SeatCategoriesTable{
		seatCategoriesTable: newSeatCategoriesTableImpl(schemaName, tableName, alias),
		NEW:                 newSeatCategoriesTableImpl("", "new", ""),
	}
}

func newSeatCategoriesTableImpl(schemaName, tableName, alias string) seatCategoriesTable {
	var (
		IDColumn           = mysql.StringColumn("id")
		CategoryNameColumn = mysql.StringColumn("category_name")
		allColumns         = mysql.ColumnList{IDColumn, CategoryNameColumn}
		mutableColumns     = mysql.ColumnList{CategoryNameColumn}
	)

	return seatCategoriesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CategoryName: CategoryNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
