package option

import (
	"reflect"

	optionenum "github.com/lyonnee/option.go/enum"
)

type SomeType[T any] struct {
	value *T
}

func (o *SomeType[T]) Value() T {
	return *o.value
}

func (o *SomeType[T]) Kind() optionenum.Enum {
	return optionenum.Some
}

func (o *SomeType[T]) Match(someFn func(*T), noneFn func()) {
	someFn(o.value)
}

func Some[T any](value T) Option[T] {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	switch t.Kind() {
	case reflect.Invalid:
		return None[T]()
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		if v.IsNil() {
			return None[T]()
		}
	default:
		if v.IsZero() {
			return None[T]()
		}
	}

	return &SomeType[T]{value: &value}
}
