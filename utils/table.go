package utils

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/enescakir/emoji"
)

func CommitToRow(commit Commit) table.Row {
	return table.Row{
		commit.sha,
		commit.kind,
		emoji.Parse(commit.gitmoji),
		commit.message,
		commit.author,
		commit.date,
	}
}

func GitmojiToRow(gitmoji Gitmoji) table.Row {
	return table.Row{
		emoji.Parse(gitmoji.emoji),
		strconv.Itoa(gitmoji.occurence),
	}
}
