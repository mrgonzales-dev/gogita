package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gogita/ui/keys"
	"gogita/ui/styles"
	"os/exec"
	"strings"
)

type errMsg string
type commitMsg []string

func getRecentCommits() tea.Cmd {
	return func() tea.Msg {
		out, err := exec.Command("git", "log", "--oneline", "--decorate=short", "--all", "-30").Output()
		if err != nil {
			return errMsg(fmt.Sprintf("Error getting recent commits: %s", err))
		}
		return commitMsg(strings.Split(strings.TrimSuffix(string(out), "\n"), "\n"))
	}
}

func Enter(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyEnter
}

type Model struct {
	width   int
	height  int
	commits []string
	cursor  int
	message string
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		getRecentCommits(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case commitMsg:
		m.commits = []string(msg)
	case tea.KeyMsg: //keys
		switch {
		// refresh using f5, re-renders the tui
		case keys.IsRefresh(msg):
			m.message = ""
			return m, getRecentCommits()
		case keys.IsEnter(msg):
			m.message = "hello"
		case keys.IsUp(msg):
			if m.cursor > 0 {
				m.cursor--
			}
		case keys.IsDown(msg):
			if m.cursor < len(m.commits)-1 {
				m.cursor++
			}
		case keys.IsQuit(msg):
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) View() string {
	// If we don't have terminal dimensions yet, show a placeholder
	if m.width == 0 {
		return "Loading..."
	}

	innerWidth := m.width - 4

	// Commit list with alternating rows, highlight current commit
	var commitLines []string

	for i, commit := range m.commits {
		switch {
		case i == m.cursor:
			commitLines = append(commitLines, styles.CommitSelected.Width(innerWidth).Render(commit))
		case strings.Contains(commit, "HEAD ->"):
			commitLines = append(commitLines, styles.CommitCurrent.Width(innerWidth).Render(commit))
		case i%2 == 0:
			commitLines = append(commitLines, styles.CommitEven.Width(innerWidth).Render(commit))
		default:
			commitLines = append(commitLines, styles.CommitOdd.Width(innerWidth).Render(commit))
		}
	}

	//render message in the middle as modal
	if m.message != "" {
		return styles.MainPanel.Width(m.width).Height(m.height).Render(m.message)
	}

	content := strings.Join(commitLines, "\n")

	hints := "Press [q] to quit, [f5] to refresh, [j] [k] to navigate"
	panel := styles.MainPanel.Width(m.width).Height(m.height - 1).Render(content)
	bar := styles.ActionBar.Width(m.width).Render(hints)

	// Full-screen black background with padded commits
	return lipgloss.JoinVertical(lipgloss.Top, panel, bar)
}
