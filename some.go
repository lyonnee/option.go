package option

import optionenum "github.com/lyonnee/option.go/enum"

type Some[T any] struct {
	value *T
}

func (o Some[T]) Value() *T {
	return o.value
}

func (o Some[T]) Kind() optionenum.Enum {
	return optionenum.Some
}

func (o Some[T]) Match(someFunc func(*T) error, noneFunc func() error) {
	someFunc(o.value)
}

func NewSome[T any](value T) Option[T] {
	return Some[T]{value: &value}
}
