package util

import (
	"github.com/asaskevich/govalidator"
	"github.com/elliotchance/pie/v2"
	"golang.org/x/exp/constraints"
)

var (
	IsInt = govalidator.IsInt
)

func Int[T constraints.Ordered](x T) int {
	return pie.Int(x)
}

func Ints[T constraints.Ordered](ss []T) []int {
	return pie.Ints(ss)
}

func Filter[T any](ss []T, condition func(T) bool) []T {
	return pie.Filter(ss, condition)
}
