package option

import (
	optionenum "github.com/lyonnee/option.go/enum"
)

type NoneType[T any] struct {
}

func (o NoneType[T]) Value() *T {
	return nil
}

func (o NoneType[T]) Kind() optionenum.Enum {
	return optionenum.None
}

func (o NoneType[T]) Match(someFn func(*T), noneFn func()) {
	noneFn()
}

func None[T any]() Option[T] {
	return NoneType[T]{}
}
