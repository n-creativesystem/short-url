package session

import (
	"context"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	cmd    redis.Cmdable
	prefix string
}

var (
	_ scs.Store    = (*RedisStore)(nil)
	_ scs.CtxStore = (*RedisStore)(nil)
)

func newRedis(cmd redis.Cmdable) *RedisStore {
	return newRedisWithPrefix(cmd, "scs::session:")
}

func newRedisWithPrefix(cmd redis.Cmdable, prefix string) *RedisStore {
	return &RedisStore{
		cmd:    cmd,
		prefix: prefix,
	}
}

func (r *RedisStore) key(token string) string {
	return r.prefix + token
}

func (r *RedisStore) DeleteCtx(ctx context.Context, token string) error {
	return r.cmd.Del(ctx, r.key(token)).Err()
}

func (r *RedisStore) FindCtx(ctx context.Context, token string) ([]byte, bool, error) {
	value, err := r.cmd.Get(ctx, r.key(token)).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	} else {
		return value, true, nil
	}
}

func (r *RedisStore) CommitCtx(ctx context.Context, token string, b []byte, expiry time.Time) error {
	err := r.cmd.Set(ctx, r.key(token), b, 0).Err()
	if err != nil {
		return err
	}
	return r.cmd.PExpireAt(ctx, r.key(token), expiry).Err()
}

func (r *RedisStore) Delete(token string) error {
	return r.DeleteCtx(context.Background(), token)
}

func (r *RedisStore) Find(token string) ([]byte, bool, error) {
	return r.FindCtx(context.Background(), token)
}

func (r *RedisStore) Commit(token string, b []byte, expiry time.Time) error {
	return r.CommitCtx(context.Background(), token, b, expiry)
}

func (r *RedisStore) All() (map[string][]byte, error) {
	keys, err := r.cmd.Keys(context.Background(), r.key("*")).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	sessions := make(map[string][]byte)

	for _, key := range keys {
		token := key[len(r.prefix):]

		data, exists, err := r.Find(token)
		if err == redis.Nil {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		if exists {
			sessions[token] = data
		}
	}

	return sessions, nil
}
