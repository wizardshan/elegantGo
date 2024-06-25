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
}

func (req *DateTimeRangeFieldV2) UnmarshalJSON(b []byte) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	elements, err := req.split(data)
	if err != nil {
		return err
	}

	if err = req.parse(elements); err != nil {
		return err
	}

	if err = req.valid(); err != nil {
		return err
	}

	return nil
}

func (req *DateTimeRangeFieldV2) unmarshal(b []byte) (data string, err error) {
	// 解析json字符串
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	// 校验字符串有效性
	if find := strings.Contains(data, ","); !find {
		err = errors.New("parameter should be separated by commas")
	}
	return
}

func (req *DateTimeRangeFieldV2) split(data string) (elements []string, err error) {
	// 解析字符串为range数组并检验range数组有效性
	elements = strings.Split(data, ",")
	capacity := len(elements)
	if capacity != 2 {
		err = errors.New(fmt.Sprintf("the rangeField capacity expected value is 2, the result is %d", capacity))
	}
	return
}

func (req *DateTimeRangeFieldV2) parse(elements []string) (err error) {
	// 解析range数组中的开始时间和结束时间
	startStr := elements[0]
	endStr := elements[1]

	if startStr != "" {
		req.Start, err = time.Parse(time.DateTime, startStr)
		if err != nil {
			return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
		}
	}

	if endStr != "" {
		req.End, err = time.Parse(time.DateTime, endStr)
		if err != nil {
			return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
		}
	}

	return
}

func (req *DateTimeRangeFieldV2) valid() error {
	// 检验开始时间和结束时间的逻辑有效性
	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("the rangeField start must lt end")
	}
	return nil
}

func (req *DateTimeRangeFieldV2) StartAble() bool {
	return !req.Start.IsZero()
}

func (req *DateTimeRangeFieldV2) EndAble() bool {
	return !req.End.IsZero()
}
