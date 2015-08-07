package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/everfore/fservice/models"
	"html/template"
	"io/ioutil"
	"os"
)

type FileServerController struct {
	beego.Controller
}

func (c *FileServerController) Prepare() {
	// beego.EnableXSRF = false
	ip := c.Ctx.Input.IP()
	beego.Debug(ip)
	if !c.CheckLogin() {
		// c.Abort("401")
		c.Redirect("/login", 401)
		// c.Ctx.Redirect(401, "/login")
	}
}

// @router / [get]
func (c *FileServerController) Get() {
	c.TplNames = "index.html"
}

// @router /upload [get]
func (c *FileServerController) LoadUpload() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "upload.html"
}

// @router /upload [post]
func (c *FileServerController) Upload() {
	_, file, err := c.GetFile("filename")
	if nil == err {
		if serr := c.SaveToFile("filename", "./file/"+file.Filename); serr == nil {
			c.TplNames = "success.html"
		} else {
			beego.Error(serr)
			c.TplNames = "error.html"
		}
		return
	}
	beego.Error(err)
	c.TplNames = "error.html"
}

// @router /show [get]
func (c *FileServerController) Show() {
	files := models.Listfiles("./file")
	c.Data["File"] = files
	c.TplNames = "show.html"
}

// @router /loadfile/* [get]
func (c *FileServerController) LoadFile() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	if file, err := os.Open("./file/" + filename); err != nil {
		beego.Error(err)
		c.TplNames = "error.html"
	} else {
		if b, err := ioutil.ReadAll(file); err != nil {
			beego.Error(err)
			c.TplNames = "error.html"
		} else {
			c.Ctx.Output.Body(b)
		}
	}
}

// @router /download/* [get]
func (c *FileServerController) Download() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	dstfilename := "./file/" + filename
	c.Ctx.Output.Download(dstfilename, filename)
}

// @router /login [get]
func (c *FileServerController) Login() {
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "login.html"
}

// @router /login [post]
func (c *FileServerController) PostLogin() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		c.TplNames = "error.html"
		return
	}
	var usr models.User
	err = c.ParseForm(&usr)
	if err != nil {
		c.TplNames = "error.html"
		return
	}
	valid := validation.Validation{}
	// ok, err := valid.Valid(&usr)
	// if !ok || err != nil || valid.HasErrors() {
	// 	c.TplNames = "error.html"
	// 	return
	// }
	// beego.Notice(ok, err, valid.HasErrors())
	usr.Check(&valid)
	// beego.Notice(ok, err, valid.HasErrors())
	if valid.HasErrors() {
		c.TplNames = "error.html"
		return
	}
	sess.Set("gosessionkey", "beego1234")
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	beego.Debug(sess)
	// c.TplNames = "index.html"
	c.Redirect("/", 302)
}

// @router /logout [get]
func (c *FileServerController) Logout() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.TplNames = "error.html"
		return
	}
	sess.Set("gosessionkey", "")
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	beego.Debug(sess)
	// c.TplNames = "login.html"
	c.Redirect("/login", 302)
}

func (c *FileServerController) CheckLogin() bool {
	if c.Ctx.Request.RequestURI != "/login" {
		sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		if err != nil || sess == nil {
			return false
		}
		sessioner := sess.Get("gosessionkey")
		beego.Debug("session:", sess)
		beego.Debug("check login gosessionkey:", sessioner)
		if fmt.Sprintf("%v", sessioner) == "beego1234" {
			return true
		}
		return false
	}
	return true
}
