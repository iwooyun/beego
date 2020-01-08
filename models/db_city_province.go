package models

import (
	"beego/models/po"
	"github.com/astaxie/beego/orm"
	"reflect"
)

var CityProvinceDaoImpl ICityProvinceDao

type ICityProvinceDao interface {
	IBaseDao
}

type CityProvinceDao struct {
	BaseDao
}

type CityProvince struct {
	po.CityProvincePo
}

var CityProvinceReturnFields = []string{
	"Id",
	"Provincename",
	"Nameshort",
	"Allspell",
	"Sort",
	"BigAreaid",
}

func NewCityProvinceDao() ICityProvinceDao {
	if CityProvinceDaoImpl == nil {
		CityProvinceDaoImpl = &CityProvinceDao{BaseDao{EntityType: reflect.TypeOf(new(CityProvince)).Elem()}}
	}
	return CityProvinceDaoImpl
}

func (t *CityProvince) GetTableName() string {
	return "city_province"
}

func (t *CityProvince) GetDBName() string {
	return "default"
}

func init() {
	orm.RegisterModel(new(CityProvince))
}
