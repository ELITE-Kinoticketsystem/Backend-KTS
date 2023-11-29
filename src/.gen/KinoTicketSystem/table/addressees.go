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

var Addressees = newAddresseesTable("KinoTicketSystem", "addressees", "")

type addresseesTable struct {
	mysql.Table

	// Columns
	ID       mysql.ColumnString
	Street   mysql.ColumnString
	StreetNr mysql.ColumnString
	Zipcode  mysql.ColumnString
	City     mysql.ColumnString
	Country  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AddresseesTable struct {
	addresseesTable

	NEW addresseesTable
}

// AS creates new AddresseesTable with assigned alias
func (a AddresseesTable) AS(alias string) *AddresseesTable {
	return newAddresseesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AddresseesTable with assigned schema name
func (a AddresseesTable) FromSchema(schemaName string) *AddresseesTable {
	return newAddresseesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AddresseesTable with assigned table prefix
func (a AddresseesTable) WithPrefix(prefix string) *AddresseesTable {
	return newAddresseesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AddresseesTable with assigned table suffix
func (a AddresseesTable) WithSuffix(suffix string) *AddresseesTable {
	return newAddresseesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAddresseesTable(schemaName, tableName, alias string) *AddresseesTable {
	return &AddresseesTable{
		addresseesTable: newAddresseesTableImpl(schemaName, tableName, alias),
		NEW:             newAddresseesTableImpl("", "new", ""),
	}
}

func newAddresseesTableImpl(schemaName, tableName, alias string) addresseesTable {
	var (
		IDColumn       = mysql.StringColumn("id")
		StreetColumn   = mysql.StringColumn("street")
		StreetNrColumn = mysql.StringColumn("street_nr")
		ZipcodeColumn  = mysql.StringColumn("zipcode")
		CityColumn     = mysql.StringColumn("city")
		CountryColumn  = mysql.StringColumn("country")
		allColumns     = mysql.ColumnList{IDColumn, StreetColumn, StreetNrColumn, ZipcodeColumn, CityColumn, CountryColumn}
		mutableColumns = mysql.ColumnList{StreetColumn, StreetNrColumn, ZipcodeColumn, CityColumn, CountryColumn}
	)

	return addresseesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Street:   StreetColumn,
		StreetNr: StreetNrColumn,
		Zipcode:  ZipcodeColumn,
		City:     CityColumn,
		Country:  CountryColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
