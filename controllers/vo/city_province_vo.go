package vo

type CityProvinceSingleVo struct {
	ApiResponse
	Data CityProvinceVo `json:"data"`
}

type CityProvinceListVo struct {
	ApiResponse
	Data []CityProvinceVo `json:"data"`
}

type CityProvincePageListVo struct {
	ApiResponse
	Data CityProvincePageList `json:"data"`
}

type CityProvincePageList struct {
	Pagination
	List []CityProvinceVo `json:"list"`
}

type CityProvinceVo struct {
	Id           int    `json:"provinceid" description:"省份id"`
	Provincename string `json:"provincename" description:"省份名"`
	Nameshort    string `json:"nameshort" description:"省份简称"`
	Allspell     string `json:"allspell" description:"全拼"`
	Sort         uint   `json:"sort" description:"排序"`
	BigAreaid    int    `json:"big_areaid" description:"后台使用大区ID"`
}
