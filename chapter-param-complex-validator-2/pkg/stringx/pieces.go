package stringx

import (
	"errors"
	"github.com/elliotchance/pie/v2"
	"strings"
	"time"
)

type Pieces struct {
	input     string
	ss        []string
	sep       string
	validator func(s string) bool
	err       error
}

func NewPieces(input string) *Pieces {
	piece := new(Pieces)
	piece.input = input
	return piece
}

func (p *Pieces) Ints() ([]int, error) {
	if p.err != nil {
		return nil, p.err
	}
	return pie.Ints(p.ss), nil
}

func (p *Pieces) PtrInts() ([]*int, error) {
	if p.err != nil {
		return nil, p.err
	}
	return ToPtrInts(p.ss), nil
}

func (p *Pieces) Strings() ([]string, error) {
	if p.err != nil {
		return nil, p.err
	}
	return p.ss, nil
}

func (p *Pieces) DateTimes() ([]time.Time, error) {
	if p.err != nil {
		return nil, p.err
	}
	return ToDateTimes(p.ss), nil
}

func (p *Pieces) Dates() ([]time.Time, error) {
	if p.err != nil {
		return nil, p.err
	}
	return ToDates(p.ss), nil
}

func (p *Pieces) Times() ([]time.Time, error) {
	if p.err != nil {
		return nil, p.err
	}
	return ToTimes(p.ss), nil
}

func (p *Pieces) Sep(sep string) *Pieces {
	p.sep = sep
	return p
}

func (p *Pieces) Validator(fn func(s string) bool) *Pieces {
	p.validator = fn
	return p
}

func (p *Pieces) Parse() *Pieces {
	p.ss = p.split()
	if p.validator != nil && !pie.All(p.ss, p.validator) {
		p.err = errors.New("one of pieces is failure")
	}
	return p
}

func (p *Pieces) split() []string {
	return strings.Split(p.input, p.getSep())
}

func (p *Pieces) getSep() string {
	sep := p.defaultSep()
	if p.sep != "" {
		sep = p.sep
	}
	return sep
}

func (p *Pieces) defaultSep() string {
	return ","
}
