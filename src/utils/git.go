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
	var formatedCommits []Commit = Map(rawCommits[:len(rawCommits)-1], formatCommit)
	return filterByUser(formatedCommits, user)
}

func GetFormatedEmoji(rawCommits []string, user string) []Commit {
	formatedCommits := Map(rawCommits[:len(rawCommits)-1], formatCommit)
	filteredCommits := filterByUser(formatedCommits, user)
	test := Reduce(filteredCommits, joinByEmoji, []Commit{})

	return test
}

func joinByEmoji(acc []Commit, cur Commit) []Commit {
	for i, a := range acc {
		if a.gitmoji == cur.gitmoji {
			acc[i].occurence += 1
			return acc
		}
	}
	acc = append(acc, Commit{gitmoji: cur.gitmoji, occurence: 1})
	return acc
}

func formatCommit(s string) Commit {
	commit := strings.Split(s, sep)
	kind, gitmoji, message := parseCommitMessage(commit[2])

	return Commit{
		kind:    strings.TrimSpace(kind),
		gitmoji: strings.TrimSpace(gitmoji),
		sha:     strings.TrimSpace(commit[0]),
		message: strings.TrimSpace(message),
		author:  strings.TrimSpace(commit[1]),
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
