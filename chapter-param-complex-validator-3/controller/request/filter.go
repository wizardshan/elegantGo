package request

import "time"

type QueryField struct {
	Sort   *SortField   `binding:"omitempty,alphanum"`
	Order  *OrderField  `binding:"omitempty,oneof=DESC ASC"`
	Offset *OffsetField `binding:"omitempty,number,min=0"`
	Limit  *LimitField  `binding:"omitempty,number,min=1,max=1000"`
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
