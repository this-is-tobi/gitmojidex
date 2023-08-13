package utils

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kyokomi/emoji/v2"
)

func TableEmoji(commits []Commit) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Gitmoji", "Number of occurence"})
	tw.AppendRows([]table.Row{
		{"🎉", "1"},
	})
	tw.AppendFooter(table.Row{"Total", len(commits)})
	tw.SetAutoIndex(true)
	tw.SetStyle(table.StyleLight) // table.Default || table.StyleColoredBright
	// tw.SetCaption("Gitmojidex for repository TEST!\n")

	fmt.Println(tw.Render())
}

func TableHistory(commits []Commit) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Gitmoji", "Kind", "Message", "Author", "Date"})
	tw.AppendRows(Map(commits, commitToRow))
	tw.AppendFooter(table.Row{"Total", len(commits)})
	tw.SetAutoIndex(true)
	tw.SetStyle(table.StyleLight) // table.Default || table.StyleColoredBright
	// tw.SetCaption("Gitmojidex for repository TEST!\n")

	fmt.Println(tw.Render())
}

func commitToRow(commit Commit) table.Row {
	return table.Row{
		emoji.Emojize(commit.gitmoji),
		commit.kind,
		commit.message,
		commit.author,
		commit.date,
	}
}
