package main

import (
	"github.com/bkand1909/song-getter/Godeps/_workspace/src/github.com/astaxie/beego"
	_ "github.com/bkand1909/song-getter/routers"
	"os"
	"strconv"
)

func configureStaticPath() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/font", "static/font")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/file", "static/file")
}

func main() {
	configureStaticPath()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	beego.HttpPort = port
	beego.Run()
}
