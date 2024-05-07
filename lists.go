package main

type List interface {
	Size()
	Item() any
}

func Filter[L ~[]E, E any](list L, fn func(E) bool) (filtered L) {
	filtered = L{}
	for _, element := range list {
		if fn(element) {
			filtered = append(filtered, element)
		}
	}
	return
}

func Each[L ~[]E, E any](list L, fn func(E)) {
	for _, element := range list {
		fn(element)
	}
}

// func Map[L ~[]E, E, R any](list L, fn func(E) R) (mapped []R) {
// 	for _, element := range list {
// 		mapped = append(mapped, fn(element))
// 	}
// 	return
// }

func Reduce[L ~[]E, E, R any](list L, fn func(R, E) R, initial R) (result R) {
	result = initial
	for _, element := range list {
		result = fn(result, element)
	}
	return
}
