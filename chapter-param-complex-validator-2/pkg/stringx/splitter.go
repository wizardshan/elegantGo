package stringx

import (
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type Splitter struct {
	input     string
	ss        []string
	sep       string
	validator func(s string) bool
	err       error
}

func NewSplitter(input string) *Splitter {
	piece := new(Splitter)
	piece.input = input
	return piece
}

func (spr *Splitter) Ints() ([]int, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return pie.Ints(spr.ss), nil
}

func (spr *Splitter) PtrInts() ([]*int, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return ToPtrInts(spr.ss), nil
}

func (spr *Splitter) Strings() ([]string, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return spr.ss, nil
}

func (spr *Splitter) DateTimes() ([]time.Time, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return ToDateTimes(spr.ss), nil
}

func (spr *Splitter) Dates() ([]time.Time, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return ToDates(spr.ss), nil
}

func (spr *Splitter) Times() ([]time.Time, error) {
	if spr.err != nil {
		return nil, spr.err
	}
	return ToTimes(spr.ss), nil
}

func (spr *Splitter) Sep(sep string) *Splitter {
	spr.sep = sep
	return spr
}

func (spr *Splitter) Validator(fn func(s string) bool) *Splitter {
	spr.validator = fn
	return spr
}

func (spr *Splitter) Parse() *Splitter {
	spr.ss = spr.split()
	if spr.validator != nil && !pie.All(spr.ss, spr.validator) {
		spr.err = errors.New("one of pieces is failure")
	}
	return spr
}

func (spr *Splitter) split() []string {
	return strings.Split(spr.input, spr.getSep())
}

func (spr *Splitter) getSep() string {
	sep := spr.defaultSep()
	if spr.sep != "" {
		sep = spr.sep
	}
	return sep
}

func (spr *Splitter) defaultSep() string {
	return ","
}
