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
	data  string
}

func (req *DateTimeRangeFieldV4) UnmarshalJSON(b []byte) error {

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

func (req *DateTimeRangeFieldV4) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateTimeRangeFieldV4) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *DateTimeRangeFieldV4) parse() error {
	elements := strings.Split(req.data, req.separator())
	layout := time.DateTime
	err := validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", layout))
	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	startStr := elements[0]
	endStr := elements[1]

	if startStr != "" {
		req.Start, _ = time.Parse(layout, startStr)
	}

	if endStr != "" {
		req.End, _ = time.Parse(layout, endStr)
	}
	return nil
}

func (req *DateTimeRangeFieldV4) valid() bool {
	return req.Start.Before(req.End)
}

func (req *DateTimeRangeFieldV4) separator() string {
	return ","
}
