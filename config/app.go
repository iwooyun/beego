package config

const (
	// beego全局配置
	AppName         = "beego" // 项目名称
	AutoRender      = false                 // 是否模板自动渲染
	CopyRequestBody = true                  // 是否允许在 HTTP 请求时，返回原始请求体数据字节
	RecoverPanic    = true                  // 是否异常恢复
	HTTPPort        = 8080                  // 应用监听端口

	// 服务名称
	AppEnName = "city-center"
	AppZhName = "城市中台"

	// apollo
	ApolloBackConfigPath = "/data/logs" // 配置备份路径

	// 日志模块配置
	SeparateLevel = `"access", "error"` // 日志目录级别配置， 支持"access", "alert", "critical", "error", "warning", "notice", "info", "debug"
	LogFilePath   = "/data/logs"        // 日志路径
)
