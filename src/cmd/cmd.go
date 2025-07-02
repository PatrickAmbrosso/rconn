package cmd

import (
	"fmt"
	"os"
	"rconn/src/constants"
	"rconn/src/out"

	"github.com/spf13/cobra"
)

const (
	helpRootCmd    = "A CLI tool to manage windows RDP connections"
	helpAddCmd     = "Add a new RDP connection"
	helpConnectCmd = "Connect to a remote computer via configured RDP connection"
	helpListCmd    = "List all configured RDP connections"
	helpRemoveCmd  = "Remove an existing RDP connection"
)

var (
	flagAddHostAddress string
	flagAddUsername    string
	flagAddPassword    string
	flagAddDomain      string

	flagConnectHostAddress string
	flagConnectUsername    string
	flagConnectPassword    string

	flagRemoveForce bool
)

var rootCmd = &cobra.Command{
	Use:           constants.AppAbbrName,
	Short:         helpRootCmd,
	Long:          out.Banner(helpRootCmd),
	Version:       constants.AppVersion,
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Help()
		fmt.Println()
		out.Logger.Error("Flag parse error", "error", err.Error())
		os.Exit(1)
		return nil
	})
}

func validateCLIArgsCount(n int, cmd *cobra.Command, args []string) {
	if len(args) != n {
		_ = cmd.Help()
		fmt.Println()
		out.Logger.Error(fmt.Sprintf("Command '%s' expects exactly %d argument(s).", cmd.CommandPath(), n))
		out.Logger.Info(fmt.Sprintf("Run '%s --help' for usage.", cmd.CommandPath()))
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
