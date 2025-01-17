package goutils

func Keys[K comparable, T any](m map[K]T) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func PtrTo[T any](t T) *T {
	return &t
}
