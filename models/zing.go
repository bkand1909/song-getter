package models

type ZingSong struct {
	Song
	BackImage string
	Hq        string
}

type ZingAlbum struct {
	Album
	Duration int
	Song     []ZingSong
}
