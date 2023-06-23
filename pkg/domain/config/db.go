package config

type DBConfig interface {
	Driver() Driver
	Build() (string, error)
	SQLDriver() string
	Dialect() string
}
