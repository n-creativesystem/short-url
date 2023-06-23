package interfaces

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
)

type DriverName int

const (
	RDB DriverName = iota
	DynamoDB
)

type PingContextExecutor interface {
	PingContext(ctx context.Context) error
}

type PingExecutor interface {
	Ping() error
}

func GetPing(name DriverName) PingContextExecutor {
	switch name {
	case RDB:
		return rdb.GetDB()
	}
	return nil
}
