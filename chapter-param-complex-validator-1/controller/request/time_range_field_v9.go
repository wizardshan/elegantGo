package request

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"time"
)

type TimeRangeFieldV9 struct {
	TimeRangeV9
}

func (req *TimeRangeFieldV9) UnmarshalJSON(b []byte) error {
	layouts := req.layouts()
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s|datetime=%s", layouts[0], layouts[1])),
		req.mapperFunc(),
	)
}

func (req *TimeRangeFieldV9) mapperFunc() func(elems []string) []time.Time {
	return func(elems []string) []time.Time {
		return pie.Map(elems, func(s string) (t time.Time) {
			for _, layout := range req.layouts() {
				t, _ = time.Parse(layout, s)
				if !t.IsZero() {
					return t
				}
			}
			return
		})
	}
}

func (req *TimeRangeFieldV9) layouts() []string {
	return []string{time.DateTime, time.DateOnly}
}
