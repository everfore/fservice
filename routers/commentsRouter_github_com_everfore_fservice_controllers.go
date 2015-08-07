package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"LoadUpload",
			`/upload`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Upload",
			`/upload`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Show",
			`/show`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"LoadFile",
			`/loadfile/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Download",
			`/download/*`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"PostLogin",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"] = append(beego.GlobalControllerRouter["github.com/everfore/fservice/controllers:FileServerController"],
		beego.ControllerComments{
			"Error",
			`/error`,
			[]string{"get"},
			nil})

}
