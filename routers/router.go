package routers

import (
	"github.com/astaxie/beego"
	"github.com/bkand1909/song-getter/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/", &controllers.ApiController{})
}
