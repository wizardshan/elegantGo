package request

import (
	"github.com/asaskevich/govalidator"
	"github.com/elliotchance/pie/v2"
	"strings"
)

type IntsFieldV4 string

func (req *IntsFieldV4) Values() []int {
	return pie.Ints(pie.Filter(req.split(), govalidator.IsInt))
}

func (req *IntsFieldV4) MustValues() ([]int, error) {
	ss := req.split()
	if err := validate.Var(ss, "dive,int"); err != nil {
		return nil, err
	}
	return pie.Ints(ss), nil
}

func (req *IntsFieldV4) ShouldValues() []int {
	return pie.Ints(req.split())
}

func (req *IntsFieldV4) split() []string {
	return strings.Split(string(*req), ",")
}
