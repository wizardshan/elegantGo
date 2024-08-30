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
	ArgumentInvalid = 200
	BizError        = 500
)

var Msg = map[int]string{
	OK:              "OK",
	UnKnown:         "出现未知错误，请稍后重试",
	ArgumentInvalid: "参数无效",
	BizError:        "业务出现错误",
}
