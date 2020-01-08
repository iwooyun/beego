package models

import (
	"beego/models/po"
	"github.com/astaxie/beego/orm"
	"reflect"
)

var CityAllDaoImpl ICityAllDao

type ICityAllDao interface {
	IBaseDao
}

type CityAllDao struct {
	BaseDao
}

type CityAll struct {
	po.CityAllPo
}

func NewCityAllDao() ICityAllDao {
	if CityAllDaoImpl == nil {
		CityAllDaoImpl = &CityAllDao{BaseDao{EntityType: reflect.TypeOf(new(CityAll)).Elem()}}
	}
	return CityAllDaoImpl
}

func (t *CityAll) GetTableName() string {
	return "city_all"
}

func (t *CityAll) GetDBName() string {
	return "default"
}

var CityAllReturnFields = []string{
	"Id",
	"Areaid",
	"BigAreaid",
	"Provinceid",
	"Cityname",
	"Ename",
	"Service",
	"Near",
	"TianrunCode",
	"Parentid",
	"CityRank",
	"Longitude",
	"Latitude",
	"CityCode",
}

func init() {
	orm.RegisterModel(new(CityAll))
}
