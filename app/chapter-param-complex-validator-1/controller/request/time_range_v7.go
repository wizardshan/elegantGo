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
	Start time.Time
	End   time.Time
	data  string
}

func (req *TimeRangeV7) Parse(b []byte, layout string) error {

	// 解析json字符串
	if err := req.unmarshal(b); err != nil {
		return err
	}
	// 验证json字符串有效性
	if !req.hasSep() {
		return errors.New("parameter should be separated by commas")
	}
	// 解析json字符串到业务数据
	if err := req.parse(layout); err != nil {
		return err
	}
	// 验证业务数据有效性
	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeV7) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *TimeRangeV7) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *TimeRangeV7) parse(layout string) error {
	elements := strings.Split(req.data, req.separator())
	err := validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", layout))
	if err != nil {
		return err
	}

	times := req.mapper(elements, layout)

	req.Start, req.End = times[0], times[1]
	return nil
}

func (req *TimeRangeV7) mapper(elements []string, layout string) []time.Time {
	return pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func (req *TimeRangeV7) valid() bool {
	return req.Start.Before(req.End)
}

func (req *TimeRangeV7) separator() string {
	return ","
}
