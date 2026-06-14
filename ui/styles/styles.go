package styles

import "github.com/charmbracelet/lipgloss"

var (
	SidePanel = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#000")).
			Background(lipgloss.Color("#000"))

	CommitEven = lipgloss.NewStyle().Background(lipgloss.Color("#111111"))
	CommitOdd  = lipgloss.NewStyle().Background(lipgloss.Color("#1e1e1e"))

	Button = lipgloss.NewStyle().
		Padding(0, 4).
		BorderStyle(lipgloss.NormalBorder())

	MainPanel = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#000")).
			Background(lipgloss.Color("#000"))
)
