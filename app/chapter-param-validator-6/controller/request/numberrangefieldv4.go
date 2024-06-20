package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type NumberRangeFieldV4 struct {
	start *int
	end   *int
}

func (req *NumberRangeFieldV4) UnmarshalJSON(b []byte) error {

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
	if req.start, err = req.parseElement(elements[0]); err != nil {
		return err
	}

	if req.end, err = req.parseElement(elements[1]); err != nil {
		return err
	}

	if req.start != nil && req.end != nil && *req.start >= *req.end {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *NumberRangeFieldV4) parseElement(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}
	number, err := strconv.Atoi(value)
	return &number, err
}

func (req *NumberRangeFieldV4) StartAble() bool {
	return req != nil && req.start != nil
}

func (req *NumberRangeFieldV4) EndAble() bool {
	return req != nil && req.End != nil
}

func (req *NumberRangeFieldV4) Start() int {
	if req.StartAble() {
		return *req.start
	}
	return 0
}

func (req *NumberRangeFieldV4) End() int {
	if req.EndAble() {
		return *req.end
	}
	return 0
}
