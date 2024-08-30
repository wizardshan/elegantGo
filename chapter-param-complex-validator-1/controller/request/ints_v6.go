package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/stringx"
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
)

type IntsFieldV6 string

func (req *IntsFieldV6) Values() []int {
	ss := req.split()

	var ssFiltered []string
	for _, s := range ss {
		if stringx.IsInt(s) {
			ssFiltered = append(ssFiltered, s)
		}
	}

	// 最终版本
	return pie.Ints(ssFiltered)

	// 版本二
	//return pie.Map(ssFiltered, numeral.ToInt)

	// 版本一
	//return pie.Map(ssFiltered, func(s string) int {
	//	return pie.Int(s)
	//})
}

func (req *IntsFieldV6) MustValues() ([]int, error) {
	ss := req.split()

	for _, s := range ss {
		if !stringx.IsInt(s) {
			return nil, errors.New(s + " is not an integer")
		}
	}

	return pie.Ints(ss), nil
}

func (req *IntsFieldV6) ShouldValues() []int {
	ss := req.split()

	return pie.Ints(ss)
}

func (req *IntsFieldV6) split() []string {
	return strings.Split(string(*req), ",")
}
