package utils

import (
	"fmt"
	"sort"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/v6/table"
)

func TableEmoji(commits []Commit) {
	sortedCommits := commits
	sort.Slice(sortedCommits, func(i, j int) bool {
		return sortedCommits[i].occurence > sortedCommits[j].occurence
	})

	t := table.NewWriter()
	t.AppendHeader(table.Row{"", "Gitmoji", "", "", "", "", "Occurence"})
	t.AppendRows(Map(commits, commitToRow))
	t.SetAutoIndex(true)
	t.SuppressEmptyColumns()
	t.SetStyle(table.StyleLight)

	fmt.Println(t.Render())
}

func TableHistory(commits []Commit, argSort string) {
	sortedCommits := commits
	sort.Slice(sortedCommits, func(i, j int) bool {
		if argSort == "asc" {
			return sortedCommits[i].date < sortedCommits[j].date
		} else {
			return sortedCommits[i].date > sortedCommits[j].date
		}
	})

	t := table.NewWriter()
	t.AppendHeader(table.Row{"Sha", "Gitmoji", "Kind", "Message", "Author", "Date"})
	t.AppendRows(Map(sortedCommits, commitToRow))
	t.AppendFooter(table.Row{"Total", len(commits)})
	t.SetAutoIndex(true)
	t.SuppressEmptyColumns()
	t.SetStyle(table.StyleLight)
	// t.Style().Options.DrawBorder = false
	// t.Style().Options.SeparateColumns = false

	fmt.Println(t.Render())
}

func commitToRow(commit Commit) table.Row {
	if commit.occurence == 0 {
		return table.Row{
			commit.sha,
			fmt.Sprintf("%s - %s", emoji.Parse(commit.gitmoji), commit.gitmoji),
			commit.kind,
			commit.message,
			commit.author,
			commit.date,
		}
	} else {
		return table.Row{
			commit.sha,
			fmt.Sprintf("%s - %s", emoji.Parse(commit.gitmoji), commit.gitmoji),
			commit.kind,
			commit.message,
			commit.author,
			commit.date,
			commit.occurence,
		}
	}
}
