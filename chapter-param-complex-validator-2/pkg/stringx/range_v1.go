package stringx

import (
	"time"
)

type DateTimeRangeV1 struct {
	Range[time.Time]
}

func (t *DateTimeRangeV1) Parse(input string) error {
	times, err := NewSplitter(input).Validator(IsDateTime).Parse().DateTimes()
	if err != nil {
		return err
	}

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
	times, err := NewSplitter(input).Validator(IsDate).Parse().Dates()
	if err != nil {
		return err
	}

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
	times, err := NewSplitter(input).Validator(IsTime).Parse().Times()
	if err != nil {
		return err
	}

	if err = t.capValid(times); err != nil {
		return err
	}

	t.Start, t.End = times[0], times[1]
	if !t.Start.IsZero() && !t.End.IsZero() && t.Start.After(t.End) {
		return t.invalidError()
	}
	return nil
}
