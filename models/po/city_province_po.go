package po

type CityProvincePo struct {
	Id           int    `orm:"column(provinceid);auto" description:"省份id" json:"provinceid"`
	Provincename string `orm:"column(provincename);size(32)" description:"省份名" json:"provincename"`
	Nameshort    string `orm:"column(nameshort);size(1)" description:"省份简称" json:"nameshort"`
	Allspell     string `orm:"column(allspell);size(32)" description:"全拼" json:"allspell"`
	Sort         uint   `orm:"column(sort)" description:"排序" json:"sort"`
	BigAreaid    int8   `orm:"column(big_areaid)" description:"后台使用大区ID" json:"big_areaid"`
}
