package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeSplitterV5 struct {
	Start time.Time
	End   time.Time
}

func (req *TimeSplitterV5) unmarshalJSON(b []byte, layout string) error {

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
	if req.Start, err = req.parseElement(elements[0], layout); err != nil {
		return err
	}

	if req.End, err = req.parseElement(elements[1], layout); err != nil {
		return err
	}

	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeSplitterV5) parseElement(value string, layout string) (time.Time, error) {
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

func (req *TimeSplitterV5) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *TimeSplitterV5) EndAble() bool {
	return !req.End.IsZero()
}
