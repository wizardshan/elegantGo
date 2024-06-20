package request

import "time"

type DateRangeFieldV7 struct {
	TimeSplitterV7
}

func (req *DateRangeFieldV7) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateRangeFieldV7) parseElement(value string) (time.Time, error) {
	return req.TimeSplitterV7.parseElement(value, time.DateOnly)
}
