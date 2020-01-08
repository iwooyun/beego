package request

import (
	"github.com/astaxie/beego/httplib"
	"time"
)

type requestConfig struct {
	Url              string
	Method           string
	Params           map[string]string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	IsJson           bool
}

var defaultRequestConfig = &requestConfig{
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	IsJson:           false,
}

// Send 发送请求.
func Send(url, method string, params map[string]string, options ...interface{}) (result interface{}, err error) {
	defaultConfig := *defaultRequestConfig
	req := httplib.NewBeegoRequest(url, method)

	for k, v := range options {
		switch k {
		case 0:
			if vv, ok := v.(bool); ok {
				defaultConfig.IsJson = vv
			}
		case 1:
			if vv, ok := v.(time.Duration); ok {
				defaultConfig.ConnectTimeout = vv
			}
		case 2:
			if vv, ok := v.(time.Duration); ok {
				defaultConfig.ReadWriteTimeout = vv
			}
		}
	}

	req.SetTimeout(defaultConfig.ConnectTimeout, defaultConfig.ReadWriteTimeout)

	if !defaultConfig.IsJson {
		for key, value := range params {
			req.Param(key, value)
		}
		goto Request
	}

	_, err = req.JSONBody(params)
	if err != nil {
		return
	}

	//数据请求
Request:
	err = req.ToJSON(&result)
	return
}
