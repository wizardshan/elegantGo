package stringx

import (
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type Option func(*Splitter)

func Sep(sep string) Option {
	return func(s *Splitter) {
		s.sep = sep
	}
}

func Validator(validator func(s string) bool) Option {
	return func(o *Splitter) {
		o.validator = validator
	}
}

type Splitter struct {
	input     string
	ss        []string
	sep       string
	validator func(s string) bool
}

func NewSplitter(input string, options ...Option) (*Splitter, error) {
	spr := Splitter{
		sep:       ",",
		input:     input,
		validator: NotEmpty,
	}

	for _, option := range options {
		option(&spr)
	}

	if err := spr.parse(); err != nil {
		return nil, err
	}
	return &spr, nil
}

func (spr *Splitter) parse() error {
	spr.ss = strings.Split(spr.input, spr.sep)
	if spr.validator != nil && !pie.All(spr.ss, spr.validator) {
		return errors.New("one of substring validation failed")
	}
	return nil
}

func (spr *Splitter) Ints() []int {
	return pie.Ints(spr.ss)
}

func (spr *Splitter) PtrInts() []*int {
	return ToPtrInts(spr.ss)
}

func (spr *Splitter) Strings() []string {
	return spr.ss
}

func (spr *Splitter) DateTimes() []time.Time {
	return ToDateTimes(spr.ss)
}

func (spr *Splitter) Dates() []time.Time {
	return ToDates(spr.ss)
}

func (spr *Splitter) Times() []time.Time {
	return ToTimes(spr.ss)
}
