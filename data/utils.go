package data

import (
	"regexp"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/enescakir/emoji"
)

func CommitToRow(commit Commit) table.Row {
	return table.Row{
		commit.sha,
		commit.kind,
		emoji.Parse(commit.emoji),
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

func FilterByUser(h []Commit, user string) []Commit {
	var filteredCommits []Commit
	for _, c := range h {
		if c == (Commit{}) {
			continue
		}
		if regexp.MustCompile("(?i)"+user).FindString(c.author) != "" {
			filteredCommits = append(filteredCommits, c)
		}
	}
	return filteredCommits
}

func joinByEmoji(acc []Gitmoji, cur Commit) []Gitmoji {
	for i, a := range acc {
		if a.emoji == cur.emoji {
			acc[i].commits = append(acc[i].commits, cur)
			acc[i].occurence += 1
			return acc
		}
	}
	acc = append(acc, Gitmoji{emoji: cur.emoji, commits: []Commit{cur}, occurence: 1})
	return acc
}
