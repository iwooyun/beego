package app

import (
	"beego/app/driver"
	_ "beego/app/recover"
	"beego/config"
	"github.com/astaxie/beego"
	"os"
)

// BcConfig全局配置初始化
func BcConfigInit() {
	beego.BConfig.RunMode = os.Getenv("ACTIVE")
	beego.BConfig.Listen.HTTPPort = config.HTTPPort
	beego.BConfig.AppName = config.AppName
	beego.BConfig.WebConfig.AutoRender = config.AutoRender
	beego.BConfig.CopyRequestBody = config.CopyRequestBody
	beego.BConfig.RecoverPanic = config.RecoverPanic
	beego.BConfig.WebConfig.EnableDocs = true
	if beego.BConfig.RunMode == beego.DEV {
		beego.BConfig.Listen.EnableAdmin = true
		beego.BConfig.Listen.AdminPort = 8088
	}
}

// 日志配置初始化
func LoggerInit() {
	driver.SetLogger()
}

// 驱动配置初始化
func configurationRegister() {
	driver.Build(func(c *driver.Container) {
		c.Register("apollo_cli", &driver.ApolloDriver{})
		c.Register("log_alert", &driver.LogDriver{})
		c.Register("mysql", &driver.MysqlDriver{})
		c.Register("redis_cli", &driver.RedisDriver{})
	})
}

func init() {
	BcConfigInit()
	LoggerInit()
	configurationRegister()
}
