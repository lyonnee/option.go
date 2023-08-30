package test

import (
	"testing"

	"github.com/lyonnee/option.go"
	"github.com/stretchr/testify/assert"
)

func TestNoneOp(t *testing.T) {
	o := option.None[any]()

	v := 1
	o.Match(func(i *any) {
		v = 2
	}, func() {
		v = 4
	})

	assert.Equal(t, 2, v, "Should not be executed someFn")
	assert.Equal(t, 4, v)
}
