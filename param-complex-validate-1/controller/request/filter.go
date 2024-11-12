package request

import (
	"github.com/samber/lo"
)

type SortField struct {
	Sort string `binding:"alphanum"`
}

func (req SortField) Value() string {
	if req.Sort == "" {
		return "ID"
	}
	return lo.SnakeCase(req.Sort) // 配合entGo框架，传入CreateTime转化为create_time，ID转化为id
}

type OrderField struct {
	Order string `binding:"oneof=DESC ASC"`
}

func (req *OrderField) Value() string {
	if req.Order == "" {
		return "DESC"
	}
	return req.Order
}

type OffsetField struct {
	Offset int `binding:"number,min=0"`
}

func (req OffsetField) Value() int {
	return req.Offset
}

type LimitField struct {
	Limit int `binding:"number,min=1,max=100"`
}

func (req *LimitField) Value() int {
	if req.Limit == 0 {
		return 100
	}
	return req.Limit
}
