package cmd

import (
	"os"
	"rconn/src/out"
	"rconn/src/utils"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: helpListCmd,
	Long:  out.Banner(helpListCmd),
	Run: func(cmd *cobra.Command, args []string) {
		store, err := utils.GetStore(flagGlobalConfigPath)
		if err != nil {
			out.Logger.Error(strings.ToUpper(err.Error()[:1]) + err.Error()[1:])
			os.Exit(1)
		}

		connections := store.List()
		if len(connections) == 0 {
			out.Logger.Info("no saved RDP connections found")
			return
		}

		tableData := pterm.TableData{
			{"Name", "Host", "User"},
		}

		for _, conn := range connections {
			tableData = append(tableData, []string{conn.Name, conn.Host, conn.User})
		}

		pterm.DefaultSection.Println("Saved RDP connections")
		_ = pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
