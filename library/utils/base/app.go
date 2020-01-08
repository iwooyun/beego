package base

import "github.com/astaxie/beego"

const (
	// 测试环境
	FAT = "fat"
	// 预上线环境
	UAT = "uat"
)

func IsProduction() bool {
	if beego.BConfig.RunMode == FAT ||
		beego.BConfig.RunMode == UAT ||
		beego.BConfig.RunMode == beego.DEV {
		return false
	}

	return true
}
