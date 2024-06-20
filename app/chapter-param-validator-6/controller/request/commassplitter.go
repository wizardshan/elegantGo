package request

import (
	"encoding/json"
	"errors"
	"strings"
)

type CommasSplitter struct {
}

func (req *CommasSplitter) Parse(b []byte) (elements []string, err error) {

	var data string
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	if find := strings.Contains(data, ","); !find {
		err = errors.New("parameter should be separated by commas")
		return
	}

	elements = strings.Split(data, ",")
	return
}
