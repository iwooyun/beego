package po

import "time"

type CityPo struct {
	Id                 int       `orm:"column(cityid);auto" description:"城市id" json:"cityid"`
	Areaid             int8      `orm:"column(areaid)" description:"大区id" json:"areaid"`
	BigAreaid          int8      `orm:"column(big_areaid)" description:"后台使用的大区划分" json:"big_areaid"`
	Provinceid         int       `orm:"column(provinceid)" description:"省份id" json:"provinceid"`
	Cityname           string    `orm:"column(cityname);size(255)" description:"城市中文名" json:"cityname"`
	Ename              string    `orm:"column(ename);size(255)" description:"城市拼音名" json:"ename"`
	Shortname          string    `orm:"column(shortname);size(4)"  json:"shortname"`
	Service            int8      `orm:"column(service)" description:"开通服务[0无1半价车]" json:"service"`
	Near               string    `orm:"column(near);size(100)" description:"附近城市" json:"near"`
	TianrunCode        string    `orm:"column(tianrun_code);size(10)" description:"天润城市对应编码" json:"tianrun_code"`
	Zhigou             int8      `orm:"column(zhigou)" description:"是否开通直购（0、未开通；1、全部开通；2、只开通迁入；3、只开通迁出）" json:"zhigou"`
	IsVisit            int8      `orm:"column(is_visit)" description:"是否上门服务（-1 否，1是）" json:"is_visit"`
	Longitude          float64   `orm:"column(longitude);digits(10);decimals(7)" description:"经度" json:"longitude"`
	Latitude           float64   `orm:"column(latitude);digits(10);decimals(7)" description:"纬度" json:"latitude"`
	CityRank           int8      `orm:"column(city_rank);null" description:"中心城=1，主城=10，百城=100" json:"city_rank"`
	CityGroup          uint8     `orm:"column(city_group)" description:"团发大区[1东北大区,2华北大区,3华东大区,4华南大区,5华中大区,6西北大区,7西南大区]" json:"city_group"`
	IsGoldPartner      int8      `orm:"column(is_gold_partner)" description:"金牌合伙人:[1:支持,-1:不支持]" json:"is_gold_partner"`
	DirectRentSupport  int8      `orm:"column(direct_rent_support)" description:"是否支持直租：1支持，0不支持" json:"direct_rent_support"`
	SalvagedSupport    int8      `orm:"column(salvaged_support)" description:"捞回方案1支持0不支持" json:"salvaged_support"`
	IsshowC            int8      `orm:"column(isshow_c)" description:"是否C端展示 1：是，-1：否。供C端/CC取用" json:"isshow_c"`
	IsLeaseBack        int8      `orm:"column(is_lease_back)" description:"是否支持回租背户方案1支持0不支持" json:"is_lease_back"`
	MortgageServiceFee int       `orm:"column(mortgage_service_fee)" description:"抵押服务费(分)" json:"mortgage_service_fee"`
	IsSmallPubHouse    uint8     `orm:"column(is_small_pub_house)" description:"是否支持小公户：1支持，0不支持" json:"is_small_pub_house"`
	IsWzMortgage       int8      `orm:"column(is_wz_mortgage)" description:"是否微众抵押城市 0否，1是" json:"is_wz_mortgage"`
	IsPurchaseDirect   int8      `orm:"column(is_purchase_direct)" description:"是否直购直租 0否 1是" json:"is_purchase_direct"`
	IsPurchaseOrigin   int8      `orm:"column(is_purchase_origin)" description:"是否支持全国直购回租 0 不支持，1直租" json:"is_purchase_origin"`
	IsMsEx             int8      `orm:"column(is_ms_ex)" description:"是否民生银行展业城市：0否，1是" json:"is_ms_ex"`
	IsMsTrans          int8      `orm:"column(is_ms_trans)" description:"是否开启民生银行转单新网：0否，1是" json:"is_ms_trans"`
	Updatetime         time.Time `orm:"column(updatetime);type(timestamp)" description:"更新时间" json:"updatetime"`
}
