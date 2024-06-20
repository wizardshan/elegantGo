package request

import "time"

type DateTimeRangeFieldV5 struct {
	TimeSplitterV5
}

func (req *DateTimeRangeFieldV5) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, time.DateTime)
}
