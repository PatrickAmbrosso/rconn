package cmd

import (
	"rconn/src/out"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: helpAddCmd,
	Long:  out.Banner(helpAddCmd),
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
