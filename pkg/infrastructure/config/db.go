package config

import (
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
)

type DatabaseConfig struct {
	DBHost             string                    `env:"DB_HOST,required"`
	DBPort             int                       `env:"DB_PORT,required"`
	DBUser             string                    `env:"DB_USER,required"`
	DBPass             *credentials.MaskedString `env:"DB_PASS,required"`
	DBName             string                    `env:"DB_NAME,required"`
	DBTlsPath          string                    `env:"DB_TLS_PATH"`
	DBTz               string                    `env:"DB_TZ,default=Asia/Tokyo"`
	ConnMaxIdleTimeSec int                       `env:"DB_CONN_MAX_IDLE_TIME_SEC,default=60"`
	ConnMaxLifetimeSec int                       `env:"DB_CONN_MAX_LIFE_TIME_SEC,default=120"`
	MaxOpenConns       int                       `env:"DB_MAX_OPEN_CONN,default=5"`
	MaxIdleConns       int                       `env:"DB_MAX_IDLE_CONN,default=5"`
}

type dbOption struct {
	migrate bool
	dbName  string
}

type DBOption interface {
	apply(*dbOption)
}

type dbOptionFn func(*dbOption)

func (fn dbOptionFn) apply(o *dbOption) {
	fn(o)
}

func WithMigration() DBOption {
	return dbOptionFn(func(do *dbOption) {
		do.migrate = true
	})
}

func WithDBName(dbName string) DBOption {
	return dbOptionFn(func(do *dbOption) {
		do.dbName = dbName
	})
}

func NewDBConfig(opts ...DBOption) config.DBConfig {
	opt := &dbOption{}
	for _, o := range opts {
		o.apply(opt)
	}
	switch config.GetDriver() {
	case config.MySQL:
		cfg := NewMySQLConfig()
		if opt.migrate {
			// migratorは常にMultiStatements許可とする
			cfg.MultiStatements = true
		}
		if opt.dbName != "" {
			cfg.DBName = opt.dbName
		}
		return cfg
	case config.PostgreSQL:
		cfg := NewPostgreSQLConfig()
		if opt.dbName != "" {
			cfg.DBName = opt.dbName
		}
		return cfg
	case config.SQLite:
		cfg := NewSQLiteConfig()
		return cfg
	default:
		panic(config.ErrNoSupportDriver())
	}
}
