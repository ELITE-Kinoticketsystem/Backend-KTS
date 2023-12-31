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

var ProducerPictures = newProducerPicturesTable("KinoTicketSystem", "producer_pictures", "")

type producerPicturesTable struct {
	mysql.Table

	// Columns
	ID         mysql.ColumnString
	ProducerID mysql.ColumnString
	PicURL     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ProducerPicturesTable struct {
	producerPicturesTable

	NEW producerPicturesTable
}

// AS creates new ProducerPicturesTable with assigned alias
func (a ProducerPicturesTable) AS(alias string) *ProducerPicturesTable {
	return newProducerPicturesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ProducerPicturesTable with assigned schema name
func (a ProducerPicturesTable) FromSchema(schemaName string) *ProducerPicturesTable {
	return newProducerPicturesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ProducerPicturesTable with assigned table prefix
func (a ProducerPicturesTable) WithPrefix(prefix string) *ProducerPicturesTable {
	return newProducerPicturesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ProducerPicturesTable with assigned table suffix
func (a ProducerPicturesTable) WithSuffix(suffix string) *ProducerPicturesTable {
	return newProducerPicturesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newProducerPicturesTable(schemaName, tableName, alias string) *ProducerPicturesTable {
	return &ProducerPicturesTable{
		producerPicturesTable: newProducerPicturesTableImpl(schemaName, tableName, alias),
		NEW:                   newProducerPicturesTableImpl("", "new", ""),
	}
}

func newProducerPicturesTableImpl(schemaName, tableName, alias string) producerPicturesTable {
	var (
		IDColumn         = mysql.StringColumn("id")
		ProducerIDColumn = mysql.StringColumn("producer_id")
		PicURLColumn     = mysql.StringColumn("pic_url")
		allColumns       = mysql.ColumnList{IDColumn, ProducerIDColumn, PicURLColumn}
		mutableColumns   = mysql.ColumnList{ProducerIDColumn, PicURLColumn}
	)

	return producerPicturesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		ProducerID: ProducerIDColumn,
		PicURL:     PicURLColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
