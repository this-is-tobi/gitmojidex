package utils

import "regexp"

func Map[E, T any](s []E, f func(E) T) []T {
	mapped := make([]T, len(s))
	for i, x := range s {
		mapped[i] = f(x)
	}
	return mapped
}

func Reduce[E any, R any](s []E, f func(R, E) R, initial R) R {
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
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

func Regex(a string, b string) bool {
	return regexp.MustCompile(a).FindString(b) != ""
}
