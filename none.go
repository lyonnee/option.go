package option

import optionenum "github.com/lyonnee/option.go/enum"

type None[T any] struct {
}

func (o None[T]) Value() *T {
	return nil
}

func (o None[T]) Kind() optionenum.Enum {
	return optionenum.None
}

func (o None[T]) Match(someFunc func(*T) error, noneFunc func() error) {
	noneFunc()
}

func NewNone[T any]() Option[T] {
	return None[T]{}
}
