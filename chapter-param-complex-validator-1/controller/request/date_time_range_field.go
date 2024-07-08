package request

import (
	"elegantGo/chapter-param-complex-validator-1/controller/pkg/mapper"
	"fmt"
	"time"
)

type DateTimeRangeField struct {
	TimeRange
}

func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s", req.layout())),
		mapper.DateTimes,
	)
}

func (req *DateTimeRangeField) layout() string {
	return time.DateTime
}
