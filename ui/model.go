package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"gogita/ui/keys"
	"gogita/ui/styles"
	"os/exec"
	"strings"
)

type branchMsg string
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
	commits    []string
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		getBranchName(),
		getRecentCommits(),
	)
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
	case commitMsg:
		m.commits = []string(msg)
	case tea.KeyMsg:
		switch {
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
		if strings.Contains(commit, "HEAD ->") {
			commitLines = append(commitLines, styles.CommitCurrent.Width(innerWidth).Render(commit))
		} else if i%2 == 0 {
			commitLines = append(commitLines, styles.CommitEven.Width(innerWidth).Render(commit))
		} else {
			commitLines = append(commitLines, styles.CommitOdd.Width(innerWidth).Render(commit))
		}
	}
	content := strings.Join(commitLines, "\n")

	// Full-screen black background with padded commits
	return styles.MainPanel.Width(m.width).Height(m.height).Render(content)
}
