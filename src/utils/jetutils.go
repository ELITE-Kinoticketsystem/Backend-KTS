package utils

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

func MysqlUuid(uuid *uuid.UUID) mysql.StringExpression {
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}
