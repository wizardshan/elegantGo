package response

type Data any

type Resp struct {
	Code    int
	Message string
	Success bool
	Data    any
}
