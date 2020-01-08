package models

import (
	"beego/models/po"
	"github.com/astaxie/beego/orm"
	"reflect"
)

var CityDaoImpl ICityDao

type ICityDao interface {
	IBaseDao
}

type CityDao struct {
	BaseDao
}

type City struct {
	po.CityPo
}

// city表要返回的字段
var CityReturnFields = []string{
	"Id",
	"Areaid",
	"BigAreaid",
	"Provinceid",
	"Cityname",
	"Ename",
	"Shortname",
	"Service",
	"Near",
	"TianrunCode",
	"Zhigou",
	"IsVisit",
	"Longitude",
	"Latitude",
	"CityRank",
	"CityGroup",
	"IsGoldPartner",
	"DirectRentSupport",
	"SalvagedSupport",
	"IsshowC",
	"IsLeaseBack",
	"MortgageServiceFee",
	"IsSmallPubHouse",
	"IsWzMortgage",
	"IsPurchaseDirect",
	"IsPurchaseOrigin",
	"IsMsEx",
	"IsMsTrans",
	"Updatetime",
}

func NewCityDao() ICityDao {
	if CityDaoImpl == nil {
		CityDaoImpl = &CityDao{BaseDao{EntityType: reflect.TypeOf(new(City)).Elem()}}
	}
	return CityDaoImpl
}

func (t *City) GetTableName() string {
	return "city"
}

func (t *City) GetDBName() string {
	return "default"
}

func init() {
	orm.RegisterModel(new(City))
}
