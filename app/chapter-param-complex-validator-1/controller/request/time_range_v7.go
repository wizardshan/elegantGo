package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type TimeRangeV7 struct {
	Start  time.Time
	End    time.Time
	layout string
}

func (req *TimeRangeV7) unmarshalJSON(b []byte, layout string) error {
	req.layout = layout

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

func (req *TimeRangeV7) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *TimeRangeV7) split(data string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", req.layout))
	return
}

func (req *TimeRangeV7) parse(elements []string) {

	times := pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(req.layout, s)
		return
	})

	req.Start, req.End = times[0], times[1]
}

func (req *TimeRangeV7) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}
