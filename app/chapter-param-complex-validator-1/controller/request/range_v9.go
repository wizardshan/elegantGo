package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type RangeV9[T time.Time | *int] struct {
}

func (req *RangeV9[T]) parse(b []byte, validateTag string, parseFunc func(s string) T, validFunc func(start T, end T) bool) (start T, end T, err error) {

	var data string
	data, err = req.unmarshal(b)
	if err != nil {
		return
	}

	var elements []string
	elements, err = req.split(data, validateTag)
	if err != nil {
		return
	}

	targets := pie.Map(elements, parseFunc)
	start = targets[0]
	end = targets[1]

	if !validFunc(start, end) {
		err = errors.New("the rangeField start must lt end")
		return
	}

	return
}

func (req *RangeV9[T]) unmarshal(b []byte) (data string, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func (req *RangeV9[T]) split(data string, validateTag string) (elements []string, err error) {
	separator := ","
	if find := strings.Contains(data, separator); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}
	elements = strings.Split(data, separator)
	err = validate.Var(elements, fmt.Sprintf("len=2,dive,omitempty,%s", validateTag))
	return
}
