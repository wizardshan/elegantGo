package separator

import (
	"github.com/elliotchance/pie/v2"
	"strconv"
	"strings"
	"time"
)

type StringV1 struct {
	data   string
	sep    string
	values []string
}

func NewV1(data string, sep string) StringV1 {
	s := StringV1{data: data, sep: sep}
	s.split()
	return s
}

func (req *StringV1) Ints() []int {
	return pie.Ints(req.values)
}

func (req *StringV1) PtrInts() []*int {
	return pie.Map(req.values, func(s string) *int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &num
	})
}

func (req *StringV1) Values() []string {
	return req.values
}

func (req *StringV1) Dates() []time.Time {
	return req.parseTimes(time.DateOnly)
}

func (req *StringV1) DateTimes() []time.Time {
	return req.parseTimes(time.DateTime)
}

func (req *StringV1) Times() []time.Time {
	layouts := []string{time.DateTime, time.DateOnly}
	return pie.Map(req.values, func(s string) (t time.Time) {
		for _, layout := range layouts {
			t, _ = time.Parse(layout, s)
			if !t.IsZero() {
				return
			}
		}
		return
	})
}

func (req *StringV1) parseTimes(layout string) []time.Time {
	return pie.Map(req.values, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func (req *StringV1) split() {
	req.values = strings.Split(req.data, req.sep)
}

func (req *StringV1) HasSep() bool {
	return strings.Contains(req.data, req.sep)
}
