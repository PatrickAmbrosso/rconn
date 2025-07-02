package cmd

import (
	"os"
	"rconn/src/models"
	"rconn/src/out"
	"rconn/src/utils"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: helpConnectCmd,
	Long:  out.Banner(helpConnectCmd),
	Run: func(cmd *cobra.Command, args []string) {

		store, err := utils.GetStore(flagGlobalConfigPath)
		if err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		if flagConnectInteractive {
			connections := store.List()
			if len(connections) == 0 {
				out.Logger.Warn("No saved RDP connections found")
				os.Exit(0)
			}

			var options []string

			for _, conn := range connections {
				options = append(options, conn.Name+": "+conn.User+"@"+conn.Host)
			}

			selected, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
			connName := strings.Split(selected, ": ")[0]
			params, err := store.Get(connName)
			if err != nil {
				out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
				os.Exit(1)
			}
			if err := utils.ConnectRDP(*params); err != nil {
				out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
				os.Exit(1)
			}
			out.Logger.Info("RDP session launched successfully")
			return
		}

		if len(args) > 0 {
			connName := args[0]
			out.Logger.Info("Connecting to saved session: " + connName)
			params, err := store.Get(connName)
			if err != nil {
				out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
				os.Exit(1)
			}
			if err := utils.ConnectRDP(*params); err != nil {
				out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
				os.Exit(1)
			}
			out.Logger.Info("RDP session launched successfully")
			return
		}

		params, err := utils.PromptRDPConnectionParams(true, models.RDPConnectionParams{})
		if err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		if err := utils.ConnectRDP(params); err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		out.Logger.Info("RDP session launched successfully")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&flagConnectHostAddress, "address", "a", "", "hostname or ip address of the target machine")
	connectCmd.Flags().StringVarP(&flagConnectUsername, "user", "u", "", "user account to use for the rdp session")
	connectCmd.Flags().StringVarP(&flagConnectPassword, "password", "p", "", "password for the user account")
	connectCmd.Flags().BoolVarP(&flagConnectInteractive, "interactive", "i", false, "choose a saved connection interactively")
}
