package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV2 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV2) UnmarshalJSON(b []byte) error {

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

	layout := time.DateTime
	layoutErrMsg := fmt.Sprintf("the time layout should be `%s`", layout)

	if startStr != "" {
		var err error
		req.Start, err = time.Parse(layout, startStr)
		if err != nil {
			return errors.New(layoutErrMsg)
		}
	}

	if endStr != "" {
		var err error
		req.End, err = time.Parse(layout, endStr)
		if err != nil {
			return errors.New(layoutErrMsg)
		}
	}

	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV2) StartAble() bool {
	if req == nil {
		return false
	}

	return !req.Start.IsZero()
}

func (req *DateTimeRangeFieldV2) EndAble() bool {
	if req == nil {
		return false
	}

	return !req.End.IsZero()
}
