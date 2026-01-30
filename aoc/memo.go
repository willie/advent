package aoc

// Memoize wraps a function with single argument and caches its results.
// Useful for recursive functions with overlapping subproblems (dynamic programming).
func Memoize[K comparable, V any](f func(K) V) func(K) V {
	cache := make(map[K]V)
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(k)
		cache[k] = v
		return v
	}
}

// Memoize2 wraps a function with two arguments and caches its results.
func Memoize2[K1, K2 comparable, V any](f func(K1, K2) V) func(K1, K2) V {
	type key struct {
		k1 K1
		k2 K2
	}
	cache := make(map[key]V)
	return func(k1 K1, k2 K2) V {
		k := key{k1, k2}
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(k1, k2)
		cache[k] = v
		return v
	}
}

// MemoizeRecursive creates a memoized recursive function.
// The function receives itself as the first argument to enable recursion.
func MemoizeRecursive[K comparable, V any](f func(recurse func(K) V, k K) V) func(K) V {
	cache := make(map[K]V)
	var memoized func(K) V
	memoized = func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(memoized, k)
		cache[k] = v
		return v
	}
	return memoized
}
