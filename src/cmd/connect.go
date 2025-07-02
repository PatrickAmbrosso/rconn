package cmd

import (
	"os"
	"rconn/src/out"
	"rconn/src/utils"
	"strings"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: helpConnectCmd,
	Long:  out.Banner(helpConnectCmd),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var valueErrs []string

		if len(args) > 0 {
			connName := args[0]
			out.Logger.Info("Connecting to saved session: " + connName)
			// You can add a lookup here later
		}

		if flagConnectHostAddress == "" {
			flagConnectHostAddress, err = out.PromptInput(out.PromptInputParams{
				Prompt:     "Hostname/IP addr of computer to RDP into",
				IsPassword: false,
			})
			if err != nil {
				valueErrs = append(valueErrs, "host address")
			}
		}

		if flagConnectUsername == "" {
			flagConnectUsername, err = out.PromptInput(out.PromptInputParams{
				Prompt:     "User account to connect to " + flagConnectHostAddress,
				IsPassword: false,
			})
			if err != nil {
				valueErrs = append(valueErrs, "username")
			}
		}

		if flagConnectPassword == "" {
			flagConnectPassword, err = out.PromptInput(out.PromptInputParams{
				Prompt:     "Password for user " + flagConnectUsername,
				IsPassword: true,
			})
			if err != nil {
				valueErrs = append(valueErrs, "password")
			}
		}

		if len(valueErrs) > 0 {
			out.Logger.Error("Missing required input(s): " + strings.Join(valueErrs, ", "))
			os.Exit(1)
		}

		err = utils.ConnectRDP(utils.RDPConnection{
			HostAddress: flagConnectHostAddress,
			Username:    flagConnectUsername,
			Password:    flagConnectPassword,
		})

		if err != nil {
			out.Logger.Error(err.Error())
			os.Exit(1)
		}

		out.Logger.Info("RDP session launched successfully")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&flagConnectHostAddress, "address", "a", "", "Host address")
	connectCmd.Flags().StringVarP(&flagConnectUsername, "username", "u", "", "Username")
	connectCmd.Flags().StringVarP(&flagConnectPassword, "password", "p", "", "Password")
}
