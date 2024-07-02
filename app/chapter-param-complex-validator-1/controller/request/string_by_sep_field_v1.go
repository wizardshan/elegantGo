package request

import (
	"encoding/json"
	"strings"
)

type StringBySepFieldV1 struct {
	Values []string
	data   string
}

func (req *StringBySepFieldV1) UnmarshalJSON(b []byte) error {

	// 解析json字符串
	if err := req.unmarshal(b); err != nil {
		return err
	}
	// 解析json字符串到业务数据
	return req.parse()
}

func (req *StringBySepFieldV1) unmarshal(b []byte) error {
	return json.Unmarshal(b, &req.data)
}

func (req *StringBySepFieldV1) parse() error {
	req.Values = strings.Split(req.data, req.separator())
	return validate.Var(req.Values, "dive,required,alphanum")
}

func (req *StringBySepFieldV1) separator() string {
	return ","
}
