package view

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/this-is-tobi/gitmojidex/data"
	"github.com/this-is-tobi/gitmojidex/utils"
)

func Render(repoPath string, user string) {
	pathInput := newInput("repo path", true)
	pathInput.SetValue(repoPath)
	userInput := newInput("username", false)
	gitmojiTable := newTable(gCols, utils.Map(data.Gitmojis, data.GitmojiToRow), 19)
	commitTable := newTable(cCols, utils.Map(data.Commits, data.CommitToRow), 25)

	m := newModel(pathInput, userInput, gitmojiTable, commitTable)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
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
			if m.state == commitsView {
				m.state = pathView
			} else {
				m.state++
			}
		case "shift+tab":
			if m.state == pathView {
				m.state = commitsView
			} else {
				m.state--
			}
		case "enter":
			if m.state == pathView {
				data.FetchHistory(m.pathInput.Value())
			} else if m.state == userView {
				data.FilterHistory(m.userInput.Value())
			}
			m.gitmojiTable.SetRows(utils.Map(data.Gitmojis, data.GitmojiToRow))
			m.commitTable.SetRows(utils.Map(data.Commits, data.CommitToRow))
		}
		switch m.state {
		case pathView:
			m.pathInput.Focus()
			m.userInput.Blur()
			m.pathInput, cmd = m.pathInput.Update(msg)
			cmds = append(cmds, cmd)
		case userView:
			m.pathInput.Blur()
			m.userInput.Focus()
			m.userInput, cmd = m.userInput.Update(msg)
			cmds = append(cmds, cmd)
		case commitsView:
			m.pathInput.Blur()
			m.userInput.Blur()
			m.commitTable, cmd = m.commitTable.Update(msg)
			cmds = append(cmds, cmd)
		case gitmojisView:
			m.pathInput.Blur()
			m.userInput.Blur()
			m.gitmojiTable, cmd = m.gitmojiTable.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s string
	pr, ur, gr, cr := updateFocus(m)
	s += lipgloss.JoinHorizontal(lipgloss.Top, lipgloss.JoinVertical(lipgloss.Top, pr, ur, gr), cr)
	s += helpStyle.Render(fmt.Sprintln("\ntab: focus next • enter: search for the focused input • q: exit"))
	return s
}

func updateFocus(m model) (string, string, string, string) {
	pr := focus(inputStyle, false).Render(m.pathInput.View())
	ur := focus(inputStyle, false).Render(m.userInput.View())
	gr := focus(tableGitmojiStyle, false).Render(m.gitmojiTable.View())
	cr := focus(tableCommitStyle, false).Render(m.commitTable.View())
	if m.state == pathView {
		pr = focus(inputStyle, true).Render(m.pathInput.View())
	} else if m.state == userView {
		ur = focus(inputStyle, true).Render(m.userInput.View())
	} else if m.state == gitmojisView {
		gr = focus(tableGitmojiStyle, true).Render(m.gitmojiTable.View())
	} else if m.state == commitsView {
		cr = focus(tableCommitStyle, true).Render(m.commitTable.View())
	}
	return pr, ur, gr, cr
}
