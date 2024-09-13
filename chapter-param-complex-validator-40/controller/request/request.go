package request

import (
	"elegantGo/chapter-param-complex-validator-4/repository/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gobeam/stringy"
)

type LimitField int

func (req *Limit) Value() int {
	if req == nil {
		return req.defaultValue()
	}
	return int(*req)
}

func (req *LimitField) defaultValue() int {
	return 100
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

type OrderField string

func (req *OrderField) By(fields ...string) func(*sql.Selector) {
	if req == nil || string(*req) == req.defaultValue() {
		return ent.Desc(fields...)
	}
	return ent.Asc(fields...)
}

func (req *OrderField) defaultValue() string {
	return "DESC"
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
