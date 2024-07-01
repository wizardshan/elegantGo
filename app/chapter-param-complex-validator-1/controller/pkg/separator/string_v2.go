package separator

import (
	"app/chapter-param-complex-validator-1/controller/pkg/mapper"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type StringV2 struct {
	data   string
	sep    string
	values []string
}

func NewV2(data string, sep string) StringV2 {
	s := StringV2{data: data, sep: sep}
	s.split()
	return s
}

func (req *StringV2) Ints() []int {
	return pie.Ints(req.values)
}

func (req *StringV2) PtrInts() []*int {
	return mapper.PtrInts(req.values)
}

func (req *StringV2) Strings() []string {
	return req.values
}

func (req *StringV2) Dates() []time.Time {
	return mapper.Dates(req.values)
}

func (req *StringV2) DateTimes() []time.Time {
	return mapper.DateTimes(req.values)
}

func (req *StringV2) Times() []time.Time {
	return mapper.Times(req.values)
}

func (req *StringV2) split() {
	req.values = strings.Split(req.data, req.sep)
}

func (req *StringV2) HasSep() bool {
	return strings.Contains(req.data, req.sep)
}
