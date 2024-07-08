package request

type OffsetFieldV1 int

func (req *OffsetFieldV1) Value() int {
	if req == nil {
		return 0
	}

	return int(*req)
}

type OffsetField int

func (req *OffsetField) Value() int {
	if req == nil {
		return req.defaultValue()
	}

	return int(*req)
}

func (req *OffsetField) defaultValue() int {
	return 0
}
