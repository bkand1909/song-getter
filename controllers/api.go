package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bkand1909/song-getter/utils"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *ApiController) Post() {

}

func (c *ApiController) Get() {
	url := c.Input().Get("url")
	resp := utils.GetUrl(url, nil)
	html := resp.String()
	var parser utils.ZingParser
	_, album := parser.ToAlbum(url, html)
	c.Data["json"] = &album
	c.ServeJson()
}
