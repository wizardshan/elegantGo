package request

import (
	"encoding/json"
	"errors"
	"strings"
)

type StringBySeparatorFieldV1 struct {
	Values []string
}

func (req *StringBySeparatorFieldV1) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	req.Values = strings.Split(data, ",")
	if err := validate.Var(req.Values, "dive,required"); err != nil {
		return err
	}

	return nil
}

func (req *StringBySeparatorFieldV1) Able() bool {
	return req != nil
}
