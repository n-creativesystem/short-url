package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/sethvargo/go-envconfig"
)

const (
	tlsMySQLConfigKey  = "custom"
	tlsMySQLSkipVerify = "skip-verify"
	tlsMySQLPreferred  = "preferred"
	tlsMySQLTrue       = "true"
	tlsMySQLFalse      = "false"
)

type MySQLConfig struct {
	DatabaseConfig
	DBTls               string   `env:"DB_TLS,default=preferred"`
	DBMaxAllowedPacket  int      `env:"DB_MAX_ALLOWED_PACKET"`
	DBInterpolateParams bool     `env:"DB_INTERPOLATE_PARAMS"`
	DBCollation         string   `env:"DB_COLLATION"`
	DBParams            []string `env:"DB_PARAMS,delimiter=|,default=transaction_isolation='READ-COMMITTED'|charset='utf8mb4'|sql_mode='TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY'"`
	MultiStatements     bool     `env:"DB_MULTI_STATEMENTS,default=false"` // migratorのみ有効とする
}

func (cfg *MySQLConfig) Driver() config.Driver {
	return config.MySQL
}

func (cfg *MySQLConfig) Build() (string, error) {
	if err := registerTLSConfig(cfg); err != nil {
		return "", err
	}
	return BuildMySQLConnectionString(cfg)
}

func (cfg *MySQLConfig) SQLDriver() string {
	return "mysql"
}

func (cfg *MySQLConfig) Dialect() string {
	return "mysql"
}

func BuildMySQLConnectionString(conf *MySQLConfig) (string, error) {
	mysqlCfg := mysql.NewConfig()

	if conf.DBHost == "" {
		return "", errors.New("db host is not set")
	}

	if conf.DBUser == "" {
		return "", errors.New("db user is not set")
	}

	mysqlCfg.Net = "tcp"
	mysqlCfg.Addr = fmt.Sprintf("%s:%d", conf.DBHost, conf.DBPort)

	mysqlCfg.DBName = conf.DBName
	mysqlCfg.User = conf.DBUser
	mysqlCfg.Passwd = conf.DBPass.UnmaskedString()

	mysqlCfg.ParseTime = true // goの場合は基本的にtrue必須
	switch conf.DBTls {
	case tlsMySQLTrue:
		if conf.DBTlsPath != "" {
			mysqlCfg.TLSConfig = tlsMySQLConfigKey
			break
		}
		mysqlCfg.TLSConfig = conf.DBTls
	case tlsMySQLFalse, tlsMySQLSkipVerify, tlsMySQLPreferred, "":
		mysqlCfg.TLSConfig = conf.DBTls
	default:
		return "", fmt.Errorf("unknown value for TLS: %v", conf.DBTls)
	}
	if conf.DBTz != "" {
		loc, err := time.LoadLocation(conf.DBTz)
		if err != nil {
			return "", err
		}
		mysqlCfg.Loc = loc
	}
	if conf.DBMaxAllowedPacket != 0 {
		mysqlCfg.MaxAllowedPacket = conf.DBMaxAllowedPacket
	}
	mysqlCfg.InterpolateParams = conf.DBInterpolateParams
	if conf.DBCollation != "" {
		mysqlCfg.Collation = conf.DBCollation
	}
	mysqlCfg.MultiStatements = conf.MultiStatements
	if len(conf.DBParams) > 0 {
		mysqlCfg.Params = make(map[string]string)
	}
	for _, param := range conf.DBParams {
		key, value, found := strings.Cut(param, "=")
		if found {
			mysqlCfg.Params[key] = value
		}
	}
	ret := mysqlCfg.FormatDSN()
	return ret, nil
}

func registerTLSConfig(conf *MySQLConfig) error {
	if conf.DBTls != tlsMySQLTrue {
		return nil
	}

	if conf.DBTlsPath == "" {
		return errors.New("must set db_tls_path when db_tls is true")
	}

	caCertPool := x509.NewCertPool()
	cert, err := os.ReadFile(conf.DBTlsPath)
	if err != nil {
		return err
	}

	if ok := caCertPool.AppendCertsFromPEM(cert); !ok {
		return errors.New("failed to append pem to ssl cert pool")
	}
	tlsConfig := tls.Config{
		MinVersion:         tls.VersionTLS12,
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}
	// NOTE: 第一引数に `custom` を渡しているので `RegisterTLSConfig` のコードが変わらない限りエラーになることはない
	if err := mysql.RegisterTLSConfig(tlsMySQLConfigKey, &tlsConfig); err != nil {
		return err
	}

	return nil
}

func NewMySQLConfig() *MySQLConfig {
	cfg := &MySQLConfig{}
	_ = envconfig.ProcessWith(context.Background(), cfg, envconfig.PrefixLookuper("MYSQL_", envconfig.OsLookuper()))
	return cfg
}
