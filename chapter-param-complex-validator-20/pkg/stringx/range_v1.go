package stringx

import (
	"time"
)

type DateTimeRangeV1 struct {
	Range[time.Time]
}

func (t *DateTimeRangeV1) Parse(input string) error {
	splitter, err := NewSplitter(input, Validator(IsDateTime))
	if err != nil {
		return err
	}

	times := splitter.DateTimes()
	if err = t.capValid(times); err != nil {
		return err
	}

	t.Start, t.End = times[0], times[1]
	if !t.Start.IsZero() && !t.End.IsZero() && t.Start.After(t.End) {
		return t.invalidError()
	}
	return nil
}

type DateRangeV1 struct {
	Range[time.Time]
}

func (t *DateRangeV1) Parse(input string) error {
	splitter, err := NewSplitter(input, Validator(IsDate))
	if err != nil {
		return err
	}

	times := splitter.Dates()
	if err = t.capValid(times); err != nil {
		return err
	}

	t.Start, t.End = times[0], times[1]
	if !t.Start.IsZero() && !t.End.IsZero() && t.Start.After(t.End) {
		return t.invalidError()
	}
	return nil
}

type TimeRangeV1 struct {
	Range[time.Time]
}

func (t *TimeRangeV1) Parse(input string) error {
	splitter, err := NewSplitter(input, Validator(IsTime))
	if err != nil {
		return err
	}

	times := splitter.Times()
	if err = t.capValid(times); err != nil {
		return err
	}

	t.Start, t.End = times[0], times[1]
	if !t.Start.IsZero() && !t.End.IsZero() && t.Start.After(t.End) {
		return t.invalidError()
	}
	return nil
}
