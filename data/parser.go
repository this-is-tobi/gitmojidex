package data

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/this-is-tobi/gitmojidex/utils"
)

var (
	sep        string = " | "
	prettyArgs string = "%h" + sep + "%s" + sep + "%an" + sep + "%as"
)

func FetchHistory(path string) {
	c := exec.Command("git", "log", "--all", "--pretty="+prettyArgs)
	c.Dir = path
	out, err := c.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	rawCommits := strings.Split(string(out), "\n")
	History = utils.Map(rawCommits[:len(rawCommits)-1], formatCommit)
	Gitmojis = utils.Reduce(History, joinByEmoji, []Gitmoji{})
	Commits = History
}

func FilterHistory(user string) {
	Commits = FilterByUser(History, user)
	Gitmojis = utils.Reduce(Commits, joinByEmoji, []Gitmoji{})
}

func formatCommit(s string) Commit {
	commit := strings.Split(s, sep)
	kind, emoji, message := parseCC(commit[1])
	return Commit{
		sha:     strings.TrimSpace(commit[0]),
		kind:    strings.TrimSpace(kind),
		emoji:   strings.TrimSpace(emoji),
		message: strings.TrimSpace(message),
		author:  strings.TrimSpace(commit[2]),
		date:    strings.TrimSpace(commit[3]),
	}
}

func parseCC(input string) (string, string, string) {
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
