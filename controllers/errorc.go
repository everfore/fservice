package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Data["content"] = "user no auth."
	c.TplNames = "error.html"
}
