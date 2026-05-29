package ui

import (
	"fmt"
	"gogita/ui/keys"
	"gogita/ui/styles"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type branchMsg string
type errMsg string

func getBranchName() tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			return errMsg(fmt.Sprintf("Error getting branch name: %s", err))
		}
		return branchMsg(strings.TrimSpace(string(out)))
	}
}

type Model struct {
	width      int
	height     int
	branchName string
	textInput  textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Focus()
	return Model{
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(getBranchName(), textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case branchMsg:
		m.branchName = string(msg)
	case errMsg:
		m.branchName = fmt.Sprintf("Error: %s", string(msg))
	case tea.KeyMsg:
		switch {
		case keys.IsQuit(msg):
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {

	mainWidth := m.width * 3 / 4
	sideWidth := m.width / 4

	mainContent := lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Button.Render(m.branchName),
		styles.TextInput.Render(m.textInput.View()),
	)

	mainPane := styles.MainPanel.Width(mainWidth).Height(m.height).Render(
		lipgloss.Place(
			mainWidth,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			mainContent,
		),
	)

	sidePane := styles.SidePanel.
		Width(sideWidth - 2).
		Height(m.height).
		Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		sidePane,
		mainPane,
	)
}
