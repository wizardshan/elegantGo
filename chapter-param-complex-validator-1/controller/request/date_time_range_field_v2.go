package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV2 struct {
	Start time.Time
	End   time.Time
	data  string
}

func (req *DateTimeRangeFieldV2) UnmarshalJSON(b []byte) error {

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

func (req *DateTimeRangeFieldV2) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateTimeRangeFieldV2) hasSep() bool {
	return strings.Contains(req.data, ",")
}

func (req *DateTimeRangeFieldV2) parse() error {
	elements := strings.Split(req.data, ",")
	if capacity := len(elements); capacity != 2 {
		return fmt.Errorf("the rangeField capacity expected value is 2, the result is %d", capacity)
	}

	startStr := elements[0]
	endStr := elements[1]

	if startStr != "" {
		var err error
		if req.Start, err = time.Parse(time.DateTime, startStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error: %w", time.DateTime, err)
		}
	}

	if endStr != "" {
		var err error
		if req.End, err = time.Parse(time.DateTime, endStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error: %w", time.DateTime, err)
		}
	}
	return nil
}

func (req *DateTimeRangeFieldV2) valid() bool {
	return req.Start.Before(req.End)
}
