package request

import (
	"encoding/json"
	"github.com/elliotchance/pie/v2"
	"strconv"
	"strings"
)

type NumberBySeparatorV1 struct {
	Values []int
}

func (req *NumberBySeparatorV1) UnmarshalJSON(b []byte) error {

	data, err := req.unmarshal(b)
	if err != nil {
		return err
	}

	var elements []string
	elements, err = req.split(data)
	if err != nil {
		return err
	}

	req.Values = pie.Map(elements, func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	})

	return nil
}

func (req *NumberBySeparatorV1) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *NumberBySeparatorV1) split(data string) (elements []string, err error) {
	separator := ","
	elements = strings.Split(data, separator)
	err = validate.Var(elements, "dive,number")
	return
}
