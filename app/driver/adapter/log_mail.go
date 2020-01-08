package adapter

import (
	"beego/library/notice"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

const LogMailAdapter = "mailAdapter"

type MailAdapter struct {
	MailConfig *notice.AdapterMailConfig
}

// Init 邮件引擎配置初始化.
func (s *MailAdapter) Init(jsonConfig string) error {
	return json.Unmarshal([]byte(jsonConfig), &s.MailConfig)
}

// WriteMsg 日志信息发送告警平台.
func (s *MailAdapter) WriteMsg(when time.Time, msg string, level int) error {
	if level > int(s.MailConfig.Level) || level == logs.LevelEmergency {
		return nil
	}
	content := fmt.Sprintf("【告警信息】=> {\n\t详细内容：%s \n\t异常时间：%s \n}", msg, when.Format("2006-01-02 15:04:05"))
	notice.SendAlarm("系统日志告警", content, s.MailConfig.SendTos)
	return nil
}

// Flush implementing method. empty.
func (s *MailAdapter) Flush() {
}

// Destroy implementing method. empty.
func (s *MailAdapter) Destroy() {
}

// NewLoggerAdapter 创建MailAdapter, 返回LoggerInterface.
func NewLoggerAdapter() logs.Logger {
	return &MailAdapter{}
}

func init() {
	logs.Register(LogMailAdapter, NewLoggerAdapter)
}
