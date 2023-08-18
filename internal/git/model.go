package git

type Commit struct {
	sha     string
	kind    string
	emoji   string
	author  string
	message string
	date    string
}

type Gitmoji struct {
	emoji     string
	commits   []Commit
	occurence int
}

var (
	History  []Commit
	Gitmojis []Gitmoji
	Commits  []Commit
)
