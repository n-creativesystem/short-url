package config

import (
	"fmt"
	"strings"

	"github.com/n-creativesystem/short-url/pkg/utils"
)

type Driver int

func (d Driver) String() string {
	switch d {
	case MySQL:
		return "mysql"
	case PostgreSQL:
		return "postgres"
	case DynamoDB:
		return "dynamodb"
	case SQLite:
		return "sqlite3"
	}
	return ""
}

const (
	None Driver = iota
	MySQL
	PostgreSQL
	DynamoDB
	SQLite
)

var (
	driver Driver
	RDB    = []Driver{MySQL, PostgreSQL, SQLite}
)

func SetDriver(d Driver) {
	driver = d
}

func GetDriver() Driver {
	return driver
}

func ConvertDriverFromString(value string) Driver {
	switch strings.ToLower(value) {
	case "mysql":
		return MySQL
	case "postgres", "pgx", "pg", "psql":
		return PostgreSQL
	case "sqlite", "sqlite3":
		return SQLite
	case "dynamo", "dynamodb":
		return DynamoDB
	default:
		panic(driverError(utils.StringerFunc(func() string { return value })))
	}
}

func driverError(value fmt.Stringer) error {
	return fmt.Errorf("No support driver: %s", value)
}

var ErrNoSupportDriver = func() error { return fmt.Errorf("No support driver: %s", GetDriver()) }
