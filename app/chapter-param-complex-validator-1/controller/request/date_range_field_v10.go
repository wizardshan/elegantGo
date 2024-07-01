package request

import (
	"fmt"
	"time"
)

type DateRangeFieldV10 struct {
	TimeRangeV10
}

func (req *DateRangeFieldV10) UnmarshalJSON(b []byte) error {
	return req.parse(
		b,
		req.validateTag(),
		req.Dates,
	)
}

func (req *DateRangeFieldV10) validateTag() string {
	return req.RangeV10.validateTag(fmt.Sprintf("datetime=%s", req.layout()))
}

func (req *DateRangeFieldV10) layout() string {
	return time.DateTime
}
