package utils

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

func MysqlUuid(uuid *uuid.UUID) (mysql.StringExpression, error) {
	binary_id, err := uuid.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return mysql.String(string(binary_id)), nil
}
