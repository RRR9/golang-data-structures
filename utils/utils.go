package utils

import (
	"algorithms/containers"
	"reflect"
)

func FindIf[T any](begin containers.Iterator[T], end containers.Iterator[T]) {
	if reflect.TypeOf(begin) != reflect.TypeOf(end) {
		panic("different iterators type")
	}

	for ; begin != end; begin.Next() {

	}
}
