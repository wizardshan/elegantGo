package request

import "time"

type DateRangeField struct {
	TimeSplitter
}

func (req *DateRangeField) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateRangeField) parseElement(value string) (time.Time, error) {
	return req.TimeSplitter.parseElement(value, time.DateOnly)
}
