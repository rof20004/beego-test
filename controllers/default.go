package controllers

import (
	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get router options
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
