package recover

import (
	"beego/library/notice"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"html/template"
	"runtime"
	"strconv"
	"time"
)

// 全局系统错误码常量定义.
const (
	//	系统公用50（22000--22050）
	Success         = 22000 //
	Failed          = 22001 //操作失败
	NoPermission    = 22002 //没有权限
	SystemError     = 22003 //系统错误
	RouterNotFound  = 22005 //路由不存在
	ParamError      = 22006 //参数错误
	ActionFrequent  = 22007 //操作过频繁
	IllegalRequest  = 22008 //非法请求
	JsonTransError  = 22009 //json解析失败
	SystemWaiting   = 22010 //操作等待
	ResponseError   = 22011 //返回数据不正确
	OutRequestError = 22012 //外部请求失败
	EnvError        = 22013 //环境变量有误
	SignatureError  = 22005 //API签名错误
	SignExprise     = 22016 //签名失效

)

// Recover 异常恢复处理.
func Recover() func(*context.Context) {
	return func(ctx *context.Context) {
		if err := recover(); err != nil {
			if err == beego.ErrAbort {
				return
			}
			if !beego.BConfig.RecoverPanic {
				panic(err)
			}
			if beego.BConfig.EnableErrorsShow {
				httpCode := fmt.Sprint(err)
				if _, ok := beego.ErrorMaps[httpCode]; ok {
					errCode, _ := strconv.ParseUint(httpCode, 10, 64)
					beego.Exception(errCode, ctx)
					return
				}
			}
			var stack string
			logs.Notice("the request url is ", ctx.Input.URL())
			logs.Notice("Handler crashed with error", err)
			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				logs.Notice(fmt.Sprintf("%s:%d", file, line))
				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
			}

			urlLog := fmt.Sprintln("the request url is " + ctx.Input.URL())
			urlLog += stack
			content := fmt.Sprintf("【错误信息】=> {\n\t错误描述：%s \n\t详细内容：%s \n\t异常时间：%s \n}",
				fmt.Sprint(err), urlLog, time.Now().Format("2006-01-02 15:04:05"))
			notice.SendAlarm("系统严重错误告警", content, []string{})
			if beego.BConfig.RunMode == beego.DEV && beego.BConfig.EnableErrorsRender {
				showError(err, ctx, stack)
				return
			}

			if ctx.Output.Status != 0 {
				ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
			} else {
				goto Error500
			}
		Error500:
			beego.Exception(500, ctx)
		}
	}
}

// showError 开发环境错误展示结果集.
func showError(err interface{}, ctx *context.Context, stack string) {
	t, _ := template.New("error.template").Parse(getTpl())
	data := map[string]string{
		"AppError":      fmt.Sprintf("%s:%v", beego.BConfig.AppName, err),
		"RequestMethod": ctx.Input.Method(),
		"RequestURL":    ctx.Input.URI(),
		"RemoteAddr":    ctx.Input.IP(),
		"Stack":         stack,
		"BeegoVersion":  beego.VERSION,
		"GoVersion":     runtime.Version(),
	}
	_ = t.Execute(ctx.ResponseWriter, data)
}

// getTpl 开发环境错误页面模板.
func getTpl() string {
	return `
		<!DOCTYPE html>
		<html>
		<head>
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
			<title>beego application error</title>
			<style>
				html, body, body * {padding: 0; margin: 0;}
				#header {background:#ffd; border-bottom:solid 2px #A31515; padding: 20px 10px;}
				#header h2{ }
				#footer {border-top:solid 1px #aaa; padding: 5px 10px; font-size: 12px; color:green;}
				#content {padding: 5px;}
				#content .stack b{ font-size: 13px; color: red;}
				#content .stack pre{padding-left: 10px;}
				table {}
				td.t {text-align: right; padding-right: 5px; color: #888;}
			</style>
			<script type="text/javascript">
			</script>
		</head>
		<body>
			<div id="header">
				<h2>{{.AppError}}</h2>
			</div>
			<div id="content">
				<table>
					<tr>
						<td class="t">Request Method: </td><td>{{.RequestMethod}}</td>
					</tr>
					<tr>
						<td class="t">Request URL: </td><td>{{.RequestURL}}</td>
					</tr>
					<tr>
						<td class="t">RemoteAddr: </td><td>{{.RemoteAddr }}</td>
					</tr>
				</table>
				<div class="stack">
					<b>Stack</b>
					<pre>{{.Stack}}</pre>
				</div>
			</div>
			<div id="footer">
				<p>beego {{ .BeegoVersion }} (beego framework)</p>
				<p>golang version: {{.GoVersion}}</p>
			</div>
		</body>
		</html>
	`
}

func init() {
	beego.BConfig.RecoverFunc = Recover()
}
