package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/n-creativesystem/short-url/fixtures"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials"
	"github.com/n-creativesystem/short-url/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testSetDefaultEnv(t *testing.T, prefix string) {
	envValue := `DB_HOST=hostname
DB_PORT=3306
DB_USER=user
DB_PASS=pass
DB_NAME=dbName
DB_TLS=false
DB_TLS_PATH=
DB_TZ=Asia/Tokyo
DB_MAX_ALLOWED_PACKET=100
DB_INTERPOLATE_PARAMS=true
DB_COLLATION=utf8mb4
DB_CONN_MAX_IDLE_TIME_SEC=100
DB_CONN_MAX_LIFE_TIME_SEC=100
DB_MAX_OPEN_CONN=10
DB_MAX_IDLE_CONN=10
DB_MULTI_STATEMENTS=false`
	envValues := strings.Split(envValue, "\n")
	buf := new(bytes.Buffer)
	for _, env := range envValues {
		if env == "" {
			continue
		}
		buf.WriteString(fmt.Sprintf("%s%s\n", prefix, env))
	}

	envMap, err := godotenv.Parse(buf)
	require.NoError(t, err)
	for key, value := range envMap {
		_ = os.Setenv(key, value)
	}
}

func TestMySQLConfig(t *testing.T) {
	tearDown := tests.EnvSetup()
	defer tearDown()

	testSetDefaultEnv(t, "MYSQL_")

	expectConfig := &MySQLConfig{
		DatabaseConfig: DatabaseConfig{
			DBHost:             "hostname",
			DBPort:             3306,
			DBUser:             "user",
			DBPass:             credentials.NewMaskedString("pass"),
			DBName:             "dbName",
			ConnMaxIdleTimeSec: 100,
			ConnMaxLifetimeSec: 100,
			MaxOpenConns:       10,
			MaxIdleConns:       10,
			DBTz:               "Asia/Tokyo",
			DBTlsPath:          "",
		},
		DBTls:               tlsMySQLFalse,
		DBMaxAllowedPacket:  100,
		DBInterpolateParams: true,
		DBCollation:         "utf8mb4",
		MultiStatements:     false,
		DBParams:            []string{"transaction_isolation='READ-COMMITTED'", "charset='utf8mb4'", "sql_mode='TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY'"},
	}

	cfg := NewMySQLConfig()
	assert.Equal(t, expectConfig, cfg)
	err := registerTLSConfig(cfg)
	require.NoError(t, err)
	conn, err := BuildMySQLConnectionString(cfg)
	require.NoError(t, err)
	expectConn := `user:pass@tcp(hostname:3306)/dbName?collation=utf8mb4&interpolateParams=true&loc=Asia%2FTokyo&parseTime=true&tls=false&maxAllowedPacket=100&charset=%27utf8mb4%27&sql_mode=%27TRADITIONAL%2CNO_AUTO_VALUE_ON_ZERO%2CONLY_FULL_GROUP_BY%27&transaction_isolation=%27READ-COMMITTED%27`
	assert.Equal(t, expectConn, conn)
}

func TestMySQLConfigError(t *testing.T) {
	checker := func(fn func() string) {
		tearDown := tests.EnvSetup()
		defer tearDown()
		expectErrMsg := fn()
		cfg := NewMySQLConfig()
		conn, err := BuildMySQLConnectionString(cfg)
		assert.Error(t, err)
		assert.Empty(t, conn)
		assert.Equal(t, expectErrMsg, err.Error())
	}
	setEnv := func(key, value string) {
		os.Setenv(fmt.Sprintf("%s%s", "MYSQL_", key), value)
	}
	checker(func() string {
		setEnv("DB_HOST", "")
		return "db host is not set"
	})
	checker(func() string {
		setEnv("DB_USER", "")
		return "db user is not set"
	})
	checker(func() string {
		setEnv("DB_TLS", "other")
		return "unknown value for TLS: other"
	})
	checker(func() string {
		setEnv("DB_TZ", "other")
		return "unknown time zone other"
	})
}

func TestRegisterTLSConfig(t *testing.T) {
	setEnv := func(key, value string) {
		os.Setenv(fmt.Sprintf("%s%s", "MYSQL_", key), value)
	}
	tearDown := tests.EnvSetup()
	defer tearDown()
	testSetDefaultEnv(t, "MYSQL_")
	setEnv("DB_TLS", "true")
	var (
		cfg *MySQLConfig
		err error
	)
	cfg = NewMySQLConfig()
	err = registerTLSConfig(cfg)
	assert.Error(t, err)
	assert.Equal(t, "must set db_tls_path when db_tls is true", err.Error())

	setEnv("DB_TLS_PATH", "true")
	dir, err := fixtures.GetDirectory()
	require.NoError(t, err)
	pemFile := filepath.Join(dir, "notfound.crt")
	setEnv("DB_TLS_PATH", pemFile)
	cfg = NewMySQLConfig()
	err = registerTLSConfig(cfg)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", pemFile), err.Error())

	pemFile = filepath.Join(dir, "ca.key")
	setEnv("DB_TLS_PATH", pemFile)
	cfg = NewMySQLConfig()
	err = registerTLSConfig(cfg)
	assert.Error(t, err)
	assert.Equal(t, "failed to append pem to ssl cert pool", err.Error())

	pemFile = filepath.Join(dir, "ca.crt")
	setEnv("DB_TLS_PATH", pemFile)
	cfg = NewMySQLConfig()
	err = registerTLSConfig(cfg)
	require.NoError(t, err)
	conn, err := BuildMySQLConnectionString(cfg)
	require.NoError(t, err)
	expectedConn := `user:pass@tcp(hostname:3306)/dbName?collation=utf8mb4&interpolateParams=true&loc=Asia%2FTokyo&parseTime=true&tls=custom&maxAllowedPacket=100&charset=%27utf8mb4%27&sql_mode=%27TRADITIONAL%2CNO_AUTO_VALUE_ON_ZERO%2CONLY_FULL_GROUP_BY%27&transaction_isolation=%27READ-COMMITTED%27`
	assert.Equal(t, expectedConn, conn)
}
