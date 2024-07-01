package request

type StringBySeparatorV2 struct {
	JsonString
	Values []string
}

func (req *StringBySeparatorV2) UnmarshalJSON(b []byte) error {

	err := req.unmarshal(b)
	if err != nil {
		return err
	}

	req.Values = req.Strings()
	return validate.Var(req.Values, "dive,required,alphanum")
}
