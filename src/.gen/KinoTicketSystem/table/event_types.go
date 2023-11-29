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

var EventTypes = newEventTypesTable("KinoTicketSystem", "event_types", "")

type eventTypesTable struct {
	mysql.Table

	// Columns
	ID       mysql.ColumnString
	Typename mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type EventTypesTable struct {
	eventTypesTable

	NEW eventTypesTable
}

// AS creates new EventTypesTable with assigned alias
func (a EventTypesTable) AS(alias string) *EventTypesTable {
	return newEventTypesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EventTypesTable with assigned schema name
func (a EventTypesTable) FromSchema(schemaName string) *EventTypesTable {
	return newEventTypesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EventTypesTable with assigned table prefix
func (a EventTypesTable) WithPrefix(prefix string) *EventTypesTable {
	return newEventTypesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EventTypesTable with assigned table suffix
func (a EventTypesTable) WithSuffix(suffix string) *EventTypesTable {
	return newEventTypesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEventTypesTable(schemaName, tableName, alias string) *EventTypesTable {
	return &EventTypesTable{
		eventTypesTable: newEventTypesTableImpl(schemaName, tableName, alias),
		NEW:             newEventTypesTableImpl("", "new", ""),
	}
}

func newEventTypesTableImpl(schemaName, tableName, alias string) eventTypesTable {
	var (
		IDColumn       = mysql.StringColumn("id")
		TypenameColumn = mysql.StringColumn("typename")
		allColumns     = mysql.ColumnList{IDColumn, TypenameColumn}
		mutableColumns = mysql.ColumnList{TypenameColumn}
	)

	return eventTypesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Typename: TypenameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
