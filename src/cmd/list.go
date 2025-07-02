package cmd

import (
	"rconn/src/out"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: helpListCmd,
	Long:  out.Banner(helpListCmd),
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
