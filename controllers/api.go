package controllers

import (
	"github.com/astaxie/beego"
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
	var res struct {
		Url string `json:"url"`
	}
	res.Url = url
	c.Data["json"] = &res
	c.ServeJson()
}
