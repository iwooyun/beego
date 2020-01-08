package driver

import (
	beeCache "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/zouyx/agollo"
	"sync"
)

var (
	redisErr error
	Redis    beeCache.Cache
)

type RedisDriver struct {
	BaseDriver
	host    string
	port    string
	cliLock sync.Mutex
}

// Make 创建Redis连接对象.
func (s *RedisDriver) Make() error {
	Redis, redisErr = beeCache.NewCache("redis", `{"conn":"`+s.host+`:`+s.port+`"}`)
	return redisErr
}

// OptionInit Redis连接配置项初始化赋值.
func (s *RedisDriver) OptionInit() error {
	s.host = agollo.GetStringValue("SITE_REDIS_SERVER", "")
	s.port = agollo.GetStringValue("SITE_REDIS_PORT", "")

	return nil
}
