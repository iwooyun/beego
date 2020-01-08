package vo

// 	ResponseVo Api响应结果集基础结构体.
type ApiResponse struct {
	Code uint16 `json:"code" description:"错误编号"`
	Msg  string `json:"msg" description:"错误信息"`
}

type Pagination struct {
	Page       uint16 `json:"page" description:"请求页数"`
	PageSize   uint16 `json:"page_size" description:"每页数量"`
	TotalCount uint64 `json:"total_count" description:"数据总数"`
}
