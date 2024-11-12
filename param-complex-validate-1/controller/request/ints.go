package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type IntsField struct {
	stringx.Ints
}

func (req *IntsField) UnmarshalJSON(b []byte) (err error) {
	var data string
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}
	return req.Parse(data)
}

type IntRangeField struct {
	stringx.IntRange
}

func (req *IntRangeField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return req.Parse(data)
}
