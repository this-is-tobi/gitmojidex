package table

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/this-is-tobi/gitmojidex/utils"
)

type model struct {
	table table.Model
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	path       string = "./"
	user       string = "tobi"
	sep        string = " | "
	prettyArgs string = "%h" + sep + "%s" + sep + "%an" + sep + "%as"
)

func GetTableData(fetchType string) ([]table.Column, []table.Row) {
	c := exec.Command("git", "log", "--all", "--pretty="+prettyArgs)
	c.Dir = path
	out, err := c.Output()

	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	rawCommits := strings.Split(string(out), "\n")
	commits := utils.GetFormatedHistory(rawCommits, user)
	var cols []table.Column
	var rows []table.Row
	if fetchType == "history" {
		cols = []table.Column{
			{Title: "SHA", Width: 10},
			{Title: "Scope", Width: 8},
			{Title: "Emoji", Width: 9},
			{Title: "Message", Width: 35},
			{Title: "Author", Width: 15},
			{Title: "Date", Width: 10},
		}
		rows = utils.Map(commits, utils.CommitToRow)
	} else if fetchType == "emoji" {
		cols = []table.Column{
			{Title: "Emoji", Width: 9},
			{Title: "Occurence", Width: 9},
		}
		gitmojis := utils.Reduce(commits, utils.JoinByEmoji, []utils.Gitmoji{})
		rows = utils.Map(gitmojis, utils.GitmojiToRow)
	}
	return cols, rows
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			cols, rows := GetTableData("history")
			m.table.SetColumns(cols)
			m.table.SetRows(rows)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func CreateTable(cols []table.Column, rows []table.Row) model {
	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(21),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := model{t}
	return m
}
