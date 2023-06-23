package cmd

import "github.com/spf13/cobra"

func makeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:           "make",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Root().PersistentPreRun(cmd, args)
		},
	}
	cmd.AddCommand(makeCryptKeyCommand())
	return &cmd
}
