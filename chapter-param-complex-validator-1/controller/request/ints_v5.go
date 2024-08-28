package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/numeral"
	"errors"
	"github.com/elliotchance/pie/v2"
	"strconv"
	"strings"
)

type IntsFieldV5 string

func (req *IntsFieldV5) Values() []int {
	ss := req.split()

	var ssFiltered []string
	for _, s := range ss {
		if numeral.IsInt(s) {
			ssFiltered = append(ssFiltered, s)
		}
	}

	return pie.Map(ssFiltered, req.toInt)
}

func (req *IntsFieldV5) MustValues() ([]int, error) {
	ss := req.split()

	for _, s := range ss {
		if !numeral.IsInt(s) {
			return nil, errors.New(s + " is not an integer")
		}
	}

	return pie.Map(ss, req.toInt), nil
}

func (req *IntsFieldV5) ShouldValues() []int {
	ss := req.split()

	return pie.Map(ss, req.toInt)
}

func (req *IntsFieldV5) split() []string {
	return strings.Split(string(*req), ",")
}

func (req *IntsFieldV5) toInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
