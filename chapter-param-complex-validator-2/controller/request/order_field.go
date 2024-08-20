package request

import (
	"elegantGo/chapter-param-complex-validator-2/repository/ent"
	"entgo.io/ent/dialect/sql"
)

type OrderFieldV1 string

func (req *OrderFieldV1) IsDesc() bool {
	return req == nil || *req == "DESC"
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
