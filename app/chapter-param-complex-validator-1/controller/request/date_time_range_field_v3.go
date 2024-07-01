package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV3 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV3) UnmarshalJSON(b []byte) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	elements, err := req.split(data)
	if err != nil {
		return err
	}

	if err = req.parse(elements); err != nil {
		return err
	}

	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV3) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *DateTimeRangeFieldV3) split(data string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	capacity := len(elements)
	if capacity != 2 {
		err = errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}
	return
}

func (req *DateTimeRangeFieldV3) parse(elements []string) (err error) {
	startStr := elements[0]
	endStr := elements[1]

	layout := time.DateTime
	timeParseErr := fmt.Sprintf("the time layout should be `%s`", layout)

	if startStr != "" {
		req.Start, err = time.Parse(layout, startStr)
		if err != nil {
			return errors.New(timeParseErr)
		}
	}

	if endStr != "" {
		req.End, err = time.Parse(layout, endStr)
		if err != nil {
			return errors.New(timeParseErr)
		}
	}

	return
}

func (req *DateTimeRangeFieldV3) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}
