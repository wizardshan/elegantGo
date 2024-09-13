package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/stringx"
	"errors"
	"strconv"
	"strings"
)

type IntsFieldV3 string

func (req *IntsFieldV3) Values() []int {
	ss := req.split()

	var ssFiltered []string
	for _, s := range ss {
		if stringx.IsInt(s) {
			ssFiltered = append(ssFiltered, s)
		}
	}

	var values []int
	for _, s := range ssFiltered {
		values = append(values, req.toInt(s))
	}
	return values
}

func (req *IntsFieldV3) MustValues() ([]int, error) {
	ss := req.split()

	for _, s := range ss {
		if !stringx.IsInt(s) {
			return nil, errors.New(s + " is not an integer")
		}
	}

	var values []int
	for _, s := range ss {
		values = append(values, req.toInt(s))
	}
	return values, nil
}

func (req *IntsFieldV3) ShouldValues() []int {
	ss := req.split()

	var values []int
	for _, s := range ss {
		values = append(values, req.toInt(s))
	}
	return values
}

func (req *IntsFieldV3) split() []string {
	return strings.Split(string(*req), ",")
}

func (req *IntsFieldV3) toInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
