package request

import (
	"github.com/elliotchance/pie/v2"
	"time"
)

type TimeRangeV9 struct {
	RangeV9[time.Time]
}

func (req *TimeRangeV9) Parse(b []byte, validateTag string, mapperFunc func([]string) []time.Time) error {
	return req.RangeV9.Parse(b, validateTag, mapperFunc, req.valid)
}

func (req *TimeRangeV9) mapper(elements []string, layout string) []time.Time {
	return pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func (req *TimeRangeV9) valid(start time.Time, end time.Time) bool {
	return start.Before(end)
}
