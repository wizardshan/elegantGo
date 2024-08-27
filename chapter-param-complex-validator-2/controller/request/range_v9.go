package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type RangeV9[T time.Time | *int] struct {
	Start T
	End   T
	data  string
}

func (req *RangeV9[T]) Parse(b []byte, validateTag string, mapperFunc func([]string) []T, validFunc func(start T, end T) bool) error {

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
	if !validFunc(req.Start, req.End) {
		return errors.New("the rangeField start must lt end")
	}

	return nil
}

func (req *RangeV9[T]) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *RangeV9[T]) hasSep() bool {
	return strings.Contains(req.data, req.separator())
}

func (req *RangeV9[T]) parse(validateTag string, mapperFunc func([]string) []T) error {
	elements := strings.Split(req.data, req.separator())
	if err := validate.Var(elements, validateTag); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	targets := mapperFunc(elements)
	req.Start, req.End = targets[0], targets[1]
	return nil
}

func (req *RangeV9[T]) separator() string {
	return ","
}

func (req *RangeV9[T]) validateTag(tag string) string {
	tags := "len=2,dive,omitempty"
	if tag != "" {
		tags = tags + "," + tag
	}
	return tags
}
