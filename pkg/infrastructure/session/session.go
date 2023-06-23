package session

import (
	"context"
	"time"

	"github.com/alexedwards/scs/boltstore"
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/redis"
	"go.etcd.io/bbolt"
)

func New(ctx context.Context, cfg *config.Session) (scs.Store, error) {
	switch {
	case cfg.Redis:
		cmd, err := redis.New(ctx, &cfg.RedisConfig)
		if err != nil {
			return nil, err
		}
		return newRedis(cmd), nil
	case cfg.Bolt:
		db, err := bbolt.Open(cfg.BoltConfig.File, 0600, nil)
		if err != nil {
			return nil, err
		}
		return boltstore.NewWithCleanupInterval(db, cfg.BoltConfig.Interval), nil
	default:
		return memstore.NewWithCleanupInterval(20 * time.Second), nil
	}
}
