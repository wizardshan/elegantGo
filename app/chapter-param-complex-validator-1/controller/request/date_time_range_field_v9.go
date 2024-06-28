package request

import (
	"fmt"
	"time"
)

type DateTimeRangeFieldV9 struct {
	TimeRangeV9
}

func (req *DateTimeRangeFieldV9) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(
		b,
		fmt.Sprintf("datetime=%s", req.layout()),
		req.parseElementFunc(),
	)
}

func (req *DateTimeRangeFieldV9) parseElementFunc() func(s string) time.Time {
	return func(s string) time.Time {
		return req.parseElement(s, req.layout())
	}
}

func (req *DateTimeRangeFieldV9) layout() string {
	return time.DateTime
}
