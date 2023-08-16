package view

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/this-is-tobi/gitmojidex/data"
	"github.com/this-is-tobi/gitmojidex/utils"
)

// sessionState is used to track which model is focused
type sessionState uint

const (
	pathView sessionState = iota
	userView
	commitsView
	gitmojisView
)

func (m model) Init() tea.Cmd {
	return textinput.Blink
	// return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == pathView {
				m.state = userView
			} else if m.state == userView {
				m.state = gitmojisView
			} else if m.state == gitmojisView {
				m.state = commitsView
			} else if m.state == commitsView {
				m.state = pathView
			}
		case "enter":
			data.FilterHistory("stan")
			gRows := utils.Map(data.Gitmojis, data.GitmojiToRow)
			cRows := utils.Map(data.Commits, data.CommitToRow)
			m.gitmojiTable.SetRows(gRows)
			m.commitTable.SetRows(cRows)
			m.userInput.SetValue("stan")
			// gitmojiTable := newTable(gCols, gRows)
			// commitTable := newTable(hCols, cRows)
			// userInput := newInput("stan")
			// m = newModel(gitmojiTable, commitTable, userInput)
		}
		switch m.state {
		case commitsView:
			m.commitTable, cmd = m.commitTable.Update(msg)
			cmds = append(cmds, cmd)
		case gitmojisView:
			m.gitmojiTable, cmd = m.gitmojiTable.Update(msg)
			cmds = append(cmds, cmd)
		case userView:
			m.userInput, cmd = m.userInput.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var (
		s  string
		pr string
		ur string
		gr string
		cr string
	)
	if m.state == pathView {
		pr = focus(inputStyle).Render(m.pathInput.View())
		ur = unfocus(inputStyle).Render(m.userInput.View())
		gr = unfocus(tableGitmojiStyle).Render(m.gitmojiTable.View())
		cr = unfocus(tableCommitStyle).Render(m.commitTable.View())
	} else if m.state == userView {
		pr = unfocus(inputStyle).Render(m.pathInput.View())
		ur = focus(inputStyle).Render(m.userInput.View())
		gr = unfocus(tableGitmojiStyle).Render(m.gitmojiTable.View())
		cr = unfocus(tableCommitStyle).Render(m.commitTable.View())
	} else if m.state == gitmojisView {
		pr = unfocus(inputStyle).Render(m.pathInput.View())
		ur = unfocus(inputStyle).Render(m.userInput.View())
		gr = focus(tableGitmojiStyle).Render(m.gitmojiTable.View())
		cr = unfocus(tableCommitStyle).Render(m.commitTable.View())
	} else if m.state == commitsView {
		pr = unfocus(inputStyle).Render(m.pathInput.View())
		ur = unfocus(inputStyle).Render(m.userInput.View())
		gr = unfocus(tableGitmojiStyle).Render(m.gitmojiTable.View())
		cr = focus(tableCommitStyle).Render(m.commitTable.View())
	}
	s += lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.JoinVertical(lipgloss.Top, pr, ur, gr), cr)
	s += helpStyle.Render(fmt.Sprintln("\ntab: focus next • enter: search focused input • q: exit"))
	return s
}

func Render(repoPath string, user string) {
	pathInput := newInput("repo path")
	userInput := newInput("username")
	data.FetchHistory("./")
	gitmojiTable, commitTable := renderTables(data.Gitmojis, data.Commits)

	m := newModel(pathInput, userInput, gitmojiTable, commitTable)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func renderTables(g []data.Gitmoji, h []data.Commit) (table.Model, table.Model) {
	gRows := utils.Map(g, data.GitmojiToRow)
	hRows := utils.Map(h, data.CommitToRow)
	return newTable(gCols, gRows), newTable(hCols, hRows)
}
