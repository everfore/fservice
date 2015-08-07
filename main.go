package main

import (
	"github.com/astaxie/beego"
	"github.com/everfore/fservice/controllers"
	_ "github.com/everfore/fservice/routers"

	"fmt"
	// "github.com/astaxie/beego/session"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/toolbox"
	"github.com/everfore/fservice/models"
)

func main() {
	app := beego.Include(&controllers.FileServerController{})
	beego.ErrorController(&controllers.ErrorController{})
	// beego.InsertFilter("/*", beego.BeforeRouter, CheckLogin)
	app.Run()
}

// var GlobalSessions *session.Manager

func init() {
	beego.SetStaticPath("/public", "static")

	// toolbox_init()
	// session_init()
}

func session_init() {
	// beego.SessionOn = true
	beego.SessionProvider = "file"
	// beego.SessionName = "begosessionid"
	// beego.SessionGCMaxLifetime = 3600
	// beego.SessionCookieLifeTime = 3600
	beego.SessionSavePath = "./sessionpath"
	// GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
}

func toolbox_init() {
	// toolbox healthcheck task
	beego.EnableAdmin = true
	beego.AdminHttpAddr = "localhost"
	beego.AdminHttpPort = 8088
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

func CheckLogin(ctx *context.Context) {
	if ctx.Request.RequestURI != "/login" {
		sess, err := models.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
		if err != nil || sess == nil {
			ctx.Abort(401, "session nil")
		}
		sessioner := sess.Get("gosessionkey")
		beego.Debug("session:", sess)
		beego.Debug("check login gosessionkey:", sessioner)
		if fmt.Sprintf("%v", sessioner) != "beego1234" {
			ctx.Redirect(401, "/login")
		}
	}
}
