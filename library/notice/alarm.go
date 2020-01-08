package notice

// 告警平台.

import (
	"beego/config"
	"beego/library/request"
	"beego/library/utils/encrypt"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/zouyx/agollo"
	"strings"
)

const (
	AlarmTypeMail   = "1" // 邮件告警
	AlarmTypeWechat = "2" // 微信告警
)

type Configuration struct {
	TitlePrefix string `json:"titlePrefix"`
	ProjectId   string `json:"projectId"`
	AlarmUrl    string `json:"alarmUrl"`
	Secret      string `json:"secret"`
}

// SendAlarm 发送告警信息.
func SendAlarm(alarmTitle string, alarmContent interface{}, alarmTo []string) bool {
	defer func() bool {
		err := recover()
		if err != nil {
			logs.Error(fmt.Errorf("告警发送失败，错误信息 => {%s}", err))
			return false
		}
		return true
	}()
	conf := &Configuration{
		TitlePrefix: "【" + beego.BConfig.RunMode + "】" + config.AppZhName + "服务",
		ProjectId:   agollo.GetStringValue("ALARM_PROJECT_ID", ""),
		AlarmUrl:    agollo.GetStringValue("ALARM_URL", ""),
		Secret:      agollo.GetStringValue("ALARM_SECRET", ""),
	}

	if len(alarmTo) == 0 {
		mailTos := agollo.GetStringValue("ALARM_SEND_TO", "")
		alarmTo = strings.Split(mailTos, ",")
	}

	if _, ok := alarmContent.(string); ok != true {
		alarmContent, _ = json.MarshalIndent(alarmContent, "", "\t")
	}

	param := make(map[string]string)
	param["project_id"] = conf.ProjectId
	param["title"] = conf.TitlePrefix + alarmTitle
	param["content"] = alarmContent.(string)
	param["type"] = AlarmTypeMail
	param["receiver"] = strings.Join(alarmTo, ",")

	token := encrypt.CreateAlarmSign(param, conf.Secret)
	param["token"] = token
	result, err := request.Send(conf.AlarmUrl, "POST", param)

	// 告警平台失败，重新发邮件
	if err != nil || int(result.(map[string]interface{})["code"].(float64)) != 1 {
		logs.Error(result)
		SendMail(alarmTitle, alarmContent, alarmTo, "")
	}

	return true
}
