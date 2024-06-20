package request

import (
	"errors"
	"fmt"
	"time"
)

type RangeParser[T time.Time | *int] struct {
	CommasSplitter
}

func (req *RangeParser[T]) Parse(b []byte, parseElement func(value string) (T, error), rangeValidate func(start T, end T) bool) (start T, end T, err error) {

	elements, err := req.CommasSplitter.Parse(b)
	if err != nil {
		return
	}

	capacity := len(elements)
	if capacity != 2 {
		err = errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
		return
	}

	if start, err = parseElement(elements[0]); err != nil {
		return
	}

	if end, err = parseElement(elements[1]); err != nil {
		return
	}

	if rangeValidate(start, end) {
		err = errors.New("the rangeField start must lt end")
		return
	}

	return
}
