package rdb

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func isErrAlreadyRecordExists(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}
