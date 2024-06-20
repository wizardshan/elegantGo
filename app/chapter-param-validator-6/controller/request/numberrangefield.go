package request

import (
	"strconv"
)

type NumberRangeField struct {
	start *int
	end   *int
	RangeParserV7[*int]
}

func (req *NumberRangeField) UnmarshalJSON(b []byte) error {
	var err error
	req.start, req.end, err = req.Parse(b, req.parseElement, req.rangeValidate)
	return err
}

func (req *NumberRangeField) rangeValidate(start *int, end *int) bool {
	return start != nil && end != nil && *start >= *end
}

func (req *NumberRangeField) parseElement(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}
	number, err := strconv.Atoi(value)
	return &number, err
}

func (req *NumberRangeField) StartAble() bool {
	return req != nil && req.start != nil
}

func (req *NumberRangeField) EndAble() bool {
	return req != nil && req.End != nil
}

func (req *NumberRangeField) Start() int {
	if req.StartAble() {
		return *req.start
	}
	return 0
}

func (req *NumberRangeField) End() int {
	if req.EndAble() {
		return *req.end
	}
	return 0
}
