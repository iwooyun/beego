package vo

type TestSuccessVo struct {
	ApiResponse
	Data TestSuccessData `json:"data" description:"返回对象"`
}

type TestSuccessData struct {
	IsSuccess bool `json:"is_success" description:"是否成功"`
}

type TestRedisVo struct {
	ApiResponse
	Data TestRedisData `json:"data" description:"返回对象"`
}

type TestRedisData struct {
	value string `json:"value" description:"redis返回值"`
}
