package aoc

import "slices"

// Reverse returns a new slice with elements in reverse order.
// Note: For in-place reversal, use slices.Reverse from stdlib.
func Reverse[T any](s []T) []T {
	result := slices.Clone(s)
	slices.Reverse(result)
	return result
}

// ReverseInPlace reverses a slice in place.
// Deprecated: Use slices.Reverse from stdlib instead.
func ReverseInPlace[T any](s []T) {
	slices.Reverse(s)
}

// Chunk splits a slice into chunks of size n.
// The last chunk may have fewer than n elements.
func Chunk[T any](s []T, n int) [][]T {
	if n <= 0 {
		return nil
	}
	var chunks [][]T
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// Window returns all sliding windows of size n.
func Window[T any](s []T, n int) [][]T {
	if n <= 0 || n > len(s) {
		return nil
	}
	windows := make([][]T, len(s)-n+1)
	for i := range windows {
		windows[i] = s[i : i+n]
	}
	return windows
}

// Pairs returns all adjacent pairs from the slice.
func Pairs[T any](s []T) [][2]T {
	if len(s) < 2 {
		return nil
	}
	pairs := make([][2]T, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		pairs[i] = [2]T{s[i], s[i+1]}
	}
	return pairs
}

// Count returns the number of elements matching the predicate.
func Count[T any](s []T, predicate func(T) bool) int {
	count := 0
	for _, v := range s {
		if predicate(v) {
			count++
		}
	}
	return count
}

// All returns true if all elements match the predicate.
func All[T any](s []T, predicate func(T) bool) bool {
	for _, v := range s {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any returns true if any element matches the predicate.
func Any[T any](s []T, predicate func(T) bool) bool {
	for _, v := range s {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Find returns the first element matching the predicate and true,
// or zero value and false if not found.
func Find[T any](s []T, predicate func(T) bool) (T, bool) {
	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element matching the predicate,
// or -1 if not found.
func FindIndex[T any](s []T, predicate func(T) bool) int {
	for i, v := range s {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Unique returns a new slice with duplicate elements removed.
// Preserves order of first occurrence.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Flatten converts a 2D slice to a 1D slice.
func Flatten[T any](s [][]T) []T {
	var result []T
	for _, inner := range s {
		result = append(result, inner...)
	}
	return result
}

// Zip combines two slices into pairs. Stops at the shorter slice.
func Zip[T, U any](a []T, b []U) [][2]any {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	result := make([][2]any, n)
	for i := 0; i < n; i++ {
		result[i] = [2]any{a[i], b[i]}
	}
	return result
}

// Reduce applies a function to accumulate values.
func Reduce[T, U any](s []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range s {
		result = f(result, v)
	}
	return result
}
