package request

import (
	"app/chapter-param-complex-validator-1/controller/pkg/separator"
	"encoding/json"
)

type JsonString struct {
	separator.StringV1
}

func (req *JsonString) unmarshal(b []byte) error {
	var data string
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	req.StringV1 = separator.NewV1(data, req.sep())
	return nil
}

func (req *JsonString) sep() string {
	return ","
}

func (req *JsonString) valid(tag string) error {
	return validate.Var(req.Values(), tag)
}
