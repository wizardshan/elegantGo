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
	if capacity := len(elements); capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	startStr := elements[0]
	endStr := elements[1]

	if startStr != "" {
		var err error
		if req.Start, err = time.Parse(time.DateTime, startStr); err != nil {
			return errors.New(fmt.Sprintf("the time layout should be `%s`", time.DateTime))
		}
	}

	if endStr != "" {
		var err error
		if req.End, err = time.Parse(time.DateTime, endStr); err != nil {
			return errors.New(fmt.Sprintf("the time layout should be `%s`", time.DateTime))
		}
	}

	if req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}
