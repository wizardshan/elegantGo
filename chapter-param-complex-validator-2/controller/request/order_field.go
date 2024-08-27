package request

type OrderFieldV1 string

func (req *OrderFieldV1) IsDesc() bool {
	return req == nil || *req == "DESC"
}

type OrderField string

func (req *OrderField) IsDesc() bool {
	return req == nil || string(*req) == req.defaultValue()
}

func (req *OrderField) defaultValue() string {
	return "DESC"
}
