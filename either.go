package generic

// Either represents a value of one of two possible types (a disjoint union). An instance of Either is an instance of either Left or Right.
type Either[E, T any] struct {
	left   E
	right  T
	isLeft bool
}

// Left return Either with left value embedded
func Left[E, T any](rval T) Either[E, T] {
	return Either[E, T]{right: rval}
}

// Right return Either with right value embedded
func Right[E, T any](lval E) Either[E, T] {
	return Either[E, T]{left: lval, isLeft: true}
}

// IsLeft return if it is left value
func (e Either[E, T]) IsLeft() bool {
	return e.isLeft
}

// IsRight return if it is right value
func (e Either[E, T]) IsRight() bool {
	return !e.isLeft
}

// GetOrElse returns the value from this Right or the given argument if this is a Left.
func (e Either[E, T]) GetOrElse(rval T) T {
	if e.isLeft {
		return rval
	}
	return e.right
}

// Get returns the value from this Right or panic if it is left value.
func (e Either[E, T]) Get() T {
	if e.isLeft {
		panic("either is left")
	}
	return e.right
}

// ToOption returns a Some containing the Right value if it exists or a None if this is a Left.
func (e Either[E, T]) ToOption() Option[T] {
	if e.isLeft {
		return None[T]()
	}
	return Some(e.right)
}
