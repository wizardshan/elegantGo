package request

type StringBySeparatorFieldV7 struct {
	ElementBySeparatorFieldV7[string]
}

func (req *StringBySeparatorFieldV7) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, "dive,required,alphanum", req.parse)
}

func (req *StringBySeparatorFieldV7) parse(elements []string) (values []string) {
	return elements
}
