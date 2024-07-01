package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type DateRangeFieldV6 struct {
	Start time.Time
	End   time.Time
}

func (req *DateRangeFieldV6) UnmarshalJSON(b []byte) error {

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

func (req *DateRangeFieldV6) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *DateRangeFieldV6) split(data string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", req.layout()))
	return
}

func (req *DateRangeFieldV6) parse(elements []string) {

	times := pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(req.layout(), s)
		return
	})

	req.Start, req.End = times[0], times[1]
}

func (req *DateRangeFieldV6) layout() string {
	return time.DateOnly
}

func (req *DateRangeFieldV6) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End)
}
