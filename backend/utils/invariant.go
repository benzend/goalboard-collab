package utils

import (
	"reflect"
)

func Invariant[K any](v K, msg string) K {
	x := reflect.ValueOf(v)
	if x.IsZero() {
		panic(msg)
	}
	return v
}