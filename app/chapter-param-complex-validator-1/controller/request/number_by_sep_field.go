package request

import "app/chapter-param-complex-validator-1/controller/pkg/mapper"

type NumberBySepField struct {
	JsonStringBySep[int]
	Values []int
}

func (req *NumberBySepField) UnmarshalJSON(b []byte) error {
	if err := req.unmarshal(b); err != nil {
		return err
	}

	if err := validate.Var(req.elements, "dive,number"); err != nil {
		return err
	}

	req.Values = req.mapper(mapper.Ints)
	return nil
}
