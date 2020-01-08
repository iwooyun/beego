package controllers

import e "beego/app/recover"

type ErrorController struct {
	BaseController
}

// Error404 请求的路由不存在.
func (c *ErrorController) Error404() {
	c.Failure(e.RouterNotFound, "no router match", "路由不存在")
}

// Error401 验签未通过.
func (c *ErrorController) Error401() {
	c.Failure(e.SignatureError, "sign valid fail", "签名验证失败")
}

// Error500 服务内部错误.
func (c *ErrorController) Error500() {
	c.Failure(e.SystemError, "system err", "服务内部错误")
}
