package cmd

import (
	"rconn/src/out"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: helpRemoveCmd,
	Long:  out.Banner(helpRemoveCmd),
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
