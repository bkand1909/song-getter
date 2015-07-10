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
	beego.Info(fmt.Sprint("Get from url[%s] return html length: [%d]", url, len(html)))
	var parser utils.ZingParser
	err, album := parser.ToAlbum(url, html)
	if err != nil {
		beego.Error(err.Error())
		return
	}
	staticDir := beego.AppPath + "/" + beego.StaticDir["/file"]
	album.Folder = staticDir + "/" + album.Title
	// if err := os.Mkdir(album.Folder, 0777); err != nil {
	// 	beego.Error(err.Error())
	// 	return
	// }
	beego.Info("Album folder: " + album.Folder + ".zip")
	zfile, err := os.Create(album.Folder + ".zip")
	if err != nil {
		beego.Error(err.Error())
		return
	}
	w := zip.NewWriter(zfile)
	// w := zip.NewWriter(c.Ctx.ResponseWriter)
	// album.Song = album.Song[0:1]
	for i, song := range album.Song {
		beego.Info("Downloading: " + song.Source)
		song.Filename = song.Title + " - " + song.Performer + "." + song.Type
		// song.Filename = album.Folder + "/" + song.Filename
		resp := grequests.Get(song.Source, nil)
		beego.Info("Download status:", resp.StatusCode)
		if resp.Error != nil {
			beego.Error(resp.Error.Error())
		} else {
			beego.Info(fmt.Sprintf("Downloaded %s. Done: %d/%d.", song.Source, i+1, len(album.Song)))
			wr, err := w.Create(song.Filename)
			if err == nil {
				_, err = wr.Write(resp.Bytes())
			}
			if err != nil {
				beego.Error(err.Error())
			}
		}
	}
	// c.Ctx.Output.Header("Content-Description", "File Transfer")
	// c.Ctx.Output.Header("Content-Type", "application/octet-stream")
	// c.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	// c.Ctx.Output.Header("Expires", "0")
	// c.Ctx.Output.Header("Cache-Control", "must-revalidate")
	// c.Ctx.Output.Header("Pragma", "public")
	// c.Ctx.Output.Header("Content-Disposition", "attachment; filename="+album.Title+".zip")
	if err := w.Close(); err != nil {
		beego.Error(err.Error())
		return
	}
	c.Data["json"] = `{"success":"true"}`
	c.ServeJson()
}
