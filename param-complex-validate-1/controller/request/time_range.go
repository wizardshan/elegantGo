package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type TimeRangeField struct {
	stringx.TimeRange
}

func (req *TimeRangeField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return req.Parse(data)
}
