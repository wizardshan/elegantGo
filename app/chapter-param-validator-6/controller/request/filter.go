package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobeam/stringy"
	"strconv"
	"strings"
	"time"
)

type QueryField struct {
	Sort     *SortField `binding:"omitempty,alphanum"`
	Order    *Order     `binding:"omitempty,oneof=DESC ASC"`
	Page     *Page      `binding:"omitempty,number,min=1"`
	PageSize *PageSize  `binding:"omitempty,number,min=10,max=100"`
}

type SortField string

func (req *SortField) Value() string {
	if req == nil {
		return "id"
	}

	return stringy.New(string(*req)).SnakeCase().ToLower()
}

type Order string

func (req *Order) IsDesc() bool {
	return req == nil || *req == "DESC"
}

type Page int

func (req *Page) Value() int {
	if req == nil {
		return 1
	}

	return int(*req)
}

type PageSize int

func (req *PageSize) Value() int {
	if req == nil {
		return 100
	}

	return int(*req)
}

type NumbersBySeparatorField struct {
	Values []int
}

func (req *NumbersBySeparatorField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	numbers := strings.Split(data, ",")
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		req.Values = append(req.Values, num)
	}
	return nil
}

func (req *NumbersBySeparatorField) Able() bool {
	return req != nil
}

type StringsBySeparatorField struct {
	Values []string
}

func (req *StringsBySeparatorField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	req.Values = strings.Split(data, ",")
	if err := validate.Var(req.Values, "dive,required"); err != nil {
		return err
	}

	return nil
}

func (req *StringsBySeparatorField) Able() bool {
	return req != nil
}

type NumberRangeField struct {
	Start     int
	End       int
	startAble bool
	endAble   bool
}

func (req *NumberRangeField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	numbers := strings.Split(data, ",")
	capacity := len(numbers)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	for i, numStr := range numbers {
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}

		if i == 0 {
			req.startAble = true
			req.Start = num
		} else {
			req.endAble = true
			req.End = num
		}
	}

	if req.startAble && req.endAble && req.Start >= req.End {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *NumberRangeField) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *NumberRangeField) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type DateTimeRangeField struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
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

		t, err := time.Parse(time.DateTime, timeStr)
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

	if req.startAble && req.endAble && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeField) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *DateTimeRangeField) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type DateRangeField struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *DateRangeField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
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

		t, err := time.Parse(time.DateOnly, timeStr)
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

	if req.startAble && req.endAble && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateRangeField) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *DateRangeField) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}

type TimeRangeField struct {
	Start     time.Time
	End       time.Time
	startAble bool
	endAble   bool
}

func (req *TimeRangeField) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
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
		t, errParseDateTime = time.Parse(time.DateTime, timeStr)
		if errParseDateTime != nil {
			var errDateOnly error
			t, errDateOnly = time.Parse(time.DateOnly, timeStr)
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

	if req.startAble && req.endAble && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeField) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *TimeRangeField) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}
