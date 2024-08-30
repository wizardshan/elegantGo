package request

import (
	"elegantGo/chapter-param-complex-validator-2/pkg/stringx"
	"encoding/json"
)

type StringsField string

func (req *StringsField) Values() ([]string, error) {
	validator := func(s string) bool {
		return s != ""
	}
	return stringx.NewSplitter(string(*req)).Validator(validator).Parse().Strings()
}

type StringsJsonField struct {
	Values []string
}

func (req *StringsJsonField) UnmarshalJSON(b []byte) (err error) {
	var data StringsField
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	req.Values, err = data.Values()
	return
}
