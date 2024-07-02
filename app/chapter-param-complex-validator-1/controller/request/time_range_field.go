package request

import (
	"app/chapter-param-complex-validator-1/controller/pkg/mapper"
	"fmt"
	"time"
)

type TimeRangeField struct {
	TimeRange
}

func (req *TimeRangeField) UnmarshalJSON(b []byte) error {
	layouts := req.layouts()
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s|datetime=%s", layouts[0], layouts[1])),
		mapper.Times,
	)
}

func (req *TimeRangeField) layouts() []string {
	return []string{time.DateTime, time.DateOnly}
}
