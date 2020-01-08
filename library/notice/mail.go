package notice

// 邮件

import (
	"beego/config"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/zouyx/agollo"
	"gopkg.in/gomail.v2"
	"net"
	"strconv"
	"strings"
	"time"
)

type AdapterMailConfig struct {
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Host        string   `json:"host"`
	Subject     string   `json:"subject"`
	FromAddress string   `json:"fromAddress"`
	SendTos     []string `json:"sendTos"`
	ContentType string   `json:"contentType"`
	Level       uint8    `json:"level"`
}

// SendMail 发送邮件
func SendMail(title string, msg interface{}, mailTo []string, path string) bool {
	defer func() bool {
		err := recover()
		if err != nil {
			logs.Error(fmt.Errorf("邮件发送失败，错误信息 => {%s}", err))
			return false
		}
		return true
	}()

	mailConfig, _ := GetMailConfig(false)
	conf := mailConfig.(*AdapterMailConfig)
	conf.Subject = "【" + beego.BConfig.RunMode + "】" + config.AppZhName + "服务" + title

	if len(mailTo) == 0 {
		mailTo = conf.SendTos
	}

	if _, ok := msg.(string); ok != true {
		msg, _ = json.MarshalIndent(msg, "<br>", "&nbsp;&nbsp;&nbsp;&nbsp;")
	}
	content := fmt.Sprintf(getContentTpl(), config.AppZhName, msg, time.Now().Format("2006-01-02 15:04:05"))

	mail := gomail.NewMessage()
	mail.SetHeader("From", conf.FromAddress)
	mail.SetHeader("To", mailTo...)
	mail.SetHeader("Subject", conf.Subject)
	mail.SetBody(conf.ContentType, content)
	if path != "" {
		mail.Attach(path)
	}
	host, _port, _ := net.SplitHostPort(conf.Host)
	port, _ := strconv.Atoi(_port)
	driver := gomail.NewDialer(host, port, conf.Username, conf.Password)
	driver.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := driver.DialAndSend(mail); err != nil {
		panic(err)
	}

	return true
}

func GetMailConfig(isJsonResult bool) (interface{}, error) {
	conf := &AdapterMailConfig{
		Username:    agollo.GetStringValue("MAIL_ALERT_USER", ""),
		Password:    agollo.GetStringValue("MAIL_ALERT_PASS", ""),
		Host:        agollo.GetStringValue("MAIL_ALERT_HOST", "") + ":" + agollo.GetStringValue("MAIL_ALERT_PORT_SSL", ""),
		FromAddress: agollo.GetStringValue("MAIL_ALERT_USER", ""),
		SendTos:     strings.Split(agollo.GetStringValue("MAIL_ALERT_SEND", ""), ","),
		Subject:     "",
		Level:       logs.LevelWarning,
		ContentType: "text/html",
	}

	if isJsonResult {
		return json.Marshal(conf)
	}

	return conf, nil
}

func getContentTpl() string {
	return `
		<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml">
		　<head>
		　　<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		　　<title>%s邮件</title>
		　　<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		　</head>
		<body style="margin: 0; padding: 0;">
			<div><span>邮件内容 => <br></span>%s<br></div>
		</body>
		<footer>
		  <p style="text-align:right">邮件发送时间:%s</p>
		</footer>
		</html>
	`
}
