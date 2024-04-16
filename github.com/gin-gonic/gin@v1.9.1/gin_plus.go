package gin

type ParamError struct {
	err string
}

func (e ParamError) Error() string {
	return e.err
}

func NewParamError(err string) ParamError {
	return ParamError{err: err}
}
