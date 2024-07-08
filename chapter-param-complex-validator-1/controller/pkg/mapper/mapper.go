package mapper

import (
	"github.com/elliotchance/pie/v2"
	"strconv"
	"time"
)

func Ints(ss []string) []int {
	return pie.Ints(ss)
}

func PtrInts(ss []string) []*int {
	return pie.Map(ss, func(s string) *int {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &num
	})
}

func Dates(ss []string) []time.Time {
	return parseTimes(ss, time.DateOnly)
}

func DateTimes(ss []string) []time.Time {
	return parseTimes(ss, time.DateTime)
}

func parseTimes(ss []string, layout string) []time.Time {
	return pie.Map(ss, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func Times(ss []string) []time.Time {
	layouts := []string{time.DateTime, time.DateOnly}
	return pie.Map(ss, func(s string) (t time.Time) {
		for _, layout := range layouts {
			t, _ = time.Parse(layout, s)
			if !t.IsZero() {
				return
			}
		}
		return
	})
}
