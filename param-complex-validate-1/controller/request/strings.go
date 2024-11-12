package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type StringsField struct {
	stringx.Strings
}

func (req *StringsField) UnmarshalJSON(b []byte) (err error) {
	var data string
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}
	return req.Parse(data)
}
