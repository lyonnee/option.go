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
	assert.Equal(t, 101, o.Value())
}

type TestStruct struct {
	i int
	s string
}

func TestNilSome(t *testing.T) {
	var bs []byte
	bsp := option.Some(&bs)
	assert.Equal(t, optionenum.Some, bsp.Kind())

	nbs := option.Some[[]byte](nil)
	assert.Equal(t, optionenum.None, nbs.Kind())

	s := option.Some[TestStruct](TestStruct{
		i: 1,
		s: "1",
	})
	assert.Equal(t, optionenum.Some, s.Kind())

	ns := option.Some[TestStruct](TestStruct{})
	assert.Equal(t, optionenum.None, ns.Kind())

	sp := option.Some[*TestStruct](&TestStruct{})
	assert.Equal(t, optionenum.Some, sp.Kind())

	nsp := option.Some[*TestStruct](nil)
	assert.Equal(t, optionenum.None, nsp.Kind())
}
