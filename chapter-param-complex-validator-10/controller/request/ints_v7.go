package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/stringx"
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
)

type IntsFieldV7 string

func (req *IntsFieldV7) Values() []int {
	ss := req.split()
	ssFiltered := pie.Filter(ss, stringx.IsInt)
	return pie.Ints(ssFiltered)
}

func (req *IntsFieldV7) MustValues() ([]int, error) {
	ss := req.split()

	if !pie.All(ss, stringx.IsInt) {
		return nil, errors.New("one of numbers is not an integer")
	}

	//if err := validate.Var(ss, "dive,int"); err != nil {
	//	return nil, err
	//}

	return pie.Ints(ss), nil
}

func (req *IntsFieldV7) ShouldValues() []int {
	ss := req.split()

	return pie.Ints(ss)
}

func (req *IntsFieldV7) split() []string {
	return strings.Split(string(*req), ",")
}
