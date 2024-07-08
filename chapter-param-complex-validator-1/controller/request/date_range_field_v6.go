package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type DateRangeFieldV6 struct {
	Start time.Time
	End   time.Time
	data  string
}

func (req *DateRangeFieldV6) UnmarshalJSON(b []byte) error {

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

func (req *DateRangeFieldV6) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *DateRangeFieldV6) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *DateRangeFieldV6) parse() error {
	elements := strings.Split(req.data, req.separator())
	layout := time.DateOnly
	err := validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,datetime=%s", layout))
	if err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	times := pie.Map(elements, func(s string) (t time.Time) {
		t, _ = time.Parse(layout, s)
		return
	})

	req.Start, req.End = times[0], times[1]
	return nil
}

func (req *DateRangeFieldV6) valid() bool {
	return !req.Start.IsZero() && !req.End.IsZero() && req.Start.Before(req.End)
}

func (req *DateRangeFieldV6) separator() string {
	return ","
}
