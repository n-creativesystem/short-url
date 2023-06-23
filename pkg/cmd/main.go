package cmd

import "github.com/n-creativesystem/short-url/pkg/utils/logging"

func Execute() {
	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		logging.Default().Fatalf("err: %v", err)
	}
}
