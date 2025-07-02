package out

import (
	"fmt"
	"rconn/src/constants"
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

type PromptInputParams struct {
	Prompt     string
	Default    string
	IsPassword bool
}

func PromptInput(params PromptInputParams) (string, error) {
	promptIcon := pterm.NewStyle(pterm.FgLightBlue, pterm.Bold).Sprint("[?]")
	promptText := pterm.NewStyle(pterm.FgLightWhite).Sprint(params.Prompt)

	input := pterm.DefaultInteractiveTextInput.
		WithTextStyle(pterm.NewStyle(pterm.FgWhite)).
		WithDefaultText(params.Default)

	if params.IsPassword {
		input = input.WithMask("*")
	}

	return input.Show(promptIcon + " " + promptText)
}
