package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type DateRangeField struct {
	stringx.DateRange
}

func (req *DateRangeField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return req.Parse(data)
}
