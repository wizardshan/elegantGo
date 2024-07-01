package request

import (
	"fmt"
	"time"
)

type DateTimeRangeFieldV10 struct {
	TimeRangeV10
}

func (req *DateTimeRangeFieldV10) UnmarshalJSON(b []byte) error {
	return req.parse(
		b,
		req.validateTag(),
		req.DateTimes,
	)
}

func (req *DateTimeRangeFieldV10) validateTag() string {
	return req.RangeV10.validateTag(fmt.Sprintf("datetime=%s", req.layout()))
}

func (req *DateTimeRangeFieldV10) layout() string {
	return time.DateTime
}
