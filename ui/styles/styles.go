package styles

import "github.com/charmbracelet/lipgloss"

var (
	SidePanel = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#000")).
			Background(lipgloss.Color("#000"))

	Button = lipgloss.NewStyle().
		Padding(0, 4).
		BorderStyle(lipgloss.NormalBorder())

	TextInput = lipgloss.NewStyle().
			Padding(0, 2).
			BorderStyle(lipgloss.NormalBorder())

	MainPanel = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#000")).
			Background(lipgloss.Color("#000"))
)
