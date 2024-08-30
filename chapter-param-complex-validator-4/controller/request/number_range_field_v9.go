package request

import (
	"github.com/elliotchance/pie/v2"
	"strconv"
)

type NumberRangeFieldV9 struct {
	RangeV9[*int]
}

func (req *NumberRangeFieldV9) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag("number"),
		req.mapperFunc(),
		req.valid,
	)
}

func (req *NumberRangeFieldV9) mapperFunc() func([]string) []*int {
	return func(elems []string) []*int {
		return pie.Map(elems, func(s string) *int {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil
			}
			return &num
		})
	}
}

func (req *NumberRangeFieldV9) valid(start *int, end *int) bool {
	return !(start != nil && end != nil && *start >= *end)
}
