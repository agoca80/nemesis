package main

func Filter[L ~[]E, E any](list L, fn func(E) bool) (filtered L) {
	filtered = L{}
	for _, element := range list {
		if fn(element) {
			filtered = append(filtered, element)
		}
	}
	return
}
