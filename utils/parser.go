package utils

import (
	"github.com/bkand1909/song-getter/models"
)

type Parser struct {
	ToSong  func(url string, html string) *models.Song
	ToAlbum func(url string, html string) *models.Album
}
