package response

type Data any

type Response struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息
	Data    Data   `json:"data"`    // 数据
}

const (
	OK = 0

	InternalError = 1000

	ParamError = 10000
)

var Msg = map[int]string{
	OK: "OK",

	InternalError: "服务器出现未知错误，请稍后重试",

	ParamError: "参数错误",
}
