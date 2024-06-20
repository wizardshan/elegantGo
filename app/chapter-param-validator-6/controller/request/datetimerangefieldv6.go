package request

import (
	"time"
)

type DateTimeRangeFieldV6 struct {
	TimeSplitterV6
}

func (req *DateTimeRangeFieldV6) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateTimeRangeFieldV6) parseElement(value string) (time.Time, error) {
	return req.TimeSplitterV6.parseElement(value, time.DateTime)
}
