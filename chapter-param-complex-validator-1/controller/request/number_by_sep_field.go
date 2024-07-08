package request

import (
	"elegantGo/chapter-param-complex-validator-1/controller/pkg/mapper"
	"fmt"
)

type NumberBySepField struct {
	JsonStringBySep[int]
	Values []int
}

func (req *NumberBySepField) UnmarshalJSON(b []byte) error {
	if err := req.unmarshal(b); err != nil {
		return err
	}

	if err := validate.Var(req.elements, "dive,number"); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	req.Values = req.mapper(mapper.Ints)
	return nil
}
