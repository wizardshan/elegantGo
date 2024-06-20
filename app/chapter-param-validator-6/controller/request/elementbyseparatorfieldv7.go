package request

import (
	"encoding/json"
	"errors"
	"strings"
)

type ElementBySeparatorFieldV7[T string | int] struct {
	values []T
}

func (req *ElementBySeparatorFieldV7[T]) unmarshalJSON(b []byte, validateTag string, parse func(elements []string) []T) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	elements := strings.Split(data, ",")
	if err := validate.Var(elements, validateTag); err != nil {
		return err
	}
	req.values = parse(elements)

	return nil
}

func (req *ElementBySeparatorFieldV7[T]) Able() bool {
	return len(req.values) > 0
}

func (req *ElementBySeparatorFieldV7[T]) Values() []T {
	return req.values
}
