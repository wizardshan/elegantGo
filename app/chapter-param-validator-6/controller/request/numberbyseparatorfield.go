package request

import (
	"strconv"
)

type NumberBySeparatorField struct {
	ElementBySeparatorField[int]
}

func (req *NumberBySeparatorField) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, "dive,number", req.parse)
}

func (req *NumberBySeparatorField) parse(elements []string) (values []int) {
	for _, element := range elements {
		num, _ := strconv.Atoi(element)
		values = append(values, num)
	}
	return
}
