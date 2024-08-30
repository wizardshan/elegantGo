package stringx

import (
	"github.com/elliotchance/pie/v2"
	"time"
)

func ToDateTimes(ss []string) []time.Time {
	return pie.Map(ss, ToDateTime)
}

func ToDates(ss []string) []time.Time {
	return pie.Map(ss, ToDate)
}

func ToTimes(ss []string) []time.Time {
	return pie.Map(ss, ToTime)
}

func ToDateTime(value string) time.Time {
	return timeParse(time.DateTime, value)
}

func ToDate(value string) time.Time {
	return timeParse(time.DateOnly, value)
}

func ToTime(value string) time.Time {
	t := ToDateTime(value)
	if !t.IsZero() {
		return t
	}

	return ToDate(value)
}

func IsDateTime(value string) bool {
	return isTime(time.DateTime, value)
}

func IsDate(value string) bool {
	return isTime(time.DateOnly, value)
}

func IsTime(value string) bool {
	if IsDateTime(value) || IsDate(value) {
		return true
	}
	return false
}

func isTime(layout string, value string) bool {
	if _, err := time.Parse(layout, value); err == nil {
		return true
	}
	return false
}

func timeParse(layout string, value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}
