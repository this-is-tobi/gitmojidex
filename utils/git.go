package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var sep string = " | "
var prettyArgs string = "%h" + sep + "%an" + sep + "%s" + sep + "%as"

func GetRawHistory(path string) []string {
	c := exec.Command("git", "log", "--all", "--pretty="+prettyArgs)
	c.Dir = path
	out, err := c.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	return strings.Split(string(out), "\n")
}

func GetFormatedHistory(rawCommits []string, user string) []Commit {
	formatedCommits := Map(rawCommits[:len(rawCommits)-1], formatCommit)
	return filterByUser(formatedCommits, user)
}

func GetFormatedEmoji(rawCommits []string, user string) []Gitmoji {
	formatedCommits := Map(rawCommits[:len(rawCommits)-1], formatCommit)
	filteredCommits := filterByUser(formatedCommits, user)
	return Reduce(filteredCommits, JoinByEmoji, []Gitmoji{})
}

func JoinByEmoji(acc []Gitmoji, cur Commit) []Gitmoji {
	for i, a := range acc {
		if a.emoji == cur.gitmoji {
			acc[i].commits = append(acc[i].commits, cur)
			acc[i].occurence += 1
			return acc
		}
	}
	acc = append(acc, Gitmoji{emoji: cur.gitmoji, commits: []Commit{cur}, occurence: 1})
	return acc
}

func formatCommit(s string) Commit {
	commit := strings.Split(s, sep)
	kind, gitmoji, message := parseCommitMessage(commit[1])
	return Commit{
		sha:     strings.TrimSpace(commit[0]),
		kind:    strings.TrimSpace(kind),
		gitmoji: strings.TrimSpace(gitmoji),
		message: strings.TrimSpace(message),
		author:  strings.TrimSpace(commit[2]),
		date:    strings.TrimSpace(commit[3]),
	}
}

func filterByUser(rawCommits []Commit, user string) []Commit {
	var formatedCommits []Commit
	for _, commit := range rawCommits {
		if commit == (Commit{}) {
			continue
		}
		if user == "" || regexp.MustCompile("(?i)"+user).FindString(commit.author) != "" {
			formatedCommits = append(formatedCommits, commit)
		}
	}
	return formatedCommits
}

func parseCommitMessage(input string) (string, string, string) {
	var kind string
	var emoji string
	kindIndexes := regexp.MustCompile(`[a-zA-Z_]*:`).FindStringIndex(input)
	emojiIndexes := regexp.MustCompile(`:[a-zA-Z_]*:`).FindStringIndex(input)
	if len(kindIndexes) > 0 {
		kind = input[kindIndexes[0] : kindIndexes[1]-1]
	} else {
		kind = ""
	}
	if len(emojiIndexes) > 0 {
		emoji = input[emojiIndexes[0]:emojiIndexes[1]]
	} else {
		emoji = ""
	}
	indexes := append(kindIndexes, emojiIndexes...)
	if len(indexes) > 0 {
		for j := 1; j < len(indexes); j++ {
			if indexes[0] < indexes[j] {
				indexes[0] = indexes[j]
			}
		}
		input = input[indexes[0]:]
	}
	return kind, emoji, input
}
