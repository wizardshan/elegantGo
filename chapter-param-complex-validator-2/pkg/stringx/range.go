package stringx

import (
	"errors"
	"time"
)

type Range[T any] struct {
	Start T
	End   T
}

func (r *Range[T]) capValid(ss []T) error {
	if len(ss) != r.capacity() {
		return errors.New("the range expected is two capacity")
	}
	return nil
}

func (r *Range[T]) invalidError() error {
	return errors.New("the range start must lt end")
}

func (r *Range[T]) capacity() int {
	return 2
}

type IntRange struct {
	Range[*int]
}

func (i *IntRange) Parse(input string) error {
	ints, err := NewSplitter(input).Validator(IsInt).Parse().PtrInts()
	if err != nil {
		return err
	}

	if err = i.capValid(ints); err != nil {
		return err
	}

	i.Start, i.End = ints[0], ints[1]
	if i.Start != nil && i.End != nil && *i.Start > *i.End {
		return i.invalidError()
	}
	return nil
}

type timeRange struct {
	Range[time.Time]
}

func (t *timeRange) Parse(input string, validator func(s string) bool, getter func(pieces *Splitter) ([]time.Time, error)) error {
	parser := NewSplitter(input).Validator(validator).Parse()
	times, err := getter(parser)
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

type DateTimeRange struct {
	timeRange
}

func (t *DateTimeRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsDateTime, func(pieces *Splitter) ([]time.Time, error) {
		return pieces.DateTimes()
	})
}

type DateRange struct {
	timeRange
}

func (t *DateRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsDate, func(pieces *Splitter) ([]time.Time, error) {
		return pieces.Dates()
	})
}

type TimeRange struct {
	timeRange
}

func (t *TimeRange) Parse(input string) error {
	return t.timeRange.Parse(input, IsTime, func(pieces *Splitter) ([]time.Time, error) {
		return pieces.Times()
	})
}
