package rdb

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestIsErrAlreadyRecordExists(t *testing.T) {
	err := &mysql.MySQLError{
		Number: 1062,
	}
	assert.True(t, isErrAlreadyRecordExists(err))
	assert.False(t, isErrAlreadyRecordExists(errors.New("not a mysql error")))
}
