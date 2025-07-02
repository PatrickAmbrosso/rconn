package out

import (
	"fmt"
	"os"
	"rconn/src/constants"
	"rconn/src/models"
	"strings"

	"github.com/pterm/pterm"
)

var Logger *SCLogger

type SCLogger struct{}

func (l SCLogger) Debug(msg string) {
	icon := pterm.NewStyle(pterm.FgLightCyan, pterm.Bold).Sprint("[?]")
	message := pterm.NewStyle(pterm.FgLightWhite).Sprint(msg)
	fmt.Printf("%s %s\n", icon, message)
}

func (l SCLogger) Info(msg string) {
	icon := pterm.NewStyle(pterm.FgLightBlue, pterm.Bold).Sprint("[i]")
	message := pterm.NewStyle(pterm.FgWhite).Sprint(msg)
	fmt.Printf("%s %s\n", icon, message)
}

func (l SCLogger) Warn(msg string) {
	icon := pterm.NewStyle(pterm.FgYellow, pterm.Bold).Sprint("[!]")
	message := pterm.NewStyle(pterm.FgWhite).Sprint(msg)
	fmt.Printf("%s %s\n", icon, message)
}

func (l SCLogger) Error(msg string) {
	icon := pterm.NewStyle(pterm.FgRed, pterm.Bold).Sprint("[x]")
	message := pterm.NewStyle(pterm.FgWhite).Sprint(msg)
	fmt.Printf("%s %s\n", icon, message)
}

func (l SCLogger) Success(msg string) {
	icon := pterm.NewStyle(pterm.FgGreen, pterm.Bold).Sprint("[✓]")
	message := pterm.NewStyle(pterm.FgWhite).Sprint(msg)
	fmt.Printf("%s %s\n", icon, message)
}

func init() {
	Logger = &SCLogger{}
}

func Banner(message string) string {
	var sb strings.Builder
	sb.WriteString(pterm.Sprintf(constants.AppBanner))
	sb.WriteString("\n")
	sb.WriteString(pterm.FgBlue.Sprint(message))
	sb.WriteString("\n")
	return sb.String()
}

func PromptInput(params models.PromptInputParams) (string, error) {
	promptIcon := pterm.NewStyle(pterm.FgLightMagenta, pterm.Bold).Sprint("[?]")
	promptLabel := pterm.NewStyle(pterm.FgLightWhite).Sprint(params.Label)

	input := pterm.DefaultInteractiveTextInput.
		WithDefaultText(params.DefaultValue).
		WithOnInterruptFunc(func() {
			Logger.Error("Aborted by user")
			os.Exit(1)
		})

	if params.IsPassword {
		input = input.WithMask("*")
	}

	return input.Show(promptIcon + " " + promptLabel)
}
