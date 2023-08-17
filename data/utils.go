package data

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/kyokomi/emoji/v2"
)

func CommitToRow(commit Commit) table.Row {
	return table.Row{
		commit.sha,
		commit.kind,
		emoji.Emojize(commit.emoji),
		commit.message,
		commit.author,
		commit.date,
	}
}

func GitmojiToRow(gitmoji Gitmoji) table.Row {
	return table.Row{
		emoji.Emojize(gitmoji.emoji),
		strconv.Itoa(gitmoji.occurence),
	}
}

func FilterByUser(c []Commit, user string) []Commit {
	var filteredCommits []Commit
	for _, c := range c {
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
	for i, g := range acc {
		if g.emoji == cur.emoji {
			acc[i].commits = append(acc[i].commits, cur)
			acc[i].occurence += 1
			return acc
		}
	}
	if cur.emoji != "" {
		acc = append(acc, Gitmoji{emoji: cur.emoji, commits: []Commit{cur}, occurence: 1})
	}
	return acc
}

func sortByEmoji(gitmojis []Gitmoji) []Gitmoji {
	sort.Slice(gitmojis, func(i, j int) bool {
		return gitmojis[i].occurence > gitmojis[j].occurence
	})
	return gitmojis
}
