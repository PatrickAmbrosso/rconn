package out

import (
	"os"
	"rconn/src/constants"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/pterm/pterm"
)

var Logger *log.Logger

func init() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Level:           log.DebugLevel,
		Formatter:       log.TextFormatter,
	})

	commonStyle := lipgloss.NewStyle().Bold(true)
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = commonStyle.SetString("[-]").Foreground(lipgloss.Color("63"))
	styles.Levels[log.InfoLevel] = commonStyle.SetString("[i]").Foreground(lipgloss.Color("86"))
	styles.Levels[log.WarnLevel] = commonStyle.SetString("[!]").Foreground(lipgloss.Color("192"))
	styles.Levels[log.ErrorLevel] = commonStyle.SetString("[✗]").Foreground(lipgloss.Color("204"))
	styles.Levels[log.FatalLevel] = commonStyle.SetString("[◈]").Foreground(lipgloss.Color("134"))

	Logger.SetStyles(styles)
}

func Banner(message string) string {
	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		PaddingTop(1).
		PaddingBottom(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4"))

	banner := bannerStyle.Render(strings.Trim(constants.AppBanner, "\n"))

	if strings.TrimSpace(message) != "" {
		helpBlock := helpStyle.Render(message)
		return lipgloss.JoinVertical(lipgloss.Left, banner, helpBlock)
	}

	return banner
}

type PromptInputParams struct {
	SectionHeading string
	Title          string
	Description    string
	Placeholder    string
	IsPassword     bool
}

func PromptInput2(params PromptInputParams) (string, error) {
	var input string
	echoMode := huh.EchoModeNormal

	if params.IsPassword {
		echoMode = huh.EchoModePassword
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Enter the following information to connect to the remote computer").Height(1),
			huh.NewInput().
				Title(params.Title).
				Description(params.Description).
				Placeholder(params.Placeholder).
				Value(&input).
				EchoMode(echoMode),
		),
	).WithTheme(GetTheme())

	if err := form.Run(); err != nil {
		// Logger.Error("Reached error state", "error", err.Error())
		return "", err
	}

	return input, nil
}

func PromptInput(params PromptInputParams) (string, error) {
	prompt := pterm.DefaultInteractiveTextInput.
		WithDefaultText(params.Placeholder).
		WithDefaultValue(params.Placeholder)

	if params.IsPassword {
		prompt.WithMask("*")
	}

}
