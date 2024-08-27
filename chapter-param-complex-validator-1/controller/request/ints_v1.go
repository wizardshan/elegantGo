package request

import (
	"strconv"
	"strings"
)

type IntsFieldV1 string

func (req *IntsFieldV1) Values() []int {
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

func (req *IntsFieldV1) MustValues() ([]int, error) {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func (req *IntsFieldV1) ShouldValues() []int {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, _ := strconv.Atoi(s)
		values = append(values, value)
	}
	return values
}
