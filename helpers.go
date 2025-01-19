package goutils

func RemoveElement[K comparable](s []K, k K) []K {
	ret := make([]K, 0, len(s)-1)
	for i, c := range s {
		if c == k {
			ret = append(ret, s[:i]...)
			ret = append(ret, s[i+1:]...)
			break
		}
	}
	if len(ret) == 0 {
		return s
	}
	return ret
}

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
