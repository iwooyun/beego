package validator

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
)

type CityValid struct {
	Validator
}

func NewCityValid() IValidator {
	return &CityValid{}
}

func (v CityValid) GetListByAreaId(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	areaId, _ := v.GetInt("area_id")
	valid.Required(areaId, "area_id")

	v.ErrorHandle(valid)
}

func (v CityValid) GetListByProvinceId(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	provinceId, _ := v.GetInt("ProvinceId")
	valid.Required(provinceId, "ProvinceId")

	v.ErrorHandle(valid)
}

func (v CityValid) GetListByEname(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	ename := v.GetString("ename")
	valid.Required(ename, "ename")

	v.ErrorHandle(valid)
}

func (v CityValid) Add(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	areaid, _ := v.GetInt("areaid")
	valid.Required(areaid, "areaid")
	provinceName := v.GetString("province_name")
	valid.Required(provinceName, "province_name")
	cityname := v.GetString("cityname")
	valid.Required(cityname, "cityname")
	ename := v.GetString("ename")
	valid.Required(ename, "ename")
	shortname, _ := v.GetInt("shortname")
	valid.Required(shortname, "shortname")
	longitude := v.GetString("longitude")
	valid.Required(longitude, "longitude")
	latitude := v.GetString("latitude")
	valid.Required(latitude, "latitude")
	cityCode := v.GetString("city_code")
	valid.Required(cityCode, "city_code")

	v.ErrorHandle(valid)
}

func (v CityValid) GetProfileByEname(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	ename := v.GetString("ename")
	valid.Required(ename, "ename")

	v.ErrorHandle(valid)
}

func (v CityValid) GetProfileById(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	cityid, _ := v.GetInt("cityid")
	valid.Required(cityid, "cityid")

	v.ErrorHandle(valid)
}

func (v CityValid) GetProfileByName(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	name := v.GetString("name")
	valid.Required(name, "name")

	v.ErrorHandle(valid)
}

func (v CityValid) Update(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	cityid, _ := v.GetInt("cityid")
	valid.Required(cityid, "cityid")
	areaid, _ := v.GetInt("areaid")
	valid.Required(areaid, "areaid")
	provinceName := v.GetString("province_name")
	valid.Required(provinceName, "province_name")
	cityname := v.GetString("cityname")
	valid.Required(cityname, "cityname")
	ename := v.GetString("ename")
	valid.Required(ename, "ename")
	shortname, _ := v.GetInt("shortname")
	valid.Required(shortname, "shortname")
	longitude := v.GetString("longitude")
	valid.Required(longitude, "longitude")
	latitude := v.GetString("latitude")
	valid.Required(latitude, "latitude")
	cityCode := v.GetString("city_code")
	valid.Required(cityCode, "city_code")

	v.ErrorHandle(valid)
}

func (v CityValid) GetAllPageList(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	query := v.GetString("query")
	valid.Required(query, "query")
	page, _ := v.GetUint16("page")
	valid.Required(page, "page")
	limit, _ := v.GetUint16("limit")
	valid.Required(limit, "limit")

	v.ErrorHandle(valid)
}

func init() {
	Register(CityValidator, NewCityValid)
}
