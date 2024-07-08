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
	data  string
}

func (req *DateTimeRangeFieldV5) UnmarshalJSON(b []byte) error {

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

func (req *DateTimeRangeFieldV5) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateTimeRangeFieldV5) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *DateTimeRangeFieldV5) parse() error {
	elements := strings.Split(req.data, req.separator())
	layout := time.DateTime
	err := validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", layout))
	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	times := make([]time.Time, len(elements))
	for i, e := range elements {
		times[i], _ = time.Parse(layout, e)
	}
	req.Start, req.End = times[0], times[1]
	return nil
}

func (req *DateTimeRangeFieldV5) valid() bool {
	return req.Start.Before(req.End)
}

func (req *DateTimeRangeFieldV5) separator() string {
	return ","
}
