package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "GetListByBigAreaId",
            Router: `/get-list-by-bigareaid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "GetListByProvinceId",
            Router: `/get-list-by-provinceid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "GetProfile",
            Router: `/get-profile`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "GetProfileByCode",
            Router: `/get-profile-by-code`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "GetProfileByEname",
            Router: `/get-profile-by-ename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityAllController"] = append(beego.GlobalControllerRouter["beego/controllers:CityAllController"],
        beego.ControllerComments{
            Method: "SearchListByEname",
            Router: `/search-list-by-ename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetAllPageList",
            Router: `/get-all-page-list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetListByAreaId",
            Router: `/get-list-by-area-id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetListByEname",
            Router: `/get-list-by-ename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetListByProvinceId",
            Router: `/get-list-by-province-id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetProfileByEname",
            Router: `/get-profile-by-ename`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetProfileById",
            Router: `/get-profile-by-id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetProfileByName",
            Router: `/get-profile-by-name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityController"] = append(beego.GlobalControllerRouter["beego/controllers:CityController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:CityProvinceController"] = append(beego.GlobalControllerRouter["beego/controllers:CityProvinceController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "Alarm",
            Router: `/alarm`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "Log",
            Router: `/log`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "Mail",
            Router: `/mail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "Mysql",
            Router: `/mysql`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego/controllers:TestController"] = append(beego.GlobalControllerRouter["beego/controllers:TestController"],
        beego.ControllerComments{
            Method: "Redis",
            Router: `/redis`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
