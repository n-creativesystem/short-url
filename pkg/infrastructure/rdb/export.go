package rdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/n-creativesystem/short-url/db"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	infra_config "github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	_ "github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/driver"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb/ent"
)

func getDBName(dbName string) string {
	if dbName == "" {
		dbName = "shorturl_test"
	}
	return dbName
}

func GetMigrateDBWithConfig(dbName string, cfg config.DBConfig) *ent.Client {
	return GetMigrateClientWithConfig(dbName, cfg).Client
}

func GetMigrateClientWithConfig(dbName string, cfg config.DBConfig) *Client {
	dbName = getDBName(dbName)
	// まずは設定されている環境変数で繋ぐ
	db, err := NewDB(cfg)
	if err != nil {
		panic(err)
	}
	if cfg.Driver() != config.SQLite {
		_ = DropDatabase(db.db, dbName)
		if err := CreateDatabase(db.db, dbName); err != nil && !IsAlreadyDatabase(err) {
			panic(err)
		}
	}
	_ = db.Close()
	cfg = infra_config.NewDBConfig(infra_config.WithDBName(dbName), infra_config.WithMigration())
	db, err = NewDB(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

func GetTestDB(dbName string) *ent.Client {
	dbName = getDBName(dbName)
	cfg := infra_config.NewDBConfig(infra_config.WithDBName(dbName))
	db, err := NewDB(cfg)
	if err != nil {
		panic(err)
	}
	return db.Client
}

func GetTestDBAndMigrate(dbName string) *ent.Client {
	// まずは設定されている環境変数で繋ぐ
	cfg := infra_config.NewDBConfig(infra_config.WithMigration())
	conn := GetMigrateDBWithConfig(dbName, cfg)
	if err := conn.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	if cfg.Driver() == config.SQLite {
		return conn
	}
	_ = conn.Close()
	return GetTestDB(dbName)
}

func RunMigrationUp(dialect string, sqlDB *sql.DB, migratorArg *MigratorArgs) error {
	if migratorArg == nil {
		migratorArg = &MigratorArgs{
			RetryCount:   2,
			RetryWait:    1000,
			AllowMissing: false,
			Dir:          db.MustGetMigrationsDirectory(),
		}
	}
	return MigrationWithDB(context.Background(), newClient(dialect, sqlDB), []string{"up"}, *migratorArg)
}

func TruncateTable(sqlDB *sql.DB, name string) error {
	_, err := sqlDB.Exec(fmt.Sprintf("truncate table %s", name))
	return err
}

func DropDatabase(db *sql.DB, dbName string) error {
	sql := fmt.Sprintf("drop database %s", dbName)
	_, err := db.Exec(sql)
	return err
}

func CreateDatabase(db *sql.DB, dbName string) error {
	var sql string
	switch config.GetDriver() {
	case config.MySQL:
		sql = "CREATE DATABASE `%s` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;"
	case config.PostgreSQL:
		sql = "create database %s"
	}
	sql = fmt.Sprintf(sql, dbName)
	_, err := db.Exec(sql)
	return err
}

func IsAlreadyDatabase(err error) bool {
	err = ErrorAsDBError(err)
	if e, ok := err.(*pgconn.PgError); e != nil && ok {
		return e.Code == "42P04"
	}
	if e, ok := err.(*mysql.MySQLError); e != nil && ok {
		return e.Number == 1007
	}
	return false
}

func ErrorAsDBError(err error) error {
	var postgresErr *pgconn.PgError
	if errors.As(err, &postgresErr) {
		return postgresErr
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr
	}
	return nil
}
