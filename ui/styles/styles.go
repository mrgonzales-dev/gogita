package styles

import "github.com/charmbracelet/lipgloss"

var (
	Button = lipgloss.NewStyle().
		Background(lipgloss.Color("#555")).
		Foreground(lipgloss.Color("#FFF")).
		Padding(0, 4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#888"))

	TextInput = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#888")).
			Padding(0, 2)
)
