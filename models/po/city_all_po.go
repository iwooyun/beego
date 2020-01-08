package po

type CityAllPo struct {
	Id          int     `orm:"column(cityid);auto" description:"城市ID" json:"cityid"`
	Areaid      int     `orm:"column(areaid)"  json:"areaid"`
	BigAreaid   uint    `orm:"column(big_areaid)" description:"大区ID" json:"big_areaid"`
	Provinceid  uint    `orm:"column(provinceid)" description:"所属省份id" json:"provinceid"`
	Cityname    string  `orm:"column(cityname);size(32)" description:"城市名称" json:"cityname"`
	Ename       string  `orm:"column(ename);size(32)" description:"城市名称拼音" json:"ename"`
	Service     int8    `orm:"column(service)" description:"开通服务：0无，1半价车" json:"service"`
	Near        string  `orm:"column(near);size(100)" description:"周边城市ID" json:"near"`
	TianrunCode string  `orm:"column(tianrun_code);size(10)" description:"天润城市对应编码" json:"tianrun_code"`
	Parentid    int     `orm:"column(parentid)" description:"上级id" json:"parentid"`
	CityRank    int8    `orm:"column(city_rank);null" description:"中心城=1，主城=10，百城=100" json:"city_rank"`
	Longitude   float64 `orm:"column(longitude);digits(10);decimals(7)" description:"经度" json:"longitude"`
	Latitude    float64 `orm:"column(latitude);digits(10);decimals(7)" description:"纬度" json:"latitude"`
	CityCode    string  `orm:"column(city_code);size(10)" description:"城市代码" json:"city_code"`
}
