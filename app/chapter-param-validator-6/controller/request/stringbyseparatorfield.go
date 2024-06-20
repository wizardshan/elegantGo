package request

type StringBySeparatorField struct {
	ElementBySeparatorField[string]
}

func (req *StringBySeparatorField) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, "dive,required,alphanum", req.parse)
}

func (req *StringBySeparatorField) parse(elements []string) (values []string) {
	return elements
}
