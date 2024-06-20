package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type RangeParserV7[T time.Time | *int] struct {
	CommasSplitter
}

func (req *RangeParserV7[T]) Parse(b []byte, parseElement func(value string) (T, error), rangeValidate func(start T, end T) bool) (start T, end T, err error) {

	var data string
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	if find := strings.Contains(data, ","); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}

	elements := strings.Split(data, ",")

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
