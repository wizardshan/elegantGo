package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/numeral"
	"errors"
	"strconv"
	"strings"
)

type IntsFieldV4 string

func (req *IntsFieldV4) Values() []int {
	ss := req.split()

	var ssFiltered []string
	for _, s := range ss {
		if numeral.IsInt(s) {
			ssFiltered = append(ssFiltered, s)
		}
	}

	return req.mapper(ssFiltered, req.toInt)
}

func (req *IntsFieldV4) MustValues() ([]int, error) {
	ss := req.split()

	for _, s := range ss {
		if !numeral.IsInt(s) {
			return nil, errors.New(s + " is not an integer")
		}
	}

	return req.mapper(ss, req.toInt), nil
}

func (req *IntsFieldV4) ShouldValues() []int {
	ss := req.split()

	return req.mapper(ss, req.toInt)
}

func (req *IntsFieldV4) split() []string {
	return strings.Split(string(*req), ",")
}

func (req *IntsFieldV4) mapper(ss []string, fn func(s string) int) []int {
	var values []int
	for _, s := range ss {
		values = append(values, fn(s))
	}
	return values
}

func (req *IntsFieldV4) toInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
