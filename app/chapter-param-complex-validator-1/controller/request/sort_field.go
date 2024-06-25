package request

import "github.com/gobeam/stringy"

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
