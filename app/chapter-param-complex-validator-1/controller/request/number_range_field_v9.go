package request

import (
	"strconv"
)

type NumberRangeFieldV9 struct {
	Start *int
	End   *int
	RangeV9[*int]
}

func (req *NumberRangeFieldV9) UnmarshalJSON(b []byte) (err error) {
	req.Start, req.End, err = req.parse(b, "number", req.parseElementFunc(), req.valid)
	return
}

func (req *NumberRangeFieldV9) parseElementFunc() func(s string) *int {
	return func(s string) *int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &num
	}
}

func (req *NumberRangeFieldV9) valid(start *int, end *int) bool {
	return !(start != nil && end != nil && *start >= *end)
}
