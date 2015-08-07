package main

import (
	"github.com/astaxie/beego"
	"github.com/everfore/fservice/controllers"
	_ "github.com/everfore/fservice/routers"

	"fmt"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	beego.EnableAdmin = true
	beego.AdminHttpAddr = "localhost"
	beego.AdminHttpPort = 8088

	app := beego.Include(&controllers.FileServerController{})

	app.Run()
}

func init() {

	beego.SetStaticPath("/public", "static")

	toolbox.AddHealthCheck("database", &DatabaseCheck{})
	go a()
}

type DatabaseCheck struct {
}

func (c *DatabaseCheck) Check() error {
	return nil
}

func a() {
	tk := toolbox.NewTask("taska", "0/10 * * * * *", func() error { fmt.Println("hello world"); return nil })
	err := tk.Run()
	if err != nil {
		beego.Error(err)
	}
	toolbox.AddTask("taska", tk)
	toolbox.StartTask()
}
