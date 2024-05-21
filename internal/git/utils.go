package git

import (
	"os/user"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	// "github.com/kyokomi/emoji/v2"
)

func expandHome(p string) string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	if p == "~" || p == "$HOME" {
		p = homeDir
	} else if strings.HasPrefix(p, "~/") {
		p = filepath.Join(homeDir, p[2:])
	} else if strings.HasPrefix(p, "$HOME/") {
		p = filepath.Join(homeDir, p[6:])
	}
	return p
}

func CommitToRow(commit Commit, _ int) table.Row {
	return table.Row{
		commit.sha,
		commit.kind,
		// emoji.Emojize(commit.emoji),
		commit.emoji,
		commit.message,
		commit.author,
		commit.date,
	}
}

func GitmojiToRow(gitmoji Gitmoji, _ int) table.Row {
	return table.Row{
		// emoji.Emojize(gitmoji.emoji),
		gitmoji.emoji,
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

func joinByEmoji(acc []Gitmoji, cur Commit, _ int) []Gitmoji {
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
