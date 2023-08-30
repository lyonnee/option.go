package option

import optionenum "github.com/lyonnee/option.go/enum"

type Option[T any] interface {
	Value() *T
	Kind() optionenum.Enum
	Match(someFunc func(*T) error, noneFunc func() error)
}
