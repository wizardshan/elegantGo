package request

import (
	"fmt"
	"time"
)

type TimeRangeFieldV10 struct {
	TimeRangeV10
}

func (req *TimeRangeFieldV10) UnmarshalJSON(b []byte) error {
	return req.parse(
		b,
		fmt.Sprintf("datetime=%s|datetime=%s", req.layouts()[0], req.layouts()[1]),
		req.Times,
	)
}

func (req *TimeRangeFieldV10) layouts() []string {
	return []string{time.DateTime, time.DateOnly}
}
