package request

import (
	"errors"
	"fmt"
	"time"
)

type Range[T time.Time | *int] struct {
	JsonStringBySep[T]
	Start T
	End   T
}

func (req *Range[T]) Parse(b []byte, validateTag string, mapperFunc func([]string) []T, validFunc func(start T, end T) bool) error {

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

func (req *Range[T]) parse(validateTag string, mapperFunc func([]string) []T) error {
	if err := validate.Var(req.elements, validateTag); err != nil {
		return fmt.Errorf("validate error: %w", err)
	}

	targets := mapperFunc(req.elements)
	req.Start, req.End = targets[0], targets[1]
	return nil
}

func (req *Range[T]) validateTag(tag string) string {
	tags := "len=2,dive,omitempty"
	if tag != "" {
		tags = tags + "," + tag
	}
	return tags
}
