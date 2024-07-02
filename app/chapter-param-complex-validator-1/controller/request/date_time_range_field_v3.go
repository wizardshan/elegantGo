package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV3 struct {
	Start time.Time
	End   time.Time
	data  string
}

func (req *DateTimeRangeFieldV3) UnmarshalJSON(b []byte) error {

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

func (req *DateTimeRangeFieldV3) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateTimeRangeFieldV3) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *DateTimeRangeFieldV3) parse() error {
	elements := strings.Split(req.data, req.separator())
	if capacity := len(elements); capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	startStr := elements[0]
	endStr := elements[1]

	layout := time.DateTime
	timeParseErrorMsg := fmt.Sprintf("the time layout should be `%s`", layout)

	if startStr != "" {
		var err error
		if req.Start, err = time.Parse(layout, startStr); err != nil {
			return errors.New(timeParseErrorMsg)
		}
	}

	if endStr != "" {
		var err error
		if req.End, err = time.Parse(layout, endStr); err != nil {
			return errors.New(timeParseErrorMsg)
		}
	}
	return nil
}

func (req *DateTimeRangeFieldV3) valid() bool {
	return req.Start.Before(req.End)
}

func (req *DateTimeRangeFieldV3) separator() string {
	return ","
}
