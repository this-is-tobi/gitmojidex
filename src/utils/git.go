package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var separator string = " | "
var prettyArgs string = "%an" + separator + "%s" + separator + "%as"

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

func formatCommit(s string) Commit {
	commit := strings.Split(s, separator)
	kind, gitmoji, message := parseCommitMessage(commit[1])

	return Commit{
		kind:    strings.TrimSpace(kind),
		gitmoji: strings.TrimSpace(gitmoji),
		message: strings.TrimSpace(message),
		author:  strings.TrimSpace(commit[0]),
		date:    strings.TrimSpace(commit[2]),
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
