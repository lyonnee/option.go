package option

import (
	optionenum "github.com/lyonnee/option.go/enum"
)

type SomeType[T any] struct {
	value *T
}

func (o SomeType[T]) Value() *T {
	return o.value
}

func (o SomeType[T]) Kind() optionenum.Enum {
	return optionenum.Some
}

func (o SomeType[T]) Match(someFn func(*T), noneFn func()) {
	someFn(o.value)
}

func Some[T any](value T) Option[T] {
	return SomeType[T]{value: &value}
}
