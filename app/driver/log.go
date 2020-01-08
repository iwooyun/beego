package driver

import (
	"beego/app/driver/adapter"
	"beego/config"
	"beego/library/notice"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"path/filepath"
	"strconv"
)

type LogDriver struct {
	BaseDriver
	SmtpConfig string
}

// OptionInit 邮件引擎注册到系统日志模块.
func (d *LogDriver) Make() error {
	err := logs.SetLogger(adapter.LogMailAdapter, d.SmtpConfig)
	return err
}

// OptionInit 获取邮件引擎配置.
func (d *LogDriver) OptionInit() error {
	mailConfig, err := notice.GetMailConfig(true)
	d.SmtpConfig = string(mailConfig.([]uint8))
	return err
}

// Reload implementing method. empty.
func (d *LogDriver) Reload(iDriver IConfigurationDriver) error {
	return nil
}

// SetLogger 系统日志模块设置.
func SetLogger() {
	beeLogger := logs.Async(1e2)
	beeLogger.SetLevel(logs.LevelInfo)
	delConsole(beeLogger)

	logPath := filepath.Join(string(os.PathSeparator), config.LogFilePath, config.AppEnName+"."+beego.BConfig.RunMode+".log")
	setLogFile(logPath, beeLogger)
}

// SetLogger 禁止控制台输出.
func delConsole(beeLogger *logs.BeeLogger) {
	_ = beeLogger.DelLogger(logs.AdapterConsole)
}

// setLogFile 设置日志输出路径.
func setLogFile(logPath string, beeLogger *logs.BeeLogger) {
	logFileConfig := `{"filename":"` + logPath + `","level":` + strconv.Itoa(logs.LevelInfo) + `,"separate":[` + config.SeparateLevel + `]}`
	if err := beeLogger.SetLogger(adapter.MultiFileAdapter, logFileConfig); err != nil {
		fmt.Printf("日志模块初始化失败，err => [%s]", err)
	}
}
