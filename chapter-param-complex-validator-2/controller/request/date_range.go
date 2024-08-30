package request

import (
	"elegantGo/chapter-param-complex-validator-2/pkg/stringx"
	"encoding/json"
	"time"
)

type DateRangeField string

func (req *DateRangeField) Values() (start time.Time, end time.Time, err error) {
	rg := new(stringx.DateRange)
	if err = rg.Parse(string(*req)); err != nil {
		return
	}
	start, end = rg.Start, rg.End
	return
}

type DateRangeJsonField struct {
	stringx.DateRange
}

func (req *DateRangeJsonField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return req.Parse(data)
}
