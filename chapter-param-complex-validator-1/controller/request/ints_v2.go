package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/numeral"
	"errors"
	"strconv"
	"strings"
)

type IntsFieldV2 string

func (req *IntsFieldV2) Values() []int {
	ss := req.split()

	var ssFiltered []string
	for _, s := range ss {
		if numeral.IsInt(s) {
			ssFiltered = append(ssFiltered, s)
		}
	}

	var values []int
	for _, s := range ssFiltered {
		value, _ := strconv.Atoi(s)
		values = append(values, value)
	}
	return values
}

func (req *IntsFieldV2) MustValues() ([]int, error) {
	ss := req.split()
	
	for _, s := range ss {
		if !numeral.IsInt(s) {
			return nil, errors.New(s + " is not an integer")
		}
	}

	var values []int
	for _, s := range ss {
		value, _ := strconv.Atoi(s)
		values = append(values, value)
	}
	return values, nil
}

func (req *IntsFieldV2) ShouldValues() []int {
	ss := req.split()

	var values []int
	for _, s := range ss {
		value, _ := strconv.Atoi(s)
		values = append(values, value)
	}
	return values
}

func (req *IntsFieldV2) split() []string {
	return strings.Split(string(*req), ",")
}
