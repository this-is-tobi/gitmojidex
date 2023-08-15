package utils

type Commit struct {
	sha     string
	kind    string
	gitmoji string
	author  string
	message string
	date    string
}

type Gitmoji struct {
	emoji     string
	commits   []Commit
	occurence int
}
