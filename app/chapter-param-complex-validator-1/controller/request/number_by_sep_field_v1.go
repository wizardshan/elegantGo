package request

import (
	"encoding/json"
	"github.com/elliotchance/pie/v2"
	"strconv"
	"strings"
)

type NumberBySepFieldV1 struct {
	Values []int
	data   string
}

func (req *NumberBySepFieldV1) UnmarshalJSON(b []byte) error {

	// 解析json字符串
	if err := req.unmarshal(b); err != nil {
		return err
	}
	// 解析json字符串到业务数据
	return req.parse()
}

func (req *NumberBySepFieldV1) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *NumberBySepFieldV1) parse() error {
	elements := strings.Split(req.data, req.separator())
	err := validate.Var(elements, "dive,number")
	if err != nil {
		return err
	}

	req.Values = req.mapper(elements)
	return nil
}

func (req *NumberBySepFieldV1) mapper(elements []string) []int {
	return pie.Map(elements, func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	})
}

func (req *NumberBySepFieldV1) separator() string {
	return ","
}
