package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeJsonFieldV1 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeJsonFieldV1) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	ss := strings.Split(data, ",")
	if len(ss) != 2 {
		return errors.New("the capacity expected value is 2")
	}

	startStr := ss[0]
	endStr := ss[1]

	if startStr != "" {
		var err error
		if req.Start, err = time.Parse(time.DateTime, startStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error:%w", time.DateTime, err)
		}
	}

	if endStr != "" {
		var err error
		if req.End, err = time.Parse(time.DateTime, endStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error:%w", time.DateTime, err)
		}
	}

	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("start must lt end")
	}

	return nil
}
