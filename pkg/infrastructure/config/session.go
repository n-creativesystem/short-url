package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Session struct {
	Redis       bool  `env:"REDIS"`
	RedisConfig Redis `env:",prefix=REDIS_"`
	Bolt        bool  `env:"BOLT"`
	BoltConfig  Bolt  `env:",prefix=BOLT_"`
	Memory      bool  `env:"MEMORY,default=True"`
}

func NewSession() *Session {
	cfg := &Session{}
	_ = envconfig.ProcessWith(context.Background(), cfg, envconfig.PrefixLookuper("SESSION_", envconfig.OsLookuper()))
	return cfg
}

type Redis struct {
	Address    string   `env:"ADDRESS"`
	Addresses  []string `env:"ADDRESSES"`
	Password   string   `env:"PASSWORD"`
	Sentinel   bool     `env:"SENTINEL"`
	Cluster    bool     `env:"CLUSTER"`
	MasterName string   `env:"MASTER_NAME"`
}

type Bolt struct {
	File     string        `env:"FILE,default=/tmp/bolt.db"`
	Interval time.Duration `env:"INTERVAL,default=20s"`
}
