package gin

import "github.com/asaskevich/govalidator"

type ParamError struct {
	err string
}

func (e ParamError) Error() string {
	return e.err
}

func NewParamError(err string) ParamError {
	return ParamError{err: err}
}

type Validator interface {
	Validate() error
}

func (c *Context) Validate(data interface{}) error {

	if err := c.ShouldBind(data); err != nil {
		return NewParamError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(data); err != nil {
		return NewParamError(err.Error())
	}

	if v, ok := data.(Validator); ok {
		if err := v.Validate(); err != nil {
			return NewParamError(err.Error())
		}
	}

	return nil
}
