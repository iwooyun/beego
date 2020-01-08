package validator

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
)

type CityAllValid struct {
	Validator
}

func NewCityAllValid() IValidator {
	return &CityAllValid{}
}

func (v CityAllValid) GetProfileByEname(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	ename := v.GetString("ename")
	valid.Required(ename, "ename")

	v.ErrorHandle(valid)
}

func (v CityAllValid) GetListByProvinceId(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	provinceid, _ := v.GetInt("provinceid")
	valid.Required(provinceid, "provinceid")

	v.ErrorHandle(valid)
}

func (v CityAllValid) SearchListByEname(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	ename := v.GetString("ename")
	valid.Required(ename, "ename")

	v.ErrorHandle(valid)
}

func (v CityAllValid) GetProfileByCode(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	cityCode := v.GetString("city_code")
	valid.Required(cityCode, "city_code")

	v.ErrorHandle(valid)
}

func (v CityAllValid) GetListByBigAreaId(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	bigAreaid, _ := v.GetInt("big_areaid")
	valid.Required(bigAreaid, "big_areaid")

	v.ErrorHandle(valid)
}

func (v CityAllValid) GetProfile(input *context.BeegoInput) {
	valid := validation.Validation{}
	v.Input = input
	cityid, _ := v.GetInt("cityid")
	valid.Required(cityid, "cityid")

	v.ErrorHandle(valid)
}

func init() {
	Register(CityAllValidator, NewCityAllValid)
}
