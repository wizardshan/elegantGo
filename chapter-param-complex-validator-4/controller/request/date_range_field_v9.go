package request

import (
	"fmt"
	"time"
)

type DateRangeFieldV9 struct {
	TimeRangeV9
}

func (req *DateRangeFieldV9) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s", req.layout())),
		req.mapperFunc(),
	)
}

func (req *DateRangeFieldV9) mapperFunc() func(elems []string) []time.Time {
	return func(elems []string) []time.Time {
		return req.mapper(elems, req.layout())
	}
}

func (req *DateRangeFieldV9) layout() string {
	return time.DateOnly
}
