package utils

import (
	"database/sql"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

func MysqlUuid(uuid *uuid.UUID) mysql.StringExpression {
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}

func MysqlUuidOrNil(uuid *uuid.UUID) mysql.Expression {
	if uuid == nil {
		return mysql.NULL
	}
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}

func MysqlTime(time *time.Time) mysql.TimeExpression {
	return mysql.TimeT(*time)
}

func MySqlString(str string) mysql.StringExpression {
	return mysql.String(str)
}

func MySqlStringPtr(str *string) mysql.StringExpression {
	if str == nil || *str == "" {
		return nil
	}
	return mysql.String(*str)
}

func MysqlTimeNow() mysql.TimestampExpression {
	return mysql.NOW()
}

func ExcecuteInsertStatement(stmt mysql.InsertStatement, dbConnection *sql.DB) *models.KTSError {
	result, err := stmt.Exec(dbConnection)

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected != 1 {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}
