package optional

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	o := Empty[string]()
	assert.True(t, o.IsEmpty())
	assert.False(t, o.IsPresent())
	assert.Panics(t, func() {
		o.Get()
	})

	o.IfPresent(func(value *string) {
		panic("this will never be called")
	})

	elsePassed := false
	o.IfPresentOrElse(func(value *string) {
		panic("this will never be called")
	}, func() {
		elsePassed = true
	})
	assert.True(t, elsePassed)

	val := o.OrElseGet(func() *string {
		str := "super-wow-string"
		return &str
	})
	assert.Equal(t, "super-wow-string", *val)

	assert.Panics(t, func() { o.OrElsePanic(errors.New("booom")) })
}

func TestWithValue(t *testing.T) {
	str := "wow-string"
	o := Of[string](&str)

	assert.False(t, o.IsEmpty())
	assert.True(t, o.IsPresent())
	assert.Equal(t, str, *o.Get())

	ifPresentCalled := false
	o.IfPresent(func(value *string) {
		ifPresentCalled = true
	})
	assert.True(t, ifPresentCalled)

	ifPresentCalled = false
	o.IfPresentOrElse(func(value *string) {
		ifPresentCalled = true
	}, func() {
		panic("this will never be called")
	})
	assert.True(t, ifPresentCalled)

	val := o.OrElseGet(func() *string {
		str := "string-that-will-be-never-seen"
		return &str
	})
	assert.Equal(t, str, *val)

	assert.NotPanics(t, func() { o.OrElsePanic(errors.New("booom")) })
}
