package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV4 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV4) UnmarshalJSON(b []byte) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	elements, err := req.split(data)
	if err != nil {
		return err
	}

	req.parse(elements)

	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV4) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *DateTimeRangeFieldV4) split(data string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", req.layout()))
	return
}

func (req *DateTimeRangeFieldV4) parse(elements []string) {
	startStr := elements[0]
	endStr := elements[1]

	if startStr != "" {
		req.Start, _ = time.Parse(req.layout(), startStr)
	}

	if endStr != "" {
		req.End, _ = time.Parse(req.layout(), endStr)
	}

}

func (req *DateTimeRangeFieldV4) layout() string {
	return time.DateTime
}

func (req *DateTimeRangeFieldV4) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}
