package request

import (
	"github.com/asaskevich/govalidator"
	"github.com/elliotchance/pie/v2"
	"strings"
)

type IntsFieldV3 string

func (req *IntsFieldV3) Values() []int {
	ss := req.split()

	ssFiltered := pie.Filter(ss, govalidator.IsInt)
	return pie.Map(ssFiltered, func(s string) int {
		return pie.Int(s)
	})
}

func (req *IntsFieldV3) MustValues() ([]int, error) {
	ss := req.split()

	if err := validate.Var(ss, "dive,int"); err != nil {
		return nil, err
	}

	return pie.Map(ss, func(s string) int {
		return pie.Int(s)
	}), nil
}

func (req *IntsFieldV3) ShouldValues() []int {
	ss := req.split()

	return pie.Map(ss, func(s string) int {
		return pie.Int(s)
	})
}

func (req *IntsFieldV3) split() []string {
	return strings.Split(string(*req), ",")
}
