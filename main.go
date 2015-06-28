package main

import (
	_ "bkand1909/song-getter/routers"
	"github.com/astaxie/beego"
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
