package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV5 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV5) UnmarshalJSON(b []byte) error {

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

func (req *DateTimeRangeFieldV5) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *DateTimeRangeFieldV5) split(data string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", req.layout()))
	return
}

func (req *DateTimeRangeFieldV5) parse(elements []string) {

	times := make([]time.Time, len(elements))
	for i, e := range elements {
		times[i], _ = time.Parse(req.layout(), e)
	}
	req.Start, req.End = times[0], times[1]
}

func (req *DateTimeRangeFieldV5) layout() string {
	return time.DateTime
}

func (req *DateTimeRangeFieldV5) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}

func (req *DateTimeRangeFieldV5) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *DateTimeRangeFieldV5) EndAble() bool {
	return !req.End.IsZero()
}
