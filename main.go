package main

import (
	"github.com/astaxie/beego"
	_ "github.com/bkand1909/song-getter/routers"
)

func configureStaticPath() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/font", "static/font")
	beego.SetStaticPath("/img", "static/img")
}

func main() {
	configureStaticPath()
	beego.Run()
}
