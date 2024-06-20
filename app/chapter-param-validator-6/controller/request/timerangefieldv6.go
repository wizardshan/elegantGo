package request

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type TimeRangeFieldV6 struct {
	TimeSplitterV6
}

func (req *TimeRangeFieldV6) UnmarshalJSON(b []byte) error {
	return req.unmarshalJSON(b, req.parseElement)
}

func (req *TimeRangeFieldV6) parseElement(value string) (time.Time, error) {
	var t time.Time
	var err error
	if value == "" {
		return t, nil
	}
	layouts := []string{time.DateTime, time.DateOnly}
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
	}

	return t, errors.New(fmt.Sprintf("the time layout should be one of %s", strings.Join(layouts, " and ")))
}
