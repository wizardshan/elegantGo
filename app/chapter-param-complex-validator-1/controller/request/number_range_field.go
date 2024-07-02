package request

import (
	"app/chapter-param-complex-validator-1/controller/pkg/mapper"
	"github.com/elliotchance/pie/v2"
	"strconv"
)

type NumberRangeField struct {
	Range[*int]
}

func (req *NumberRangeField) UnmarshalJSON(b []byte) error {
	return req.Parse(
		b,
		req.validateTag("number"),
		mapper.PtrInts,
		req.valid,
	)
}

func (req *NumberRangeField) mapperFunc() func([]string) []*int {
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

func (req *NumberRangeField) valid(start *int, end *int) bool {
	return !(start != nil && end != nil && *start >= *end)
}
