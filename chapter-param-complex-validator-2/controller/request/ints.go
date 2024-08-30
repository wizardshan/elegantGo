package request

import (
	"elegantGo/chapter-param-complex-validator-2/pkg/stringx"
	"encoding/json"
)

type IntsField string

func (req *IntsField) Values() ([]int, error) {
	return stringx.NewSplitter(string(*req)).Validator(stringx.IsInt).Parse().Ints()
}

type IntsJsonField struct {
	Values []int
}

func (req *IntsJsonField) UnmarshalJSON(b []byte) (err error) {
	var data IntsField
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	req.Values, err = data.Values()
	return
}

type IntRangeField string

func (req *IntRangeField) Values() (start *int, end *int, err error) {
	rg := new(stringx.IntRange)
	if err = rg.Parse(string(*req)); err != nil {
		return
	}
	start, end = rg.Start, rg.End
	return
}

type IntRangeJsonField struct {
	stringx.IntRange
}

func (req *IntRangeJsonField) UnmarshalJSON(b []byte) error {
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	return req.Parse(data)
}
