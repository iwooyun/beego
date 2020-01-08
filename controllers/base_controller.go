package controllers

import (
	e "beego/app/recover"
	"beego/controllers/validator"
	"beego/library/utils/base"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type BaseController struct {
	startTime time.Time
	beego.Controller
}

// 	ApiResponse Api响应结果集结构体.
type ApiResponse struct {
	Code uint16      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) Prepare() {
	c.startTime = time.Now()
	requestUri := c.Ctx.Input.URL()
	if _, ok := validator.GlobalControllerValidator[requestUri]; ok {
		err := validator.NewValidator(validator.GlobalControllerValidator[requestUri], c.Ctx.Input)
		if err != nil {
			c.Failure(e.ParamError, err.Error(), nil)
		}
	}
}

// SuccessWithoutData 返回data为空的成功信息.
func (c *BaseController) SuccessWithoutData() {
	c.JsonResult(e.Success, "suc", [...]string{})
}

// Success 返回成功信息.
func (c *BaseController) Success(data interface{}) {
	c.JsonResult(e.Success, "success", data)
}

// Failure 返回错误信息.
func (c *BaseController) Failure(code uint16, msg string, data interface{}) {
	c.JsonResult(code, msg, data)
}

// JsonResultWithOutData 返回json格式数据.
func (c *BaseController) JsonResultWithOutData(code uint16, msg string) {
	c.JsonResult(code, msg, [...]string{})
}

// JsonResult 返回json格式数据.
func (c *BaseController) JsonResult(code uint16, msg string, data interface{}) {
	if base.IsNil(data) {
		data = [...]string{}
	}
	c.Data["json"] = &ApiResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.ServeJSON()
}

func (c *BaseController) Finish() {
	var (
		requestTime time.Time
		elapsedTime float64
		r           = c.Ctx.Request
	)
	requestTime = c.startTime
	elapsedTime = time.Since(c.startTime).Seconds()

	statusCode := c.Ctx.ResponseWriter.Status
	if statusCode == 0 {
		statusCode = 200
	}

	record := map[string]interface{}{
		"remote_addr":     c.Ctx.Input.IP(),
		"rs_time":         requestTime.Format("2006-06-02 03:04:05"),
		"request_method":  r.Method,
		"request_uri":     r.RequestURI,
		"server_protocol": r.Proto,
		"host":            r.Host,
		"status":          statusCode,
		"elapsed_time":    elapsedTime,
		"http_referrer":   r.Referer(),
		"http_user_agent": r.UserAgent(),
		"remote_user":     r.RemoteAddr,
		"post_data":       r.PostForm.Encode(),
		"response_data":   c.Data["json"],
	}

	msg, _ := json.Marshal(record)
	logs.Emergency(string(msg))
}
