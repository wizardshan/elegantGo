package request

import "time"

type DateRangeFieldV7 struct {
	TimeRangeV7
}

func (req *DateRangeFieldV7) UnmarshalJSON(b []byte) error {
	return req.Parse(b, time.DateOnly)
}
