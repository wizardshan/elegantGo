package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type TimeRangeV8 struct {
	Start time.Time
	End   time.Time
}

func (req *TimeRangeV8) unmarshalJSON(b []byte, validateTag string, parseFunc func(s string) time.Time) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	elements, err := req.split(data, validateTag)
	if err != nil {
		return err
	}

	req.parse(elements, parseFunc)

	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeV8) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *TimeRangeV8) split(data string, validateTag string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,%s", validateTag))
	return
}

func (req *TimeRangeV8) parse(elements []string, f func(s string) time.Time) {

	times := pie.Map(elements, f)
	req.Start, req.End = times[0], times[1]
}

func (req *TimeRangeV8) parseElement(s string, layout string) (t time.Time) {
	t, _ = time.Parse(layout, s)
	return
}

func (req *TimeRangeV8) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}

func (req *TimeRangeV8) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *TimeRangeV8) EndAble() bool {
	return !req.End.IsZero()
}
