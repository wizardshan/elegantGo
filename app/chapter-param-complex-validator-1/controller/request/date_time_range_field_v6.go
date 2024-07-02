package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type DateTimeRangeFieldV6 struct {
	Start time.Time
	End   time.Time
	data  string
}

func (req *DateTimeRangeFieldV6) UnmarshalJSON(b []byte) error {

	// 解析json字符串
	if err := req.unmarshal(b); err != nil {
		return err
	}
	// 验证json字符串有效性
	if !req.hasSep() {
		return errors.New("parameter should be separated by commas")
	}
	// 解析json字符串到业务数据
	if err := req.parse(); err != nil {
		return err
	}
	// 验证业务数据有效性
	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV6) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateTimeRangeFieldV6) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *DateTimeRangeFieldV6) parse() error {
	elements := strings.Split(req.data, req.separator())
	layout := time.DateTime
	err := validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", layout))
	if err != nil {
		return err
	}

	times := req.mapper(elements, layout)

	//times := pie.Map(elements, func(s string) (t time.Time) {
	//	t, _ = time.Parse(layout, s)
	//	return
	//})

	//times := iterator.Map(elements, func(s string) (t time.Time) {
	//	t, _ = time.Parse(req.layout(), s)
	//	return
	//})

	req.Start, req.End = times[0], times[1]
	return nil
}

func (req *DateTimeRangeFieldV6) mapper(elements []string, layout string) []time.Time {
	return pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func (req *DateTimeRangeFieldV6) valid() bool {
	return req.Start.Before(req.End)
}

func (req *DateTimeRangeFieldV6) separator() string {
	return ","
}
