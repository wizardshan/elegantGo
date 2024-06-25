package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeRangeFieldV1 struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *TimeRangeFieldV1) UnmarshalJSON(b []byte) error {

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

	startStr := elements[0]
	endStr := elements[0]

	if startStr != "" {
		t, err := req.parse(startStr)
		if err != nil {
			return err
		}
		req.startAble = true
		req.Start = t
	}

	if endStr != "" {
		t, err := req.parse(startStr)
		if err != nil {
			return err
		}
		req.endAble = true
		req.End = t
	}

	if req.startAble && req.endAble && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeFieldV1) parse(value string) (t time.Time, err error) {
	layouts := []string{time.DateTime, time.DateOnly}
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return
		}
	}

	err = errors.New(fmt.Sprintf("the time layout should be one of `%s` and `%s`", layouts))
	return
}

func (req *TimeRangeFieldV1) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *TimeRangeFieldV1) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}
