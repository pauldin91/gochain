package utils

func FilterBy[K any, V any](items []K, v V, predicate func(K, V) bool) []K {
	var filtered []K
	for _, c := range items {
		if predicate(c, v) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func Aggregate[K any](items []K, fn func(first K, second K) K) K {
	var reduced K = items[0]
	for _, t := range items {
		reduced = fn(reduced, t)
	}
	return reduced

}

func SelectMany[K any, V any](items []K, m *map[string]V, fn func(f *K, m *map[string]V)) {
	for _, v := range items {
		fn(&v, m)
	}
}
