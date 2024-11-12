package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type DateTimeRangeField struct {
	stringx.DateTimeRange
}

func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return req.Parse(data)
}
