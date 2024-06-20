package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV1 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV1) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	elements := strings.Split(data, ",")
	capacity := len(elements)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	startStr := elements[0]
	endStr := elements[1]

	var err error
	if startStr != "" {
		req.Start, err = time.Parse(time.DateTime, startStr)
		if err != nil {
			return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
		}
	}

	if endStr != "" {
		req.End, err = time.Parse(time.DateTime, endStr)
		if err != nil {
			return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
		}
	}

	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV1) parse(value string) (time.Time, error) {
	var t time.Time
	if value == "" {
		return t, nil
	}
	layout := time.DateTime
	layoutErrMsg := fmt.Sprintf("the time layout should be `%s`", layout)

	var err error
	t, err = time.Parse(layout, value)
	if err != nil {
		return t, errors.New(layoutErrMsg)
	}

	return t, nil
}

func (req *DateTimeRangeFieldV1) StartAble() bool {
	if req == nil {
		return false
	}

	return !req.Start.IsZero()
}

func (req *DateTimeRangeFieldV1) EndAble() bool {
	if req == nil {
		return false
	}

	return !req.End.IsZero()
}
