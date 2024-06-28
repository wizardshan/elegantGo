package request

import (
	"fmt"
	"time"
)

type DateRangeFieldV8 struct {
	TimeRangeV8
}

func (req *DateRangeFieldV8) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(
		b,
		fmt.Sprintf("datetime=%s", req.layout()),
		req.parseElementFunc(),
	)
}

func (req *DateRangeFieldV8) parseElementFunc() func(s string) time.Time {
	return func(s string) time.Time {
		return req.parseElement(s, req.layout())
	}
}

func (req *DateRangeFieldV8) layout() string {
	return time.DateOnly
}
