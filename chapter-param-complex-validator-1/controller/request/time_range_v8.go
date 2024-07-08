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
	data  string
}

func (req *TimeRangeV8) Parse(b []byte, validateTag string, mapperFunc func([]string) []time.Time) error {

	// 解析json字符串
	if err := req.unmarshal(b); err != nil {
		return err
	}
	// 验证json字符串有效性
	if !req.hasSep() {
		return errors.New("parameter should be separated by commas")
	}
	// 解析json字符串到业务数据
	if err := req.parse(validateTag, mapperFunc); err != nil {
		return err
	}
	// 验证业务数据有效性
	if !req.valid() {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *TimeRangeV8) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *TimeRangeV8) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *TimeRangeV8) parse(validateTag string, mapperFunc func([]string) []time.Time) error {
	elements := strings.Split(req.data, req.separator())
	if err := validate.Var(elements, validateTag); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	times := mapperFunc(elements)
	req.Start, req.End = times[0], times[1]
	return nil
}

func (req *TimeRangeV8) mapper(elements []string, layout string) []time.Time {
	return pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})
}

func (req *TimeRangeV8) valid() bool {
	return req.Start.Before(req.End)
}

func (req *TimeRangeV8) separator() string {
	return ","
}

func (req *TimeRangeV8) validateTag(tag string) string {
	tags := "len=2,dive,omitempty"
	if tag != "" {
		tags = tags + "," + tag
	}
	return tags
}
