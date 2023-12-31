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

var Tickets = newTicketsTable("KinoTicketSystem", "tickets", "")

type ticketsTable struct {
	mysql.Table

	// Columns
	ID              mysql.ColumnString
	Validated       mysql.ColumnBool
	Price           mysql.ColumnInteger
	PriceCategoryID mysql.ColumnString
	OrderID         mysql.ColumnString
	EventSeatID     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TicketsTable struct {
	ticketsTable

	NEW ticketsTable
}

// AS creates new TicketsTable with assigned alias
func (a TicketsTable) AS(alias string) *TicketsTable {
	return newTicketsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TicketsTable with assigned schema name
func (a TicketsTable) FromSchema(schemaName string) *TicketsTable {
	return newTicketsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TicketsTable with assigned table prefix
func (a TicketsTable) WithPrefix(prefix string) *TicketsTable {
	return newTicketsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TicketsTable with assigned table suffix
func (a TicketsTable) WithSuffix(suffix string) *TicketsTable {
	return newTicketsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTicketsTable(schemaName, tableName, alias string) *TicketsTable {
	return &TicketsTable{
		ticketsTable: newTicketsTableImpl(schemaName, tableName, alias),
		NEW:          newTicketsTableImpl("", "new", ""),
	}
}

func newTicketsTableImpl(schemaName, tableName, alias string) ticketsTable {
	var (
		IDColumn              = mysql.StringColumn("id")
		ValidatedColumn       = mysql.BoolColumn("validated")
		PriceColumn           = mysql.IntegerColumn("price")
		PriceCategoryIDColumn = mysql.StringColumn("price_category_id")
		OrderIDColumn         = mysql.StringColumn("order_id")
		EventSeatIDColumn     = mysql.StringColumn("event_seat_id")
		allColumns            = mysql.ColumnList{IDColumn, ValidatedColumn, PriceColumn, PriceCategoryIDColumn, OrderIDColumn, EventSeatIDColumn}
		mutableColumns        = mysql.ColumnList{ValidatedColumn, PriceColumn, PriceCategoryIDColumn, OrderIDColumn, EventSeatIDColumn}
	)

	return ticketsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Validated:       ValidatedColumn,
		Price:           PriceColumn,
		PriceCategoryID: PriceCategoryIDColumn,
		OrderID:         OrderIDColumn,
		EventSeatID:     EventSeatIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
