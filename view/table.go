package view

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	gCols []table.Column = []table.Column{
		{Title: "Emoji", Width: 18},
		{Title: "Occurence", Width: 18},
	}
	hCols []table.Column = []table.Column{
		{Title: "SHA", Width: 10},
		{Title: "Scope", Width: 8},
		{Title: "Emoji", Width: 9},
		{Title: "Message", Width: 35},
		{Title: "Author", Width: 15},
		{Title: "Date", Width: 10},
	}
)

func newTable(cols []table.Column, rows []table.Row) table.Model {
	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Background(lipgloss.Color("201")).
		Bold(false)

	t.SetStyles(s)
	return t
}
