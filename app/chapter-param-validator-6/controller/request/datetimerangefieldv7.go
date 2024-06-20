package request

import (
	"time"
)

type DateTimeRangeFieldV7 struct {
	TimeSplitterV7
}

func (req *DateTimeRangeFieldV7) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateTimeRangeFieldV7) parseElement(value string) (time.Time, error) {
	return req.TimeSplitterV7.parseElement(value, time.DateTime)
}
