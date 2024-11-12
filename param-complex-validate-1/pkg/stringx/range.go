package stringx

import (
	"errors"
	"time"
)

var rangeInvalidError = errors.New("the range start must lt end")

type Range[T any] struct {
	Start T
	End   T
}

func (r Range[T]) capValid(ss []T) error {
	if len(ss) != 2 {
		return errors.New("the range expected is two capacity")
	}
	return nil
}

type IntRange struct {
	Range[*int]
}

func (i *IntRange) Parse(input string) error {
	splitter, err := NewSplitter(input, Validator(IsInt))
	if err != nil {
		return err
	}

	ints := splitter.PtrInts()
	if err = i.capValid(ints); err != nil {
		return err
	}

	i.Start, i.End = ints[0], ints[1]
	if i.Start != nil && i.End != nil && *i.Start > *i.End {
		return rangeInvalidError
	}
	return nil
}

type timeRange struct {
	Range[time.Time]
}

func (t *timeRange) Parse(input string, validator func(s string) bool, fn func(spr *Splitter) []time.Time) error {
	splitter, err := NewSplitter(input, Validator(validator))
	if err != nil {
		return err
	}
	times := fn(splitter)
	if err = t.capValid(times); err != nil {
		return err
	}

	t.Start, t.End = times[0], times[1]
	if !t.Start.IsZero() && !t.End.IsZero() && t.Start.After(t.End) {
		return rangeInvalidError
	}
	return nil
}

type DateTimeRange struct {
	timeRange
}

func (t *DateTimeRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsDateTime, func(spr *Splitter) []time.Time {
		return spr.DateTimes()
	})
}

type DateRange struct {
	timeRange
}

func (t *DateRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsDate, func(spr *Splitter) []time.Time {
		return spr.Dates()
	})
}

type TimeRange struct {
	timeRange
}

func (t *TimeRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsTime, func(spr *Splitter) []time.Time {
		return spr.Times()
	})
}
