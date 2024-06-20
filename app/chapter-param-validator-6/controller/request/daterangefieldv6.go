package request

import "time"

type DateRangeFieldV6 struct {
	TimeSplitterV6
}

func (req *DateRangeFieldV6) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *DateRangeFieldV6) parseElement(value string) (time.Time, error) {
	return req.TimeSplitterV6.parseElement(value, time.DateOnly)
}
