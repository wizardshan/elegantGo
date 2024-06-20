package request

import (
	"strconv"
)

type NumberBySeparatorFieldV7 struct {
	ElementBySeparatorFieldV7[int]
}

func (req *NumberBySeparatorFieldV7) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, "dive,number", req.parse)
}

func (req *NumberBySeparatorFieldV7) parse(elements []string) (values []int) {
	for _, element := range elements {
		num, _ := strconv.Atoi(element)
		values = append(values, num)
	}
	return
}
