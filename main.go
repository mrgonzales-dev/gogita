package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	btnStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#555")).
		Foreground(lipgloss.Color("#FFF")).
		Padding(0, 4).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#888"))
)

type mrg_model struct {
	width  int
	height int
	count  int
}

func (m mrg_model) Init() tea.Cmd {
	return nil
}

func (m mrg_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.count++
		}
	}
	return m, nil
}

func (m mrg_model) View() string {
	button := btnStyle.Render("Click Up")
	header := fmt.Sprintf("Hello World!  Count: %d", m.count)

	content := lipgloss.JoinVertical(lipgloss.Center,
		header,
		"",
		button,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

func main() {
	p := tea.NewProgram(mrg_model{})
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error running program:", err)
		os.Exit(1)
	}
}
