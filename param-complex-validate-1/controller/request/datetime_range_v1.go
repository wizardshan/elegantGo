package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type DateTimeRangeJsonFieldV1 struct {
	Start time.Time
	End   time.Time
}

func (req *DateTimeRangeJsonFieldV1) UnmarshalJSON(b []byte) error {
	// 解析json字符串
	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	// 解析字符串为range数组并检验range数组长度有效性
	ss := strings.Split(data, ",")
	if len(ss) != 2 {
		return errors.New("the capacity expected value is 2")
	}

	// 解析range数组中的开始时间和结束时间
	startStr := ss[0]
	endStr := ss[1]

	if startStr != "" {
		var err error
		if req.Start, err = time.Parse(time.DateTime, startStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error:%w", time.DateTime, err)
		}
	}

	if endStr != "" {
		var err error
		if req.End, err = time.Parse(time.DateTime, endStr); err != nil {
			return fmt.Errorf("the time layout should be `%s`, error:%w", time.DateTime, err)
		}
	}

	// 检验范围有效性
	if !req.Start.IsZero() && !req.End.IsZero() && req.Start.After(req.End) {
		return errors.New("start must lt end")
	}

	return nil
}
