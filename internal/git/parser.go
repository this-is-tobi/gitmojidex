package git

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

var (
	sep        string = " | "
	prettyArgs string = "%h" + sep + "%s" + sep + "%an" + sep + "%as"
)

func FetchHistory(path string) {
	c := exec.Command("git", "log", "--all", "--pretty="+prettyArgs)
	c.Dir = expandHome(path)
	out, err := c.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	rawCommits := strings.Split(string(out), "\n")
	History = lo.Map(rawCommits[:len(rawCommits)-1], formatCommit)
	Gitmojis = sortByEmoji(lo.Reduce(History, joinByEmoji, []Gitmoji{}))
	Commits = History
}

func FilterHistory(user string) {
	Commits = FilterByUser(History, user)
	Gitmojis = sortByEmoji(lo.Reduce(Commits, joinByEmoji, []Gitmoji{}))
}

func formatCommit(s string, _ int) Commit {
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
	kindIds := regexp.MustCompile(`[a-zA-Z_]*:`).FindStringIndex(input)
	emojiIds := regexp.MustCompile(`:[a-zA-Z_]*:`).FindStringIndex(input)
	if len(kindIds) > 0 {
		kind = input[kindIds[0] : kindIds[1]-1]
	} else {
		kind = ""
	}
	if len(emojiIds) > 0 {
		emoji = input[emojiIds[0]:emojiIds[1]]
	} else {
		emoji = ""
	}
	ids := append(kindIds, emojiIds...)
	if len(ids) > 0 {
		for j := 1; j < len(ids); j++ {
			if ids[0] < ids[j] {
				ids[0] = ids[j]
			}
		}
		input = input[ids[0]:]
	}
	return kind, emoji, input
}
