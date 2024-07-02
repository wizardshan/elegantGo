package request

type StringBySepFieldV2 struct {
	JsonStringBySep[string]
	Values []string
}

func (req *StringBySepFieldV2) UnmarshalJSON(b []byte) error {
	if err := req.unmarshal(b); err != nil {
		return err
	}

	if err := validate.Var(req.Values, "dive,required,alphanum"); err != nil {
		return err
	}

	req.Values = req.elements
	return nil
}
