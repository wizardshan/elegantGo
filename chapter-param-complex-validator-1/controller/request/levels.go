package request

import (
	"strconv"
	"strings"
)

type LevelsField string

func (req *LevelsField) Values() []int {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	return values
}
