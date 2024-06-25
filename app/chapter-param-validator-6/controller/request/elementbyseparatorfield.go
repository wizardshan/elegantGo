package request

type ElementBySeparatorField[T string | int] struct {
	values []T
	CommasSplitter
}

func (req *ElementBySeparatorField[T]) unmarshalJSON(b []byte, validateTag string, parse func(elements []string) []T) error {

	elements, err := req.Parse(b)
	if err != nil {
		return err
	}

	if err = validate.Var(elements, validateTag); err != nil {
		return err
	}
	req.values = parse(elements)

	return nil
}

func (req *ElementBySeparatorField[T]) Able() bool {
	return len(req.values) > 0
}

func (req *ElementBySeparatorField[T]) Values() []T {
	return req.values
}
