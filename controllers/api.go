package controllers

import (
	"archive/zip"
	"fmt"
	"github.com/bkand1909/song-getter/Godeps/_workspace/src/github.com/astaxie/beego"
	"github.com/bkand1909/song-getter/Godeps/_workspace/src/github.com/levigross/grequests"
	"github.com/bkand1909/song-getter/utils"
	"os"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

func (c *ApiController) Post() {

}

func (c *ApiController) Get() {
	url := c.Input().Get("url")
	resp := utils.GetUrl(url, nil)
	html := resp.String()
	var parser utils.ZingParser
	_, album := parser.ToAlbum(url, html)
	staticDir := beego.AppPath + "/" + beego.StaticDir["/file"]
	album.Folder = staticDir + "/" + album.Title
	if err := os.Mkdir(album.Folder, 0777); err != nil {
		beego.Error(err.Error())
		return
	}
	zfile, err := os.Create(album.Folder + ".zip")
	if err != nil {
		beego.Error(err.Error())
		return
	}
	w := zip.NewWriter(zfile)
	for i, song := range album.Song {
		beego.Info("Downloading: " + song.Source)
		song.Filename = song.Title + " - " + song.Performer + "." + song.Type
		// song.Filename = album.Folder + "/" + song.Filename
		resp := grequests.Get(song.Source, nil)
		if resp.Error != nil {
			beego.Error(resp.Error.Error())
		} else {
			beego.Info(fmt.Sprintf("Downloaded %s. Done: %d/%d.", song.Source, i+1, len(album.Song)))
			wr, err := w.Create(song.Filename)
			if err != nil {
				beego.Error(err.Error())
			} else {
				wr.Write(resp.Bytes())
			}
		}
	}
	w.Close()
	c.Data["json"] = &album
	c.ServeJson()
}
