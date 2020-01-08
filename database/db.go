package database

import "github.com/astaxie/beego/logs"

type IDatabase interface {
	ConnConfInit(db *Database) *Database
}

type databaseFunc func() IDatabase

var databaseConfAdapters = make(map[string]databaseFunc)

type DatabaseConf map[string]*Database

var databaseConfMap = make(DatabaseConf)

type Database struct {
	Name      string `json:"user"`
	Host      string `json:"pass"`
	Port      string `json:"host"`
	User      string `json:"port"`
	Pass      string `json:"name"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
	MaxIdle   int    `json:"max_idle"`
	MaxOpen   int    `json:"max_open"`
	MaxLife   int    `json:"max_life"`
}

// RegisterDatabase 配置连接数据库注册.
func RegisterDatabase(alias string, database databaseFunc) {
	if database == nil {
		panic("Database: Register provide is nil")
	}

	if _, dup := databaseConfAdapters[alias]; dup {
		logs.Error("configs: Register provide is nil")
		panic("databaseConf: Register called twice for provider " + alias)
	}

	databaseConfAdapters[alias] = database
}

// NewDatabase 创建数据库连接配置对象.
func NewDatabase() *Database {
	return &Database{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_general_ci",
		MaxIdle:   30, // 设置数据库的最大空闲连接
		MaxOpen:   30, // 设置数据库的最大数据库连接
		MaxLife:   50,
	}
}

// FillApolloConfig 从配置中心获取数据填充数据库连接配置.
func FillApolloConfig() {
	for alias, databaseConfAdapter := range databaseConfAdapters {
		databaseConfMap[alias] = databaseConfAdapter().ConnConfInit(NewDatabase())
	}
}

// GetDatabaseConnConf 获取数据库连接配置.
func GetDatabaseConnConf() DatabaseConf {
	return databaseConfMap
}
