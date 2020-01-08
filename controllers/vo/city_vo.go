package vo

import "time"

type CitySingleVo struct {
	ApiResponse
	Data CityVo `json:"data"`
}

type CityListVo struct {
	ApiResponse
	Data []CityVo `json:"data"`
}

type CityPageListVo struct {
	ApiResponse
	Data CityPageList `json:"data"`
}

type CityPageList struct {
	Pagination
	List []CityVo `json:"list"`
}

type CityVo struct {
	Id                 int       `json:"cityid" description:"城市id"`
	Areaid             int8      `json:"areaid" description:"大区id"`
	BigAreaid          int8      `json:"big_areaid" description:"后台使用的大区划分"`
	Provinceid         int       `json:"provinceid" description:"省份id"`
	Cityname           string    `json:"cityname" description:"城市中文名"`
	Ename              string    `json:"ename" description:"城市拼音名"`
	Shortname          string    `json:"shortname"`
	Service            int8      `json:"service" description:"开通服务[0无1半价车]"`
	Near               string    `json:"near" description:"附近城市"`
	TianrunCode        string    `json:"tianrun_code" description:"天润城市对应编码"`
	Zhigou             int8      `json:"zhigou" description:"是否开通直购（0、未开通；1、全部开通；2、只开通迁入；3、只开通迁出）"`
	IsVisit            int8      `json:"is_visit" description:"是否上门服务（-1 否，1是）"`
	Longitude          float64   `json:"longitude" description:"经度"`
	Latitude           float64   `json:"latitude" description:"纬度"`
	CityRank           int8      `json:"city_rank" description:"中心城=1，主城=10，百城=100"`
	CityGroup          uint8     `json:"city_group" description:"团发大区[1东北大区,2华北大区,3华东大区,4华南大区,5华中大区,6西北大区,7西南大区]"`
	IsGoldPartner      int8      `json:"is_gold_partner" description:"金牌合伙人:[1:支持,-1:不支持]"`
	DirectRentSupport  int8      `json:"direct_rent_support" description:"是否支持直租：1支持，0不支持"`
	SalvagedSupport    int8      `json:"salvaged_support" description:"捞回方案1支持0不支持"`
	IsshowC            int8      `json:"isshow_c" description:"是否C端展示 1：是，-1：否。供C端/CC取用"`
	IsLeaseBack        int8      `json:"is_lease_back" description:"是否支持回租背户方案1支持0不支持"`
	MortgageServiceFee int       `json:"mortgage_service_fee" description:"抵押服务费(分)"`
	IsSmallPubHouse    uint8     `json:"is_small_pub_house" description:"是否支持小公户：1支持，0不支持"`
	IsWzMortgage       int8      `json:"is_wz_mortgage" description:"是否微众抵押城市 0否，1是"`
	IsPurchaseDirect   int8      `json:"is_purchase_direct" description:"是否直购直租 0否 1是"`
	IsPurchaseOrigin   int8      `json:"is_purchase_origin" description:"是否支持全国直购回租 0 不支持，1直租"`
	IsMsEx             int8      `json:"is_ms_ex" description:"是否民生银行展业城市：0否，1是"`
	IsMsTrans          int8      `json:"is_ms_trans" description:"是否开启民生银行转单新网：0否，1是"`
	Updatetime         time.Time `json:"updatetime" description:"更新时间"`
}
