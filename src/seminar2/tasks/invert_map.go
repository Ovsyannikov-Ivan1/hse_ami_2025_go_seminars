package tasks

func invertMap[K comparable, V comparable](s map[K]V) map[V]K {
	a := make(map[V]K)
	for k, v := range s {
		a[v] = k
	}
	return a
}
