package request

import (
	"fmt"
	"time"
)

type TimeRangeFieldV8 struct {
	TimeRangeV8
}

func (req *TimeRangeFieldV8) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(
		b,
		fmt.Sprintf("datetime=%s|datetime=%s", req.layouts()[0], req.layouts()[1]),
		req.parseElementFunc(),
	)
}

func (req *TimeRangeFieldV8) parseElementFunc() func(s string) time.Time {
	return func(s string) (t time.Time) {
		for _, layout := range req.layouts() {
			t = req.parseElement(s, layout)
			if !t.IsZero() {
				break
			}
		}
		return
	}
}

func (req *TimeRangeFieldV8) layouts() []string {
	return []string{time.DateTime, time.DateOnly}
}
