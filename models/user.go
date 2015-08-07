package models

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	Name   string `form:"username"`
	Passwd string `form:"password"`
	XSRF   string `form:"_xsrf"`
}

func (u *User) Check(v *validation.Validation) {
	if u.Name == "shaalx" && u.Passwd == "1234" {
		v.Clear()
	} else {
		v.SetError("login", "user is abnormal.")
	}
}
