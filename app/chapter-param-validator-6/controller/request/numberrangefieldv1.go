package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type NumberRangeFieldV1 struct {
	Start     int
	End       int
	startAble bool
	endAble   bool
}

func (req *NumberRangeFieldV1) UnmarshalJSON(b []byte) error {

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

	var err error
	if startStr != "" {
		req.Start, err = strconv.Atoi(startStr)
		if err != nil {
			return err
		}
		req.startAble = true
	}

	if endStr != "" {
		req.End, err = strconv.Atoi(endStr)
		if err != nil {
			return err
		}
		req.startAble = true
	}

	if req.startAble && req.endAble && req.Start >= req.End {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *NumberRangeFieldV1) StartAble() bool {
	if req == nil {
		return false
	}

	return req.startAble
}

func (req *NumberRangeFieldV1) EndAble() bool {
	if req == nil {
		return false
	}

	return req.endAble
}
