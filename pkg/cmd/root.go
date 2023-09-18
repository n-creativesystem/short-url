package cmd

import (
	"github.com/n-creativesystem/short-url/pkg/cmd/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type rootArgs struct {
	output *flags.EnumFlags
	debug  bool
}

func (arg *rootArgs) setFlag(flag *pflag.FlagSet) {
	flag.VarP(arg.output, "output", "o", "log output format")
	flag.BoolVarP(&arg.debug, "debug", "d", false, "debug mode")
}

var (
	rootArg = &rootArgs{
		output: flags.New("json", "console"),
	}
)

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "short-url",
		SilenceErrors:    true,
		SilenceUsage:     true,
		PersistentPreRun: rootPersistentExecute,
	}
	rootArg.setFlag(cmd.PersistentFlags())
	cmd.AddCommand(serverCommand())
	cmd.AddCommand(migratorCommand())
	cmd.AddCommand(makeCommand())
	return cmd
}

func rootPersistentExecute(cmd *cobra.Command, args []string) {
	// if rootArg.debug {
	// 	logging.DebugMode()
	// } else {
	// 	logging.SetFormat(rootArg.output.String())
	// }
}
