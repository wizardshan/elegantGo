package request

type NumberRangeFieldV10 struct {
	Start *int
	End   *int
	RangeV10[*int]
}

func (req *NumberRangeFieldV10) UnmarshalJSON(b []byte) (err error) {

	return req.parse(
		b,
		req.validateTag(),
		req.PtrInts,
		req.validRange,
	)
}

func (req *NumberRangeFieldV10) validateTag() string {
	return req.RangeV10.validateTag("number")
}

func (req *NumberRangeFieldV10) validRange(start *int, end *int) bool {
	return !(start != nil && end != nil && *start >= *end)
}
