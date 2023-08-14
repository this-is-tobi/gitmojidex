package utils

func Map[E, T any](s []E, f func(E) T) []T {
	mapped := make([]T, len(s))
	for i, x := range s {
		mapped[i] = f(x)
	}
	return mapped
}

func Reduce[E any, R any](s []E, f func(R, E) R, initial R) R {
	acc := initial
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Filter[E any](s []E, f func(E) bool) []E {
	var filtered []E
	for _, x := range s {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}
