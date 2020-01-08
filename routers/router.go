// @APIVersion 1.0.0
// @Host 127.0.0.1:8080
// @Title 城市中台服务
// @Description 提供城市相关信息新增、查询及修改服务
// @Name admin@mail.com
package routers

import (
	"beego/controllers"
	"beego/filters"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/api",
		beego.NSBefore(filters.Verify),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		beego.NSNamespace("/city-base",
			beego.NSInclude(
				&controllers.CityAllController{},
			),
		),
        beego.NSNamespace("/city",
			beego.NSInclude(
				&controllers.CityController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
