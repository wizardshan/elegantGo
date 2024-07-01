package request

import (
	"errors"
	"time"
)

type RangeV10[T time.Time | *int] struct {
	JsonString
	Start T
	End   T
}

func (req *RangeV10[T]) parse(b []byte, validateTag string, parseElements func() []T, validRangeFunc func(start T, end T) bool) error {
	if err := req.unmarshal(b); err != nil {
		return err
	}

	if !req.HasSep() {
		return errors.New("parameter should be separated by commas")
	}

	if err := req.valid(validateTag); err != nil {
		return err
	}

	elements := parseElements()
	req.Start, req.End = elements[0], elements[1]

	if validRangeFunc(req.Start, req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *RangeV10[T]) validateTag(tag string) string {
	tags := "len=2,dive,omitempty"
	if tag != "" {
		tags = tags + "," + tag
	}
	return tags
}
