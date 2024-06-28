package request

import "time"

type DateTimeRangeFieldV7 struct {
	TimeRangeV7
}

func (req *DateTimeRangeFieldV7) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, time.DateTime)
}
