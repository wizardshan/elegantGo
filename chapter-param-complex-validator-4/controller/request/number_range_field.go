package request

import (
	"elegantGo/chapter-param-complex-validator-2/controller/pkg/mapper"
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

func (req *NumberRangeField) valid(start *int, end *int) bool {
	return !(start != nil && end != nil && *start >= *end)
}
