package interfaces

import (
	"context"

	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
)

type PingContextExecutor interface {
	PingContext(ctx context.Context) error
}

type PingExecutor interface {
	Ping() error
}

func GetPing() PingContextExecutor {
	driver := config.GetDriver()
	switch driver {
	case config.MySQL, config.PostgreSQL, config.SQLite:
		return rdb.GetDB()
	default:
		return noopPing{}
	}
}

type noopPing struct{}

func (noopPing) PingContext(ctx context.Context) error { return nil }
