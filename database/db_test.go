package database

// 多库配置.
// 1、添加新增数据库常量.
// 2、创建struct, 实现IDatabase接口.
// 3、数据库配置注册.

import (
	"github.com/zouyx/agollo"
)

// Test test库配置
type Test struct {
	Database
}

// NewTest 创建test库配置对象.
func NewTest() IDatabase {
	return &Test{}
}

// ConnConfInit 获取配置中心test库连接配置, 完成各配置项赋值.
func (t *Test) ConnConfInit(db *Database) *Database {
	db.Host = agollo.GetStringValue("DB_TEST_HOST", "")
	db.Port = agollo.GetStringValue("DB_TEST_PORT", "")
	db.User = agollo.GetStringValue("DB_TEST_USER", "")
	db.Pass = agollo.GetStringValue("DB_TEST_PASS", "")
	db.Name = agollo.GetStringValue("DB_TEST_NAME", "")
	return db
}

func init() {
	RegisterDatabase("default", NewTest)
}
