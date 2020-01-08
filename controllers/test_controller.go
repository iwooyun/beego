package controllers

import (
	"beego/app/driver"
	"beego/config"
	"beego/library/notice"
	"beego/models"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"log"
	"time"
)

// @TagName 基础测试
// @Description 框架基础功能测试模块
type TestController struct {
	BaseController
}

// URLMapping ...
func (c *TestController) URLMapping() {
	c.Mapping("Redis", c.Redis)
	c.Mapping("Mail", c.Mail)
	c.Mapping("Alarm", c.Alarm)
	c.Mapping("Log", c.Log)
	c.Mapping("Mysql", c.Mysql)
}

// @Title Redis
// @Summary Redis读写
// @Description Redis读写测试
// @Success 200 {object} vo.TestRedisVo
// @Failure 401 sign valid fail
// @router /redis [get]
func (c *TestController) Redis() {
	cacheKey := config.CacheTestKey
	err := driver.Redis.Put(cacheKey, time.Now().Format("2006-01-02 15:04:05"), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	c.Success(cache.GetString(driver.Redis.Get(cacheKey)))
}

// @Title Mail
// @Summary 发送邮件
// @Description 车源中心邮件测试
// @Success 200 {object} vo.TestSuccessVo
// @Failure 401 sign valid fail
// @router /mail [get]
func (c *TestController) Mail() {
	flag := notice.SendMail("车源中心邮件测试", "邮件功能测试", []string{"admin@mail.com"}, "")

	c.Success(map[string]bool{
		"is_success": flag,
	})
}

// @Title Alarm
// @Summary 发送告警
// @Description 车源中心告警测试
// @Success 200 {object} vo.TestSuccessVo
// @Failure 401 sign valid fail
// @router /alarm [get]
func (c *TestController) Alarm() {
	flag := notice.SendAlarm("车源中心告警测试", "告警平台测试", []string{"admin@mail.com"})

	c.Success(map[string]bool{
		"is_success": flag,
	})
}

// @Title Log
// @Summary 日志告警
// @Description 日志邮件告警测试
// @Success 200 {object} vo.TestSuccessVo
// @Failure 401 sign valid fail
// @router /log [get]
func (c *TestController) Log() {
	logs.Warning("日志邮件告警测试")
	c.Success(map[string]bool{
		"is_success": true,
	})
}

// @Title Mysql
// @Summary 数据库查询
// @Description 数据库查询测试
// @Success 200 {object} vo.ApiResponse
// @Failure 401 sign valid fail
// @router /mysql [get]
func (c *TestController) Mysql() {
	where := make(map[string]interface{})
	where["cityid__in"] = []int{101, 102}
	var cityList []models.City
	fields := []string{"cityid", "areaid", "provinceid", "cityname"}
	pagination, _ := models.NewCityDao().SelectPageList(&cityList, where, fields, nil, nil, 1, 20)
	pagination.List = cityList
	c.Success(pagination)
}
