package tests

import (
	"database/sql"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
)

func GetTestDB(dbName string) *ent.Client {
	return rdb.GetTestDB(dbName)
}

func GetTestDBAndMigrate(dbName string) *ent.Client {
	return rdb.GetTestDBAndMigrate(dbName)
}

func CreateDatabase(db *sql.DB, dbName string) error {
	return rdb.CreateDatabase(db, dbName)
}

func IsAlreadyDatabase(err error) bool {
	return rdb.IsAlreadyDatabase(err)
}

func ErrorAsDBError(err error) error {
	return rdb.ErrorAsDBError(err)
}

func RunMigrationUp(dialect string, sqlDB *sql.DB, migratorArg *rdb.MigratorArgs) error {
	return rdb.RunMigrationUp(dialect, sqlDB, migratorArg)
}

func TruncateTable(sqlDB *sql.DB, name string) error {
	return rdb.TruncateTable(sqlDB, name)
}
