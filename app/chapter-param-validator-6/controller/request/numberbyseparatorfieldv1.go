package request

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type NumberBySeparatorFieldV1 struct {
	Values []int
}

func (req *NumberBySeparatorFieldV1) UnmarshalJSON(b []byte) error {

	var data string
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if find := strings.Contains(data, ","); !find {
		return errors.New("parameter should be separated by commas")
	}

	elements := strings.Split(data, ",")
	if err := validate.Var(req.Values, "dive,number"); err != nil {
		return err
	}
	for _, element := range elements {
		num, _ := strconv.Atoi(element)
		req.Values = append(req.Values, num)
	}
	return nil
}

func (req *NumberBySeparatorFieldV1) Able() bool {
	return req != nil
}
