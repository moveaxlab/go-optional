package optional

import "errors"

var (
	TriedToGetEmptyValueError = errors.New("tried to get empty value")
)

// The Optional interface may or may not contain a non-null value.
type Optional[T any] interface {
	// IsEmpty returns true if the Optional is empty.
	IsEmpty() bool
	// IsPresent returns true if the Optional is not empty.
	IsPresent() bool
	// Get returns the value of Optional. It will panic if no value is available.
	// Before calling Get, make sure that there is a value by calling IsPresent.
	Get() *T
	// IfPresent runs the fn function given in input if Optional has a value.
	// The fn function receives in input the actual value inside Optional
	IfPresent(fn func(value *T))
	// IfPresentOrElse runs the first fn function given in input if Optional has a value,
	// otherwise it runs the second fn2 function given in input.
	IfPresentOrElse(fn func(value *T), fn2 func())
	// OrElseGet returns the current value inside Optional value,
	// or calls the fn function if no value is available, and returns its result.
	OrElseGet(fn func() *T) *T
	// OrElsePanic returns the current value inside Optional value,
	// or panics with the error given in input.
	OrElsePanic(err error) *T
}

type optional[T any] struct {
	value *T
}

// Empty returns an Optional without a value.
func Empty[T any]() Optional[T] {
	return &optional[T]{}
}

// Of returns an Optional wrapping the
// value given in input.
func Of[T any](value *T) Optional[T] {
	return &optional[T]{value: value}
}

func (o *optional[T]) IfPresent(fn func(value *T)) {
	if o.IsPresent() {
		fn(o.value)
	}
}

func (o *optional[T]) IfPresentOrElse(fn func(value *T), fn2 func()) {
	if o.IsPresent() {
		fn(o.value)
	} else {
		fn2()
	}
}

func (o *optional[T]) Get() *T {
	if !o.IsPresent() {
		panic(TriedToGetEmptyValueError)
	}
	return o.value
}

func (o *optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o *optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o *optional[T]) OrElseGet(fn func() *T) *T {
	if o.IsPresent() {
		return o.value
	} else {
		return fn()
	}
}

func (o *optional[T]) OrElsePanic(err error) *T {
	if !o.IsPresent() {
		panic(err)
	}
	return o.value
}
