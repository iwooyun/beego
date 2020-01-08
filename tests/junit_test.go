package test

import (
	_ "beego/app"
	"beego/models"
	_ "beego/routers"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService(t *testing.T) {
	where := make(map[string]interface{})
	where["cityid__in"] = []int{101, 102}
	var cityList []models.City
	fields := []string{"cityid", "areaid", "provinceid", "cityname"}
	pagination, _ := models.NewCityDao().SelectPageList(&cityList, where, fields, nil, nil, 1, 20)
	pagination.List = cityList
	msg, _ := json.MarshalIndent(pagination, "", "  ")
	t.Logf("TestService output content=> %s", msg)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(len(cityList), ShouldNotEqual, 0)
		})
	})
}
