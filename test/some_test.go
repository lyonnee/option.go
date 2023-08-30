package test

import (
	"testing"

	"github.com/lyonnee/option.go"
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

	assert.Equal(t, 4, v, "Should not be executed noneFn")
	assert.Equal(t, 2, v)
	assert.Equal(t, 101, *(o.Value()))
}
