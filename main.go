package main

import (
	"github.com/astaxie/beego"
	"github.com/everfore/fservice/controllers"
	_ "github.com/everfore/fservice/routers"
)

func main() {
	app := beego.Include(&controllers.FileServerController{})

	app.Run()
}

func init() {
	beego.SetStaticPath("/public", "static")
}
