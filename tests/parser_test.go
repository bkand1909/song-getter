package tests

import (
	"fmt"
	"github.com/bkand1909/song-getter/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestZingSongParser(t *testing.T) {
	url := "http://mp3.zing.vn/bai-hat/Co-Don-Cung-Khong-Khoc-Nguyen-Vu/ZW7I9BAB.html"
	html := ""
	Convey("Get html from "+url, t, func() {
		resp := utils.GetUrl(url, nil)
		So(resp.Ok, ShouldBeTrue)
		html = resp.String()
		Convey("Should return valid html", func() {
			So(html, ShouldNotBeEmpty)
		})
	})
	var parser utils.ZingParser
	Convey("ZingParser - ToSong", t, func() {
		err, song := parser.ToSong(url, html)
		fmt.Println(utils.JsonIndentStruct(song))
		Convey("Should return valid song", func() {
			So(err, ShouldBeNil)
			So(song, ShouldNotBeNil)
		})
	})
}

func TestZingAlbumParser(t *testing.T) {
	url := "http://mp3.zing.vn/album/Cafe-Sai-Gon-Nguyen-Vu/ZWZBFIU0.html"
	html := ""
	Convey("Get html from "+url, t, func() {
		resp := utils.GetUrl(url, nil)
		So(resp.Ok, ShouldBeTrue)
		html = resp.String()
		Convey("Should return valid html", func() {
			So(html, ShouldNotBeEmpty)
		})
	})
	var parser utils.ZingParser
	Convey("ZingParser - ToAlbum", t, func() {
		err, album := parser.ToAlbum(url, html)
		fmt.Println(utils.JsonIndentStruct(album))
		Convey("Should return valid album", func() {
			So(err, ShouldBeNil)
			So(album, ShouldNotBeNil)
		})
	})
}
