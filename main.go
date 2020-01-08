package main

import (
	_ "beego/app"
	_ "beego/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == beego.DEV {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
