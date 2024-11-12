package request

import (
	"elegantGo/param-complex-validate-1/pkg/stringx"
	"encoding/json"
)

type StringsField struct {
	Values []string
}

func (req *StringsField) UnmarshalJSON(b []byte) (err error) {
	var data string
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}
	splitter, err := stringx.NewSplitter(data)
	if err != nil {
		return
	}
	req.Values = splitter.Strings()
	return
}
