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

var Producers = newProducersTable("KinoTicketSystem", "producers", "")

type producersTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnString
	Name        mysql.ColumnString
	Birthdate   mysql.ColumnDate
	Description mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ProducersTable struct {
	producersTable

	NEW producersTable
}

// AS creates new ProducersTable with assigned alias
func (a ProducersTable) AS(alias string) *ProducersTable {
	return newProducersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ProducersTable with assigned schema name
func (a ProducersTable) FromSchema(schemaName string) *ProducersTable {
	return newProducersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ProducersTable with assigned table prefix
func (a ProducersTable) WithPrefix(prefix string) *ProducersTable {
	return newProducersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ProducersTable with assigned table suffix
func (a ProducersTable) WithSuffix(suffix string) *ProducersTable {
	return newProducersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newProducersTable(schemaName, tableName, alias string) *ProducersTable {
	return &ProducersTable{
		producersTable: newProducersTableImpl(schemaName, tableName, alias),
		NEW:            newProducersTableImpl("", "new", ""),
	}
}

func newProducersTableImpl(schemaName, tableName, alias string) producersTable {
	var (
		IDColumn          = mysql.StringColumn("id")
		NameColumn        = mysql.StringColumn("name")
		BirthdateColumn   = mysql.DateColumn("birthdate")
		DescriptionColumn = mysql.StringColumn("description")
		allColumns        = mysql.ColumnList{IDColumn, NameColumn, BirthdateColumn, DescriptionColumn}
		mutableColumns    = mysql.ColumnList{NameColumn, BirthdateColumn, DescriptionColumn}
	)

	return producersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Birthdate:   BirthdateColumn,
		Description: DescriptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
