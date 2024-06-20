package request

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeRangeField struct {
	TimeSplitter
}

func (req *TimeRangeField) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *TimeRangeField) parseElement(value string) (t time.Time, err error) {
	if value == "" {
		return
	}
	layouts := []string{time.DateTime, time.DateOnly}
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return
		}
	}

	return t, errors.New(fmt.Sprintf("the time layout should be one of %s", strings.Join(layouts, " and ")))
}
