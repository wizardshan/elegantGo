package request

import (
	"elegantGo/chapter-param-complex-validator-2/pkg/stringx"
	"encoding/json"
)

type StringsField string

func (req *StringsField) Values() ([]string, error) {
	splitter, err := stringx.NewSplitter(string(*req))
	if err != nil {
		return nil, err
	}
	return splitter.Strings(), nil
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
