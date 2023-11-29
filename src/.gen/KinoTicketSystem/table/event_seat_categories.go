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

var EventSeatCategories = newEventSeatCategoriesTable("KinoTicketSystem", "event_seat_categories", "")

type eventSeatCategoriesTable struct {
	mysql.Table

	// Columns
	EventID        mysql.ColumnString
	SeatCategoryID mysql.ColumnString
	Price          mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EventSeatCategoriesTable struct {
	eventSeatCategoriesTable

	NEW eventSeatCategoriesTable
}

// AS creates new EventSeatCategoriesTable with assigned alias
func (a EventSeatCategoriesTable) AS(alias string) *EventSeatCategoriesTable {
	return newEventSeatCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventSeatCategoriesTable with assigned schema name
func (a EventSeatCategoriesTable) FromSchema(schemaName string) *EventSeatCategoriesTable {
	return newEventSeatCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EventSeatCategoriesTable with assigned table prefix
func (a EventSeatCategoriesTable) WithPrefix(prefix string) *EventSeatCategoriesTable {
	return newEventSeatCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EventSeatCategoriesTable with assigned table suffix
func (a EventSeatCategoriesTable) WithSuffix(suffix string) *EventSeatCategoriesTable {
	return newEventSeatCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEventSeatCategoriesTable(schemaName, tableName, alias string) *EventSeatCategoriesTable {
	return &EventSeatCategoriesTable{
		eventSeatCategoriesTable: newEventSeatCategoriesTableImpl(schemaName, tableName, alias),
		NEW:                      newEventSeatCategoriesTableImpl("", "new", ""),
	}
}

func newEventSeatCategoriesTableImpl(schemaName, tableName, alias string) eventSeatCategoriesTable {
	var (
		EventIDColumn        = mysql.StringColumn("event_id")
		SeatCategoryIDColumn = mysql.StringColumn("seat_category_id")
		PriceColumn          = mysql.IntegerColumn("price")
		allColumns           = mysql.ColumnList{EventIDColumn, SeatCategoryIDColumn, PriceColumn}
		mutableColumns       = mysql.ColumnList{PriceColumn}
	)

	return eventSeatCategoriesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		EventID:        EventIDColumn,
		SeatCategoryID: SeatCategoryIDColumn,
		Price:          PriceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
