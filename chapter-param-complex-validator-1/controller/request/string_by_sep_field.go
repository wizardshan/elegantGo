package request

import "fmt"

type StringBySepField struct {
	JsonStringBySep[string]
	Values []string
}

func (req *StringBySepField) UnmarshalJSON(b []byte) error {
	if err := req.unmarshal(b); err != nil {
		return err
	}

	if err := validate.Var(req.elements, "dive,required,alphanum"); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	req.Values = req.elements
	return nil
}
