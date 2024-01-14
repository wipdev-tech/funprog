package funprog

// func Map implements the common high-order function `map`. It takes a
// function and a slice of some type, applies the function on each element of
// the slice, and returns a new slice of the same type.
func Map[T any](f func(T) T, s []T) []T {
	out := make([]T, len(s))
	for i, el := range s {
		out[i] = f(el)
	}
	return out
}
