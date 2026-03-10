package aoc

import (
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

// NewSet[T] returns a new Set[T]
func NewSet[T comparable](values ...T) Set[T] {
	s := Set[T]{}
	s.AddSlice(values)
	return s
}

// Add values to the set
func (s Set[T]) Add(values ...T) Set[T] { return s.AddSlice(values) }

// AddSlice to the set
func (s Set[T]) AddSlice(values []T) Set[T] {
	for _, value := range values {
		s[value] = struct{}{}
	}
	return s
}

// AddSet to the set
func (s Set[T]) AddSet(set Set[T]) Set[T] {
	for key := range set {
		s[key] = struct{}{}
	}
	return s
}

// Remove values from set
func (s Set[T]) Remove(values ...T) Set[T] {
	for _, value := range values {
		delete(s, value)
	}
	return s
}

// Contains returns if a value is in the set
func (s Set[T]) Contains(value T) bool {
	_, c := s[value]
	return c
}

// ContainsAll returns if all values are in the set
func (s Set[T]) ContainsAll(values []T) bool {
	for _, v := range values {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

// ContainsAny returns if any of the values are in the set
func (s Set[T]) ContainsAny(values []T) bool {
	for _, v := range values {
		if s.Contains(v) {
			return true
		}
	}
	return false
}

// Subtract returns the differences
func (s Set[T]) Subtract(x Set[T]) (difference Set[T]) {
	difference = Set[T]{}
	for k := range s {
		if !x.Contains(k) {
			difference.Add(k)
		}
	}
	return
}

// Values returns the values in set
func (s Set[T]) Values() []T { return slices.Collect(maps.Keys(s)) }

// Intersect returns the differences
func (s Set[T]) Intersect(x Set[T]) (intersection Set[T]) {
	intersection = Set[T]{}
	for k := range s {
		if x.Contains(k) {
			intersection.Add(k)
		}
	}
	return
}

// Union returns the combination of two sets
func (s Set[T]) Union(x Set[T]) (union Set[T]) {
	union = Set[T]{}
	union.AddSet(s)
	union.AddSet(x)
	return
}
