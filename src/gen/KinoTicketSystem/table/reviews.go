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

var Reviews = newReviewsTable("KinoTicketSystem", "reviews", "")

type reviewsTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString
	Rating    mysql.ColumnInteger
	Comment   mysql.ColumnString
	Datetime  mysql.ColumnTimestamp
	IsSpoiler mysql.ColumnBool
	UserID    mysql.ColumnString
	MovieID   mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ReviewsTable struct {
	reviewsTable

	NEW reviewsTable
}

// AS creates new ReviewsTable with assigned alias
func (a ReviewsTable) AS(alias string) *ReviewsTable {
	return newReviewsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ReviewsTable with assigned schema name
func (a ReviewsTable) FromSchema(schemaName string) *ReviewsTable {
	return newReviewsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ReviewsTable with assigned table prefix
func (a ReviewsTable) WithPrefix(prefix string) *ReviewsTable {
	return newReviewsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ReviewsTable with assigned table suffix
func (a ReviewsTable) WithSuffix(suffix string) *ReviewsTable {
	return newReviewsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newReviewsTable(schemaName, tableName, alias string) *ReviewsTable {
	return &ReviewsTable{
		reviewsTable: newReviewsTableImpl(schemaName, tableName, alias),
		NEW:          newReviewsTableImpl("", "new", ""),
	}
}

func newReviewsTableImpl(schemaName, tableName, alias string) reviewsTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		RatingColumn    = mysql.IntegerColumn("rating")
		CommentColumn   = mysql.StringColumn("comment")
		DatetimeColumn  = mysql.TimestampColumn("datetime")
		IsSpoilerColumn = mysql.BoolColumn("is_spoiler")
		UserIDColumn    = mysql.StringColumn("user_id")
		MovieIDColumn   = mysql.StringColumn("movie_id")
		allColumns      = mysql.ColumnList{IDColumn, RatingColumn, CommentColumn, DatetimeColumn, IsSpoilerColumn, UserIDColumn, MovieIDColumn}
		mutableColumns  = mysql.ColumnList{RatingColumn, CommentColumn, DatetimeColumn, IsSpoilerColumn, UserIDColumn, MovieIDColumn}
	)

	return reviewsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Rating:    RatingColumn,
		Comment:   CommentColumn,
		Datetime:  DatetimeColumn,
		IsSpoiler: IsSpoilerColumn,
		UserID:    UserIDColumn,
		MovieID:   MovieIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
