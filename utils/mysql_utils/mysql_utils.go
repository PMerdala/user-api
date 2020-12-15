package mysql_utils

import (
	"fmt"
	"github.com/PMerdala/users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	mysqlIndexUniqueEmailErrorNumber    = 1062
	mysqlNotRowsError                   = "sql: no rows in result set"
	errorParsingDatabaseResponseMessage = "error parsing database response: %s"
	unexpectDatabaseErrorMessage        = "error on try execute database statement: %s"
)

func ParserError(err error, msg string) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), mysqlNotRowsError) {
			return errors.NewNotFoundError(msg)
		}
		return errors.NewInternalServerError(fmt.Sprintf(errorParsingDatabaseResponseMessage, err.Error()))
	}
	switch sqlErr.Number {
	case mysqlIndexUniqueEmailErrorNumber:
		return errors.NewBadRequestError(msg)
	}
	return errors.NewInternalServerError(fmt.Sprintf(unexpectDatabaseErrorMessage, err.Error()))
}
