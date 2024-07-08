package response

type Data any

type Response struct {
	Code    int
	Message string
	Success bool
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
	OK              = 0
	UnKnown         = 100
	InvalidArgument = 200
	BusinessError   = 500
)

var Msg = map[int]string{
	OK:              "OK",
	UnKnown:         "出现未知错误，请稍后重试",
	InvalidArgument: "参数错误",
	BusinessError:   "业务出现错误",
}
