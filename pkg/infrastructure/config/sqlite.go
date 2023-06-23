package config

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/sethvargo/go-envconfig"
)

type SQLiteConfig struct {
	DSN                string `env:"DSN,default=file:short_url?mode=memory"`
	ConnMaxIdleTimeSec int    `env:"DB_CONN_MAX_IDLE_TIME_SEC,default=60"`
	ConnMaxLifetimeSec int    `env:"DB_CONN_MAX_LIFE_TIME_SEC,default=120"`
	MaxOpenConns       int    `env:"DB_MAX_OPEN_CONN,default=5"`
	MaxIdleConns       int    `env:"DB_MAX_IDLE_CONN,default=5"`
}

func (cfg *SQLiteConfig) Driver() config.Driver {
	return config.SQLite
}

func (cfg *SQLiteConfig) Build() (string, error) {
	return cfg.DSN, nil
}

func (cfg *SQLiteConfig) SQLDriver() string {
	return "sqlite3"
}

func (cfg *SQLiteConfig) Dialect() string {
	return "sqlite3"
}

func NewSQLiteConfig() *SQLiteConfig {
	cfg := &SQLiteConfig{}
	_ = envconfig.ProcessWith(context.Background(), cfg, envconfig.PrefixLookuper("SQLITE_", envconfig.OsLookuper()))
	return cfg
}
