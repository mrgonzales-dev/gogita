package styles

import "github.com/charmbracelet/lipgloss"

var (
	CommitEven    = lipgloss.NewStyle().Background(lipgloss.Color("#111111"))
	CommitOdd     = lipgloss.NewStyle().Background(lipgloss.Color("#1e1e1e"))
	CommitCurrent = lipgloss.NewStyle().Background(lipgloss.Color("#4a148c"))

	MainPanel = lipgloss.NewStyle().
			Background(lipgloss.Color("#000")).
			PaddingTop(1).
			PaddingLeft(2).
			PaddingRight(2)
)
