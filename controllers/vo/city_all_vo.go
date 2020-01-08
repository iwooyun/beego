package vo

type CityAllSingleVo struct {
	ApiResponse
	Data CityAllVo `json:"data"`
}

type CityAllListVo struct {
	ApiResponse
	Data []CityAllVo `json:"data"`
}

type CityAllPageListVo struct {
	ApiResponse
	Data CityAllPageList `json:"data"`
}

type CityAllPageList struct {
	Pagination
	List []CityAllVo `json:"list"`
}

type CityAllVo struct {
	Id          int     `json:"cityid" description:"城市ID"`
	Areaid      int     `json:"areaid"`
	BigAreaid   uint    `json:"big_areaid" description:"大区ID"`
	Provinceid  uint    `json:"provinceid" description:"所属省份id"`
	Cityname    string  `json:"cityname" description:"城市名称"`
	Ename       string  `json:"ename" description:"城市名称拼音"`
	Service     int8    `json:"service" description:"开通服务：0无，1半价车"`
	Near        string  `json:"near" description:"周边城市ID"`
	TianrunCode string  `json:"tianrun_code" description:"天润城市对应编码"`
	Parentid    int     `json:"parentid" description:"上级id"`
	CityRank    int8    `json:"city_rank" description:"中心城=1，主城=10，百城=100"`
	Longitude   float64 `json:"longitude" description:"经度"`
	Latitude    float64 `json:"latitude" description:"纬度"`
	CityCode    string  `json:"city_code" description:"城市代码"`
}
