package optional

type Optional[T any] struct {
	value *T
}

// EmptyOptional returns an Optional without a value.
func EmptyOptional[T any]() Optional[T] {
	return Optional[T]{}
}

// OptionalOf returns an Optional wrapping the
// value given in input.
func OptionalOf[T any](value *T) Optional[T] {
	return Optional[T]{value: value}
}

// IfPresent runs the fn function given in input if Optional has a value.
// The fn function receives in input the actual value inside Optional
func (o Optional[T]) IfPresent(fn func(value *T)) {
	if o.IsPresent() {
		fn(o.value)
	}
}

// IfPresentOrElse runs the first fn function given in input if Optional has a value,
// otherwise it runs the second fn2 function given in input.
func (o Optional[T]) IfPresentOrElse(fn func(value *T), fn2 func()) {
	if o.IsPresent() {
		fn(o.value)
	} else {
		fn2()
	}
}

// Get returns the value of Optional. It will panic if no value is available.
// Before calling Get, make sure that there is a value by calling IsPresent.
func (o Optional[T]) Get() *T {
	if !o.IsPresent() {
		panic("tried to get an empty value!")
	}
	return o.value
}

// IsPresent returns true if the Optional is not empty.
func (o Optional[T]) IsPresent() bool {
	return o.value != nil
}

// OrElseGet returns the current value inside Optional value,
// or calls the fn function if no value is available, and returns its result.
func (o Optional[T]) OrElseGet(fn func() *T) *T {
	if o.IsPresent() {
		return o.value
	} else {
		return fn()
	}
}

// OrElsePanic returns the current value inside Optional value,
// or panics with the error given in input.
func (o Optional[T]) OrElsePanic(err error) *T {
	if !o.IsPresent() {
		panic(err)
	}
	return o.value
}
