package response

type Data any

type Response struct {
	Code    int
	Message string
	Data    Data
}

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

const (
	OK = 0

	InternalError = 1000

	ParamError = 2000
)

var Msg = map[int]string{
	OK: "OK",

	InternalError: "服务器出现未知错误，请稍后重试",

	ParamError: "参数错误",
}
