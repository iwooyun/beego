package driver

import (
	"beego/app/driver/adapter"
	"beego/config"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/zouyx/agollo"
	"os"
	"path/filepath"
)

type ApolloDriver struct {
	BaseDriver
	config agollo.AppConfig
}

// Make 从阿波罗配置中心获取项目所有配置.
func (c *ApolloDriver) Make() error {
	readyConfig := &c.config
	agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
		return readyConfig, nil
	})

	err := agollo.StartWithLogger(adapter.ApolloLogger{})
	if err != nil {
		logs.Error("apollo request err, content => %s", err.Error())
		return err
	}

	c.syncListenChangeEvent()

	return nil
}

// OptionInit 阿波罗连接配置初始化赋值.
func (c *ApolloDriver) OptionInit() error {
	c.config = agollo.AppConfig{
		AppId:            os.Getenv("APPID"),
		Cluster:          os.Getenv("CLUSTER"),
		NamespaceName:    os.Getenv("NAMESPACENAME"),
		Ip:               os.Getenv("IP"),
		BackupConfigPath: c.getBackupConfigPath(),
	}

	return nil
}

// Reload implementing method. empty.
func (c *ApolloDriver) Reload(iDriver IConfigurationDriver) error {
	return nil
}

// getBackupConfigPath 获取阿波罗配置备份路径
func (c *ApolloDriver) getBackupConfigPath() string {
	return filepath.Join(string(os.PathSeparator), config.ApolloBackConfigPath)
}

// syncListenChangeEvent 监听阿波罗配置变更事件，并重新加载驱动配置
func (c *ApolloDriver) syncListenChangeEvent() {
	event := agollo.ListenChangeEvent()
	go func() {
		for {
			select {
			case changeEvent := <-event:
				bytes, _ := json.Marshal(changeEvent)
				logs.Info("event:", string(bytes))
				ReBuild()
			}
		}
	}()
}
