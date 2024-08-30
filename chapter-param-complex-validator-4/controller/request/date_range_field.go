package request

import (
	"elegantGo/chapter-param-complex-validator-2/controller/pkg/mapper"
	"fmt"
	"time"
)

type DateRangeField struct {
	TimeRange
}

func (req *DateRangeField) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag(fmt.Sprintf("datetime=%s", req.layout())),
		mapper.Dates,
	)
}

func (req *DateRangeField) layout() string {
	return time.DateOnly
}
