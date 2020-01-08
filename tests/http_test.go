package test

import (
	_ "beego/app"
	e "beego/app/recover"
	_ "beego/routers"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGet is a sample to run an endpoint test
func TestIndex(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test/mysql", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	t.Logf("TestIndex Request Code[%d], Response => \n%s ", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Code Should Be 20000", func() {
			var messageMap map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &messageMap)
			if err != nil {
				logs.Error("testing result trans map error, err content => {%s}", err)
			}
			So(messageMap["code"], ShouldEqual, e.Success)
		})
	})
}
