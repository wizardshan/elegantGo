package request

import (
	"encoding/json"
	"strings"
)

type StringBySeparatorV1 struct {
	Values []string
}

func (req *StringBySeparatorV1) UnmarshalJSON(b []byte) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	req.Values, err = req.split(data)
	return err
}

func (req *StringBySeparatorV1) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *StringBySeparatorV1) split(data string) (elements []string, err error) {
	separator := ","
	elements = strings.Split(data, separator)
	err = validate.Var(elements, "dive,required,alphanum")
	return
}
