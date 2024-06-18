package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldRaw struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *DateTimeRangeFieldRaw) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if data == "" {
		req.startAble = false
		req.endAble = false
		return nil
	}

	times := strings.Split(data, ",")
	capacity := len(times)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	for i, timeStr := range times {
		if timeStr == "" {
			continue
		}

		t, err := time.Parse(time.DateTime, data)
		if err != nil {
			return err
		}

		if i == 0 {
			req.startAble = true
			req.Start = t
		} else {
			req.endAble = true
			req.End = t
		}
	}

	if req.startAble && req.endAble && req.Start.Before(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldRaw) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *DateTimeRangeFieldRaw) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type DateRangeFieldRaw struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *DateRangeFieldRaw) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if data == "" {
		req.startAble = false
		req.endAble = false
		return nil
	}

	times := strings.Split(data, ",")
	capacity := len(times)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	for i, timeStr := range times {
		if timeStr == "" {
			continue
		}

		t, err := time.Parse(time.DateOnly, data)
		if err != nil {
			return err
		}

		if i == 0 {
			req.startAble = true
			req.Start = t
		} else {
			req.endAble = true
			req.End = t
		}
	}

	if req.startAble && req.endAble && req.Start.Before(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateRangeFieldRaw) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *DateRangeFieldRaw) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type TimeRangeFieldRaw struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *TimeRangeFieldRaw) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if data == "" {
		req.startAble = false
		req.endAble = false
		return nil
	}

	times := strings.Split(data, ",")
	capacity := len(times)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	for i, timeStr := range times {
		if timeStr == "" {
			continue
		}

		var t time.Time
		var errParseDateTime error
		t, errParseDateTime = time.Parse(time.DateTime, data)
		if errParseDateTime != nil {
			var errDateOnly error
			t, errDateOnly = time.Parse(time.DateOnly, data)
			if errDateOnly != nil {
				return errors.New(fmt.Sprintf("the time layout should be one of %s and %s", time.DateTime, time.DateOnly))
			}
		}

		if i == 0 {
			req.startAble = true
			req.Start = t
		} else {
			req.endAble = true
			req.End = t
		}
	}

	if req.startAble && req.endAble && req.Start.Before(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeFieldRaw) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *TimeRangeFieldRaw) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}
