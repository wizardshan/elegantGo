package request

import (
	"time"
)

type TimeRangeV10 struct {
	RangeV10[time.Time]
}

func (req *TimeRangeV10) parse(b []byte, validateTag string, parseTimes func() []time.Time) error {
	return req.RangeV10.parse(b, validateTag, parseTimes, req.validRange)
}

func (req *TimeRangeV10) validRange(start time.Time, end time.Time) bool {
	return start.Before(end)
}
