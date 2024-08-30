package request

import (
	"time"
)

type TimeRange struct {
	Range[time.Time]
}

func (req *TimeRange) Parse(b []byte, validateTag string, mapperFunc func([]string) []time.Time) error {
	return req.Range.Parse(b, validateTag, mapperFunc, req.valid)
}

func (req *TimeRange) valid(start time.Time, end time.Time) bool {
	return start.Before(end)
}
