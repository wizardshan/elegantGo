package request

import (
	"elegantGo/chapter-param-complex-validator-1/pkg/util"
	"strings"
)

type IntsField string

func (req *IntsField) Values() []int {
	return util.Ints(util.Filter(req.split(), util.IsInt))
}

func (req *IntsField) MustValues() ([]int, error) {
	ss := req.split()
	if err := util.Validate(ss, "dive,int"); err != nil {
		return nil, err
	}
	return util.Ints(ss), nil
}

func (req *IntsField) ShouldValues() []int {
	return util.Ints(req.split())
}

func (req *IntsField) split() []string {
	return strings.Split(string(*req), ",")
}
