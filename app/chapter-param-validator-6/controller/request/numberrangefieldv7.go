package request

import (
	"strconv"
)

type NumberRangeFieldV7 struct {
	start *int
	end   *int
	RangeParserV7[*int]
}

func (req *NumberRangeFieldV7) UnmarshalJSON(b []byte) error {
	var err error
	req.start, req.end, err = req.Parse(b, req.parseElement, req.rangeValidate)
	return err
}

func (req *NumberRangeFieldV7) rangeValidate(start *int, end *int) bool {
	return start != nil && end != nil && *start >= *end
}

func (req *NumberRangeFieldV7) parseElement(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}
	number, err := strconv.Atoi(value)
	return &number, err
}

func (req *NumberRangeFieldV7) StartAble() bool {
	return req != nil && req.start != nil
}

func (req *NumberRangeFieldV7) EndAble() bool {
	return req != nil && req.End != nil
}

func (req *NumberRangeFieldV7) Start() int {
	if req.StartAble() {
		return *req.start
	}
	return 0
}

func (req *NumberRangeFieldV7) End() int {
	if req.EndAble() {
		return *req.end
	}
	return 0
}
