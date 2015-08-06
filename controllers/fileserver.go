package controllers

import (
	"github.com/astaxie/beego"
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

// @router /download/*
func (c *FileServerController) Download() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	dstfilename := "./file/" + filename
	c.Ctx.Output.Download(dstfilename, filename)
}
