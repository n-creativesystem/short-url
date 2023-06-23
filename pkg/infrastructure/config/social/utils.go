package social

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

func envProcess(i interface{}, prefix string) error {
	return envconfig.ProcessWith(context.Background(), i, envconfig.PrefixLookuper(prefix, envconfig.OsLookuper()))
}
