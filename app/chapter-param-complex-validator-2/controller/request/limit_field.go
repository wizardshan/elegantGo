package request

type LimitFieldV1 int

func (req *LimitFieldV1) Value() int {
	if req == nil {
		return 100
	}

	return int(*req)
}

type LimitField int

func (req *LimitField) Value() int {
	if req == nil {
		return req.defaultValue()
	}

	return int(*req)
}

func (req *LimitField) defaultValue() int {
	return 100
}
