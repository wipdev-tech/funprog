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

// func Reduce implements the common high-order function `reduce`. It takes a
// function and a slice. The input function must have two parameters so that
// the first one would be the "accumulator" and the second would be the next
// element in the slice.
//
// Note that the initial value of the accumulator is set to the zero value of
// whatever type (T) returned by Reduce.
func Reduce[T any](f func(T, T) T, s []T) T {
	var acc T
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// func Comp is a high-order function that implements composition. Given an
// arbitrary number of functions as input, Comp will return a function so that
// Comp(f, g)(x) == f(g(x)).
func Comp[T any](fs ...func(T) T) func(T) T {
	if len(fs) == 0 {
		return func(x T) T {
			return x
		}
	}

	return func(x T) T {
		f1 := fs[0]
		rest := fs[1:]
		return f1(Comp(rest...)(x))
	}
}
