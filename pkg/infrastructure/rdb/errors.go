package rdb

import (
	"database/sql"
	"errors"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
)

func IsNotFoundRecord(err error) bool {
	if ent.IsNotFound(err) {
		return true
	}
	return errors.Is(err, sql.ErrNoRows)
}
