package ui

import (
	"fmt"

	"gogita/ui/keys"
	"gogita/ui/styles"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width     int
	height    int
	count     int
	textInput textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Focus()
	return Model{
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
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
		default:
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
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
		"",
		styles.TextInput.Render(m.textInput.View()),
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}
