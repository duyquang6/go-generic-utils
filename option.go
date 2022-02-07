package genericutils

// Option represents optional values. Instances of Option are either an instance of scala.Some or the object None.
type Option[T any] struct {
	val     T
	defined bool
}

// Some return option with value embedded
func Some[T any](val T) Option[T] {
	return Option[T]{val: val, defined: true}
}

// None return option with empty value
func None[T any]() Option[T] {
	return Option[T]{defined: false}
}

// IsDefined returns true if the option is an instance of Some, false otherwise.
func (o Option[A]) IsDefined() bool {
	return o.defined
}

// IsEmpty returns true if the option is an instance of None, false otherwise.
func (o Option[A]) IsEmpty() bool {
	return !o.defined
}

// Get returns embedded value if option is defined, panic otherwise
func (o Option[A]) Get() A {
	if !o.defined {
		panic("option is empty")
	}
	return o.val
}

// GetOrElse returns embedded value if option is defined, default value otherwise
func (o Option[A]) GetOrElse(alt A) A {
	if !o.defined {
		return alt
	}
	return o.val
}
