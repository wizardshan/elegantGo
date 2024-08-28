package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/numeral"
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
)

type IntsField string

func (req *IntsField) Values() []int {
	return pie.Ints(pie.Filter(req.split(), numeral.IsInt))
}

func (req *IntsField) MustValues() ([]int, error) {
	ss := req.split()
	if !pie.All(ss, numeral.IsInt) {
		return nil, errors.New("one of numbers is not an integer")
	}

	return pie.Ints(ss), nil
}

func (req *IntsField) ShouldValues() []int {
	return pie.Ints(req.split())
}

func (req *IntsField) split() []string {
	return strings.Split(string(*req), ",")
}
