package request

import "github.com/gobeam/stringy"

type QueryField struct {
	Sort   SortField   `binding:"omitempty,alphanum"`
	Order  OrderField  `binding:"omitempty,oneof=DESC ASC"`
	Offset OffsetField `binding:"omitempty,number,min=0"`
	Limit  LimitField  `binding:"omitempty,number,min=1,max=1000"`
}

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

type SortFieldV1 string

func (req *SortFieldV1) Value() string {
	if req == nil {
		return "id"
	}

	return stringy.New(string(*req)).SnakeCase().ToLower()
}

type SortField string

func (req *SortField) Value() string {
	if req == nil {
		return req.defaultValue()
	}

	return stringy.New(string(*req)).SnakeCase().ToLower()
}

func (req *SortField) defaultValue() string {
	return "id"
}

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
