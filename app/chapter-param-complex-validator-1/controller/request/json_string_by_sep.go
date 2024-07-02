package request

import (
	"encoding/json"
	"strings"
)

type JsonStringBySep[T any] struct {
	data     string
	elements []string
}

func (req *JsonStringBySep[T]) unmarshal(b []byte) error {
	if err := json.Unmarshal(b, &req.data); err != nil {
		return err
	}

	req.split()

	return nil
}

func (req *JsonStringBySep[T]) split() {
	req.elements = strings.Split(req.data, req.sep())
}

func (req *JsonStringBySep[T]) hasSep() bool {
	return strings.Contains(req.data, req.sep())
}

func (req *JsonStringBySep[T]) sep() string {
	return ","
}

func (req *JsonStringBySep[T]) mapper(f func(values []string) []T) []T {
	return f(req.elements)
}
