package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeSplitterV6 struct {
	Start time.Time
	End   time.Time
}

func (req *TimeSplitterV6) unmarshalJSON(b []byte, parseElement func(value string) (time.Time, error)) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	elements := strings.Split(data, ",")
	capacity := len(elements)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	var err error
	if req.Start, err = parseElement(elements[0]); err != nil {
		return err
	}

	if req.End, err = parseElement(elements[1]); err != nil {
		return err
	}

	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeSplitterV6) parseElement(value string, layout string) (time.Time, error) {
	var t time.Time
	var err error
	if value == "" {
		return t, nil
	}

	t, err = time.Parse(layout, value)
	if err != nil {
		return t, errors.New(fmt.Sprintf("the time layout should be `%s`", layout))
	}

	return t, nil
}

func (req *TimeSplitterV6) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *TimeSplitterV6) EndAble() bool {
	return !req.End.IsZero()
}
