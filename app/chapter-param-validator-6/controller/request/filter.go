package request

import (
	"github.com/gobeam/stringy"
)

var (
	separatedByCommasMsg  = "parameter should be separated by commas"
	rangeFieldCapacityMsg = "the rangeField capacity expected value is 2, the result is %d"
	rangeFieldCompareMsg  = "the rangeField start must lt end"
)

type QueryField struct {
	Sort     *SortField `binding:"omitempty,alphanum"`
	Order    *Order     `binding:"omitempty,oneof=DESC ASC"`
	Page     *Page      `binding:"omitempty,number,min=1"`
	PageSize *PageSize  `binding:"omitempty,number,min=10,max=100"`
}

type SortField string

func (req *SortField) Value() string {
	if req == nil {
		return "id"
	}

	return stringy.New(string(*req)).SnakeCase().ToLower()
}

type Order string

func (req *Order) IsDesc() bool {
	return req == nil || *req == "DESC"
}

type Page int

func (req *Page) Value() int {
	if req == nil {
		return 1
	}

	return int(*req)
}

type PageSize int

func (req *PageSize) Value() int {
	if req == nil {
		return 100
	}

	return int(*req)
}
