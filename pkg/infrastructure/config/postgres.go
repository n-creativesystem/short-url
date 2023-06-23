package config

import (
	"context"
	"errors"
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/sethvargo/go-envconfig"
)

type PostgreSQLConfig struct {
	DatabaseConfig
	DBTls     string   `env:"DB_TLS,default=prefer"`
	DBTlsCert string   `env:"DB_TLS_CERT"`
	DBTlsKey  string   `env:"DB_TLS_KEY"`
	DBParams  []string `env:"DB_PARAMS,delimiter=|"`
}

func (cfg *PostgreSQLConfig) Driver() config.Driver {
	return config.PostgreSQL
}

func (cfg *PostgreSQLConfig) Build() (string, error) {
	return BuildPostgreSQLConnectionString(cfg)
}

func (cfg *PostgreSQLConfig) SQLDriver() string {
	return "pgx"
}

func (cfg *PostgreSQLConfig) Dialect() string {
	return "postgres"
}

func BuildPostgreSQLConnectionString(conf *PostgreSQLConfig) (string, error) {
	if conf.DBHost == "" {
		return "", errors.New("db host is not set")
	}

	if conf.DBUser == "" {
		return "", errors.New("db user is not set")
	}

	switch conf.DBTls {
	case "", tlsMySQLFalse:
		conf.DBTls = "disabled"
	case tlsMySQLPreferred:
		conf.DBTls = "prefer"
	case tlsMySQLSkipVerify, tlsMySQLTrue:
		conf.DBTls = "require"
	}

	mp := map[string]string{
		"host":        conf.DBHost,
		"port":        fmt.Sprintf("%d", conf.DBPort),
		"user":        conf.DBUser,
		"password":    conf.DBPass.UnmaskedString(),
		"dbname":      conf.DBName,
		"sslmode":     conf.DBTls,
		"sslrootcert": conf.DBTlsPath,
		"sslcert":     conf.DBTlsCert,
		"sslkey":      conf.DBTlsKey,
		"timezone":    conf.DBTz,
	}
	for _, param := range conf.DBParams {
		key, value, found := strings.Cut(param, "=")
		if found {
			mp[key] = value
		}
	}
	values := make([]string, 0, len(mp))
	for k, v := range mp {
		if v == "" {
			continue
		}
		values = append(values, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(values, " "), nil
}

func NewPostgreSQLConfig() *PostgreSQLConfig {
	cfg := &PostgreSQLConfig{}
	_ = envconfig.ProcessWith(context.Background(), cfg, envconfig.PrefixLookuper("PG_", envconfig.OsLookuper()))
	return cfg
}
