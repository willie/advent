package aoc

func Map[T, V any](f func(T) V, in []T) (out []V) {
	out = make([]V, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return
}

func Filter[T any](f func(T) bool, in []T) (out []T) {
	for _, v := range in {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

func FilterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

/*
func Reduce[E1, E2 any](f func(E2, E1) E2, in []E1, init E2) E2 {
	r := init
	for _, v := range in {
		r = f(r, v)
	}
	return r
}
*/
