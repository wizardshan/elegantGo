package request

import (
	"fmt"
	"time"
)

type DateTimeRangeFieldV8 struct {
	TimeRangeV8
}

func (req *DateTimeRangeFieldV8) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s", req.layout())),
		req.mapperFunc(),
	)
}

func (req *DateTimeRangeFieldV8) mapperFunc() func(elems []string) []time.Time {
	return func(elems []string) []time.Time {
		return req.mapper(elems, req.layout())
	}
}

func (req *DateTimeRangeFieldV8) layout() string {
	return time.DateTime
}
