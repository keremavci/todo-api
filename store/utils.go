package store

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	e "errors"

)

const (
	MYSQL_DUP_TABLE_ERROR_CODE = uint16(1050) // see https://dev.mysql.com/doc/mysql-errors/5.7/en/server-error-reference.html#error_er_table_exists_error
	DB_PING_ATTEMPTS           = 6
	DB_PING_TIMEOUT_SECS       = 10
	EXIT_CREATE_TABLE                = 100
	EXIT_DB_OPEN                     = 101
	EXIT_PING                        = 102
)

func IsDuplicate(err error) bool {
	var mysqlErr *mysql.MySQLError
	switch {
	case e.As(errors.Cause(err), &mysqlErr):
		if mysqlErr.Number == MYSQL_DUP_TABLE_ERROR_CODE {
			return true
		}
	}

	return false
}