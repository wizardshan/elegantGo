package request

import "time"

type DateRangeFieldV5 struct {
	TimeSplitterV5
}

func (req *DateRangeFieldV5) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, time.DateOnly)
}
