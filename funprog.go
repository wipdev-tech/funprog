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

// func Filter implements the common high-order function `filter`. It takes a
// function (predicate) and a slice of some type, and it returns a slice
// containing the elements of the input slice for which the predicate returns
// true.
func Filter[T any](p func(T) bool, s []T) []T {
	out := make([]T, 0, len(s))
	for _, el := range s {
		if p(el) {
			out = append(out, el)
		}
	}
	return out
}
