package out

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

// Base color palette
var (
	ColorAccent      = lipgloss.Color("10") // Bright green
	ColorSubtle      = lipgloss.Color("8")  // Bright black / gray
	ColorError       = lipgloss.Color("9")  // Red
	ColorPrimary     = lipgloss.Color("10") // Bright green
	ColorPlaceholder = lipgloss.Color("7")  // Light gray
)

func GetTerminalWidth() int {
	width, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil || width <= 0 {
		return 80 // fallback width
	}
	return width
}

func GetTheme() *huh.Theme {
	width := GetTerminalWidth()

	formStyles := huh.FormStyles{
		Base: lipgloss.NewStyle().
			Width(width-2).
			Padding(1, 2).
			Margin(1, 0).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(ColorAccent),
	}

	groupStyles := huh.GroupStyles{
		Base: lipgloss.NewStyle(),
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary),
		Description: lipgloss.NewStyle().
			Italic(true).
			Foreground(ColorSubtle),
	}

	fieldStyles := huh.FieldStyles{
		Base: lipgloss.NewStyle(),

		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorAccent),

		Description: lipgloss.NewStyle().
			Foreground(ColorSubtle),

		ErrorIndicator: lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true),

		ErrorMessage: lipgloss.NewStyle().
			Foreground(ColorError),

		TextInput: huh.TextInputStyles{
			Cursor:      lipgloss.NewStyle().Reverse(true),
			CursorText:  lipgloss.NewStyle().Foreground(ColorPrimary),
			Placeholder: lipgloss.NewStyle().Italic(true).Foreground(ColorPlaceholder),
			Prompt:      lipgloss.NewStyle().Foreground(ColorAccent),
			Text:        lipgloss.NewStyle().Foreground(lipgloss.Color("15")),
		},

		FocusedButton: lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(ColorAccent).
			Bold(true).
			Padding(0, 2),

		BlurredButton: lipgloss.NewStyle().
			Foreground(ColorAccent).
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorAccent).
			Padding(0, 2),
	}

	helpStyles := help.Styles{
		ShortKey:       lipgloss.NewStyle().Foreground(ColorAccent),
		ShortDesc:      lipgloss.NewStyle().Foreground(ColorSubtle),
		ShortSeparator: lipgloss.NewStyle().Foreground(ColorSubtle),
		FullKey:        lipgloss.NewStyle().Foreground(ColorAccent),
		FullDesc:       lipgloss.NewStyle().Foreground(ColorSubtle),
		FullSeparator:  lipgloss.NewStyle().Foreground(ColorSubtle),
	}

	return &huh.Theme{
		Form:           formStyles,
		Group:          groupStyles,
		FieldSeparator: lipgloss.NewStyle().Foreground(ColorSubtle),
		Blurred:        fieldStyles,
		Focused:        fieldStyles,
		Help:           helpStyles,
	}
}
