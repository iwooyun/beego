package driver

import (
	"github.com/astaxie/beego/logs"
	"sync"
)

// IConfigurationDriver 驱动配置接口.
type IConfigurationDriver interface {
	Make() error
	OptionInit() error
	Reload(iDriver IConfigurationDriver) error
}

// BaseDriver 驱动配置基础结构体.
type BaseDriver struct {
	CliLock sync.Mutex
}

// Reload 重新加载驱动.
func (s *BaseDriver) Reload(iDriver IConfigurationDriver) error {
	_ = iDriver.OptionInit()
	s.CliLock.Lock()
	if err := iDriver.Make(); err == nil {
		s.CliLock.Unlock()
		return err
	}
	s.CliLock.Unlock()
	return nil
}

var driverContainer *Container

// 驱动容器结构体.
type Container struct {
	Instances  map[string]IConfigurationDriver
	DriverName []string
}

// NewDriverContainer 创建驱动容器结构体.
func NewDriverContainer() *Container {
	return &Container{}
}

// Register 驱动注册.
func (c *Container) Register(name string, iDriver IConfigurationDriver) {
	if iDriver == nil {
		logs.Error("configs: Register provide is nil")
		panic("configs: Register provide is nil")
	}
	if _, dup := c.Instances[name]; dup {
		logs.Error("configs: Register called twice for provider " + name)
		panic("configs: Register called twice for provider " + name)
	}

	if c.Instances == nil {
		c.Instances = make(map[string]IConfigurationDriver)
	}

	c.Instances[name] = iDriver
	c.DriverName = append(c.DriverName, name)
}

// Build 加载所有驱动.
func Build(callback func(c *Container)) {
	driverContainer = NewDriverContainer()
	callback(driverContainer)

	for _, name := range driverContainer.DriverName {
		if iConfigurationDriver, ok := driverContainer.Instances[name]; ok {
			if err := iConfigurationDriver.OptionInit(); err != nil {
				logs.Error(name+"驱动配置初始化失败", err.Error())
			}
			if err := iConfigurationDriver.Make(); err != nil {
				logs.Error(name+"驱动加载启动失败", err.Error())
			}
		}
	}
}

// Build 重新加载所有驱动.
func ReBuild() {
	for _, name := range driverContainer.DriverName {
		if iConfigurationDriver, ok := driverContainer.Instances[name]; ok {
			if err := iConfigurationDriver.Reload(iConfigurationDriver); err != nil {
				logs.Error(name+"配置变更，驱动配置重新加载失败", err.Error())
			}
		}
	}
}
