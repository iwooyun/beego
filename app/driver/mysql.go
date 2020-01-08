package driver

import (
	"beego/config"
	"beego/database"
	"beego/library/utils/base"
	"fmt"
	"github.com/astaxie/beego"
	beeConfig "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type MysqlDriver struct {
	BaseDriver
	beeConfig.IniConfig
	DbConfigs database.DatabaseConf
}

// Make 注册数据库驱动.
func (s *MysqlDriver) Make() (err error) {
	if !base.IsProduction() {
		orm.Debug = true
		logger := logs.NewLogger()
		logPath := filepath.Join(string(os.PathSeparator), config.LogFilePath, config.AppEnName+"."+beego.BConfig.RunMode+".sql.log")
		_ = logger.Async().SetLogger(logs.AdapterFile, `{"filename":"`+logPath+`","level":`+strconv.Itoa(logs.LevelInfo)+`}`)
		orm.DebugLog = orm.NewLog(logger)
	}

	for aliasName, dbConfig := range s.DbConfigs {
		dataSource := s.dataSourceFormat(dbConfig)
		if err := orm.RegisterDataBase(aliasName, "mysql", dataSource,
			dbConfig.MaxIdle, dbConfig.MaxOpen); err != nil {
			logs.Error("%s数据源注册失败，驱动连接信息：{%s}，错误信息：{%s}", aliasName, dataSource, err)
			continue
		}

		db, _ := orm.GetDB(aliasName)
		db.SetConnMaxLifetime(time.Duration(dbConfig.MaxIdle) * time.Second)

		logs.Informational("%s数据库注册成功", aliasName)
	}
	return nil
}

// OptionInit 获取数据库连接信息.
func (s *MysqlDriver) OptionInit() error {
	database.FillApolloConfig()
	s.DbConfigs = database.GetDatabaseConnConf()
	return nil
}

// Reload implementing method. empty.
func (s *MysqlDriver) Reload(iDriver IConfigurationDriver) error {
	return nil
}

// dataSourceFormat 数据源连接信息格式化
func (s *MysqlDriver) dataSourceFormat(dbConfig *database.Database) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=true&loc=Asia%%2fShanghai",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset,
		dbConfig.Collation,
	)
}
