package cmd

import (
	"github.com/joho/godotenv"
	"github.com/n-creativesystem/short-url/pkg/domain/config"
	"github.com/n-creativesystem/short-url/pkg/infrastructure/rdb"
	"github.com/n-creativesystem/short-url/pkg/utils/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type migratorArgs struct {
	retryWait    uint64
	retryCount   uint64
	allowMissing bool
	table        string
	envFiles     []string
	driver       string
}

func (arg *migratorArgs) Convert() rdb.MigratorArgs {
	return rdb.MigratorArgs{
		Dir:          arg.driver,
		RetryCount:   arg.retryCount,
		RetryWait:    arg.retryWait,
		AllowMissing: arg.allowMissing,
		Table:        arg.table,
		Driver:       arg.driver,
	}
}

func (arg *migratorArgs) setFlag(flags *pflag.FlagSet) {
	flags.Uint64Var(&arg.retryWait, "retryWait", 1000, "connection open retry interval in milliseconds")
	flags.Uint64Var(&arg.retryCount, "retryCount", 0, "maximum connection open retry count (default = 0,no retry)")
	flags.BoolVar(&arg.allowMissing, "allowMissing", false, "allow out-of-order migrations (default = false, test-use-only)")
	flags.StringVar(&arg.table, "table", "", "migrations table name (default goose_db_version)")
	flags.StringArrayVar(&arg.envFiles, "dotenvs", nil, "dotenv files")
	flags.StringVar(&arg.driver, "driver", "mysql", "driver name")
}

func migratorCommand() *cobra.Command {
	var migratorArgs migratorArgs
	cmd := &cobra.Command{
		Use:           "migrator",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ArbitraryArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			config.SetDriver(config.ConvertDriverFromString(migratorArgs.driver))
		},
		Run: func(cmd *cobra.Command, args []string) {
			_ = godotenv.Load(migratorArgs.envFiles...)
			if err := rdb.Migration(cmd.Context(), args, migratorArgs.Convert()); err != nil {
				logging.Default().Error(err)
			}
		},
	}
	flags := cmd.Flags()
	migratorArgs.setFlag(flags)
	return cmd
}
