package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/n-creativesystem/short-url/pkg/infrastructure/config"
	"github.com/redis/go-redis/v9"
)

func New(ctx context.Context, cfg *config.Redis) (Cmdable, error) {
	if cfg.Sentinel && cfg.Cluster {
		return nil, errors.New("redis sentinelとredis clusterのどちらかのみ指定して下さい")
	}
	var (
		c   Cmdable
		err error
	)
	if cfg.Sentinel {
		c, err = buildSentinel(cfg)
		if err != nil {
			return nil, err
		}
	}
	if cfg.Cluster {
		c, err = buildCluster(cfg)
		if err != nil {
			return nil, err
		}
	}

	if c == nil {
		c, err = buildStandalone(cfg)
		if err != nil {
			return nil, err
		}
	}

	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return c, nil
}

func buildSentinel(cfg *config.Redis) (Cmdable, error) {
	addrs, opt, err := parseRedisURLs(cfg.Addresses)
	if err != nil {
		return nil, err
	}
	return redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs:    addrs,
		SentinelPassword: cfg.Password,
		MasterName:       cfg.MasterName,
		TLSConfig:        opt.TLSConfig,
	}), nil
}

func buildCluster(cfg *config.Redis) (Cmdable, error) {
	addrs, opt, err := parseRedisURLs(cfg.Addresses)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:     addrs,
		Password:  cfg.Password,
		TLSConfig: opt.TLSConfig,
	}), nil
}

func buildStandalone(cfg *config.Redis) (Cmdable, error) {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	}), nil
}

func parseRedisURLs(urls []string) ([]string, *redis.Options, error) {
	addrs := []string{}
	var redisOptions *redis.Options
	for _, u := range urls {
		parsedURL, err := redis.ParseURL(u)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse redis url: %v", err)
		}
		addrs = append(addrs, parsedURL.Addr)
		redisOptions = parsedURL
	}
	return addrs, redisOptions, nil
}

type Cmdable interface {
	redis.Cmdable
	Close() error
}
