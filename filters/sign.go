package filters

// 接口验签规则过滤器.
// 验签必传参数：
// 		业务线标识	app_id
// 		请求时间戳	ts
// 		验签加密串	sign

import (
	"beego/library/utils/base"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/parkingwang/go-sign"
	"github.com/zouyx/agollo"
	"strconv"
	"strings"
	"time"
)

// Verify 验签过滤器
func Verify(ctx *context.Context) {
	var (
		err       error
		signer    *sign.GoSigner
		debugFlag bool
	)
	verifier := sign.NewGoVerifier()
	verifier.DefaultKeyName.SetKeyNameAppId("app_id")
	verifier.DefaultKeyName.SetKeyNameTimestamp("ts")
	if err = ctx.Request.ParseForm(); err != nil {
		goto Error401
	}

	verifier.ParseValues(ctx.Request.Form)
	debugFlag, _ = strconv.ParseBool(verifier.MustString("debug"))
	if !base.IsProduction() && debugFlag {
		return
	}

	verifier.SetTimeout(time.Minute * 5)
	if err := verifier.CheckTimeStamp(); nil != err {
		goto Error401
	}

	signer = sign.NewGoSignerMd5()
	signer.SetBody(verifier.GetBodyWithoutSign())
	signer.SetAppSecretWrapBody(agollo.GetStringValue("API_SECRET", ""))

	if strings.Compare(verifier.MustString("sign"), signer.GetSignature()) == 0 {
		return
	}

Error401:
	beego.Exception(401, ctx)
}
