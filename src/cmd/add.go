package cmd

import (
	"os"
	"rconn/src/models"
	"rconn/src/out"
	"rconn/src/utils"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: helpAddCmd,
	Long:  out.Banner(helpAddCmd),
	Run: func(cmd *cobra.Command, args []string) {
		connName := ""
		if len(args) > 0 {
			connName = args[0]
		}

		params, err := utils.PromptRDPConnectionParams(false, models.RDPConnectionParams{
			Name:     connName,
			Host:     flagAddHostAddress,
			User:     flagAddUsername,
			Password: flagAddPassword,
		})

		if err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		store, err := utils.GetStore(flagGlobalConfigPath)
		if err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		if err := store.Add(params); err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		out.Logger.Success("RDP connection " + params.Name + " saved successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
