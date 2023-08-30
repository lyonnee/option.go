package test

import (
	"testing"

	"github.com/lyonnee/option.go"
	optionenum "github.com/lyonnee/option.go/enum"
	"github.com/stretchr/testify/assert"
)

func TestSomeOp(t *testing.T) {
	o := option.Some(1)

	v := 1
	o.Match(func(i *int) {
		// 修改 o 的value值
		*i = *i + 100
		v = 2
	}, func() {
		v = 4
	})

	assert.NotEqual(t, 4, v, "Should not be executed noneFn")
	assert.Equal(t, 2, v)
	assert.Equal(t, 101, *(o.Value()))
}

type TestStruct struct {
	i int
	s string
}

func TestNilSome(t *testing.T) {
	v := option.Some[[]byte](nil)
	assert.Equal(t, optionenum.None, v.Kind())

	s := option.Some[*TestStruct](nil)
	assert.Equal(t, optionenum.None, s.Kind())

	s1 := option.Some[*TestStruct](&TestStruct{})
	assert.Equal(t, optionenum.Some, s1.Kind())

	var bs []byte
	bss := option.Some(&bs)
	assert.Equal(t, optionenum.Some, bss.Kind())
}
