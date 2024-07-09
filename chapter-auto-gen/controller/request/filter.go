package request

import (
	"elegantGo/chapter-auto-gen/repository/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/gobeam/stringy"
	"time"
)

type QueryField struct {
	Sort   *SortField   `binding:"omitempty,alphanum"`
	Order  *OrderField  `binding:"omitempty,oneof=DESC ASC"`
	Offset *OffsetField `binding:"omitempty,number,min=0"`
	Limit  *LimitField  `binding:"omitempty,number,min=1,max=1000"`
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

type TimeRange struct {
	Start time.Time
	End   time.Time `binding:"omitempty,gtfield=Start"`
}

func (req TimeRange) StartAble() bool {
	return !req.Start.IsZero()
}

func (req TimeRange) EndAble() bool {
	return !req.End.IsZero()
}

type NumberRange struct {
	Start *int
	End   *int `binding:"omitempty,nilfield=Start|gtfield=Start"`
}

func (req NumberRange) StartAble() bool {
	return req.Start != nil
}

func (req NumberRange) EndAble() bool {
	return req.End != nil
}

type Numbers []int

func (req Numbers) Able() bool {
	return req != nil && len(req) != 0
}

type Strings []string

func (req Strings) Able() bool {
	return req != nil && len(req) != 0
}
