package request

import (
	"time"
)

type TimeRangeV9 struct {
	Start time.Time
	End   time.Time
	RangeV9[time.Time]
}

func (req *TimeRangeV9) unmarshalJSON(b []byte, validateTag string, parseFunc func(s string) time.Time) (err error) {
	req.Start, req.End, err = req.RangeV9.parse(b, validateTag, parseFunc, req.valid)
	return
}

func (req *TimeRangeV9) parseElement(s string, layout string) (t time.Time) {
	t, _ = time.Parse(layout, s)
	return
}

func (req *TimeRangeV9) valid(start time.Time, end time.Time) bool {
	return start.Before(end)
}
