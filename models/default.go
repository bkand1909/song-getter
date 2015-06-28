package models

type Song struct {
	Id        string
	Title     string
	Url       string
	Source    string
	Performer string
	Lyric     string
	Type      string
}

type Album struct {
	Title string
	Url   string
}
