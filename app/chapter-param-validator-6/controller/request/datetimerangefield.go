package request

import (
	"time"
)

type DateTimeRangeField struct {
	TimeSplitter
}

func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateTimeRangeField) parseElement(value string) (time.Time, error) {
	return req.TimeSplitter.parseElement(value, time.DateTime)
}
