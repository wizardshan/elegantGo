package request

import (
	"errors"
	"fmt"
	"time"
)

type TimeSplitter struct {
	Start time.Time
	End   time.Time
	RangeParserV7[time.Time]
}

func (req *TimeSplitter) unmarshalJSON(b []byte, parseElement func(value string) (time.Time, error)) (err error) {
	req.Start, req.End, err = req.Parse(b, parseElement, req.rangeValidate)
	return
}

func (req *TimeSplitter) rangeValidate(start time.Time, end time.Time) bool {
	return !start.IsZero() && !end.IsZero() && start.After(end)
}

func (req *TimeSplitter) parseElement(value string, layout string) (t time.Time, err error) {
	if value == "" {
		return
	}

	t, err = time.Parse(layout, value)
	if err != nil {
		err = errors.New(fmt.Sprintf("the time layout should be `%s`", layout))
	}

	return
}

func (req *TimeSplitter) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *TimeSplitter) EndAble() bool {
	return !req.End.IsZero()
}
