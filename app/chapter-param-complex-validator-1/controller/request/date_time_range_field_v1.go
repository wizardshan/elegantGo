package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeFieldV1 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeFieldV1) UnmarshalJSON(b []byte) error {
	// 解析json字符串
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	// 校验字符串有效性
	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	// 解析字符串为range数组并检验range数组有效性
	elements := strings.Split(data, ",")
	capacity := len(elements)
	if capacity != 2 {
		return errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}

	// 解析range数组中的开始时间和结束时间
	startStr := elements[0]
	endStr := elements[1]

	var err error
	if startStr != "" {
		req.Start, err = time.Parse(time.DateTime, startStr)
		if err != nil {
			return errors.New(fmt.Sprintf("the time layout should be `%s`", time.DateTime))
		}
	}

	if endStr != "" {
		req.End, err = time.Parse(time.DateTime, endStr)
		if err != nil {
			return errors.New(fmt.Sprintf("the time layout should be `%s`", time.DateTime))
		}
	}

	// 检验开始时间和结束时间的逻辑有效性
	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *DateTimeRangeFieldV1) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *DateTimeRangeFieldV1) EndAble() bool {
	return !req.End.IsZero()
}
