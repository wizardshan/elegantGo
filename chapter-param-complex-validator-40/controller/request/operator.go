package request

type Operator struct {
	Value string
}

func (op *Operator) defaultValue() string {
	return "="
}
