package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gogita/ui/keys"
	"gogita/ui/styles"
)

type Model struct {
	width  int
	height int
	count  int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case keys.IsQuit(msg):
			return m, tea.Quit
		case keys.IsIncrement(msg):
			m.count++
		}
	}
	return m, nil
}

func (m Model) View() string {
	button := styles.Button.Render("Click Up")
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
