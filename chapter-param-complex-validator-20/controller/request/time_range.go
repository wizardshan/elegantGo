package request

import (
	"elegantGo/chapter-param-complex-validator-2/pkg/stringx"
	"encoding/json"
	"time"
)

type TimeRangeField string

func (req *TimeRangeField) Values() (start time.Time, end time.Time, err error) {
	rg := new(stringx.TimeRange)
	if err = rg.Parse(string(*req)); err != nil {
		return
	}
	start, end = rg.Start, rg.End
	return
}

type TimeRangeJsonField struct {
	stringx.TimeRange
}

func (req *TimeRangeJsonField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return req.Parse(data)
}