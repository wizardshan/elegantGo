package request

type NumberBySeparatorV2 struct {
	JsonString
	Values []int
}

func (req *NumberBySeparatorV2) UnmarshalJSON(b []byte) error {

	err := req.unmarshal(b)
	if err != nil {
		return err
	}

	req.Values = req.Ints()
	return validate.Var(req.Values, "dive,number")
}
