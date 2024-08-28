package request

import (
	"strconv"
	"strings"
)

type IntsFieldV1 string

// 只转换有效的数据
func (req *IntsFieldV1) Values() []int {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	return values
}

// 无效数据报错，停止转换
func (req *IntsFieldV1) MustValues() ([]int, error) {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

// 尽力而为，无效数据转换失败，默认零值
func (req *IntsFieldV1) ShouldValues() []int {
	ss := strings.Split(string(*req), ",")
	var values []int
	for _, s := range ss {
		value, _ := strconv.Atoi(s)
		values = append(values, value)
	}
	return values
}
