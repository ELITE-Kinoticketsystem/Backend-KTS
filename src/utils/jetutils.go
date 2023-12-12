package utils

import (
	"database/sql"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

func MysqlUuid(uuid *uuid.UUID) mysql.StringExpression {
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}

func MySqlString(str string) mysql.StringExpression {
	return mysql.String(str)
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
